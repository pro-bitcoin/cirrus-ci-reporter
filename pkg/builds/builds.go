package builds

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hasura/go-graphql-client"
	"github.com/pro-bitcoin/cirrus-reporter/pkg/cirrus"
	cdb "github.com/pro-bitcoin/cirrus-reporter/pkg/db"
	"github.com/uptrace/bun"
)

type Syncer interface {
	Get(ctx context.Context, ID uint64, after time.Time) ([]cirrus.Edge, error)
}

type cirrusClient struct {
	c *graphql.Client
}

func (c *cirrusClient) Get(ctx context.Context, id uint64, after time.Time) ([]cirrus.Edge, error) {
	query := cirrus.Query{}
	graphVariables := map[string]interface{}{
		"id":    fmt.Sprintf("%d", id),
		"after": graphql.String(fmt.Sprintf("%d", after.UnixMilli())),
	}
	err := c.c.Query(ctx, &query, graphVariables)
	if err != nil {
		return nil, err
	}
	edges := query.Repository.Builds.Edges
	if query.Repository.Builds.PageInfo.HasNextPage {
		t := epochMSToTime(string(query.Repository.Builds.PageInfo.EndCursor))
		fmt.Printf("getting builds after %s\n", t)
		moreEdges, err := c.Get(ctx, id, t)
		if err != nil {
			return nil, err
		}
		edges = append(edges, moreEdges...)
	}

	return edges, nil
}

const (
	millisPerSecond     = int64(time.Second / time.Millisecond)
	nanosPerMillisecond = int64(time.Millisecond / time.Nanosecond)
)

func epochToTime(v graphql.Float) time.Time {
	msInt := int64(v)

	return time.Unix(msInt/millisPerSecond,
		(msInt%millisPerSecond)*nanosPerMillisecond).UTC()
}

func epochMSToTime(epoch string) time.Time {
	msInt, _ := strconv.ParseInt(epoch, 10, 64)

	return time.Unix(msInt/millisPerSecond,
		(msInt%millisPerSecond)*nanosPerMillisecond).UTC()
}

// nolint:wrapcheck
func Process(ctx context.Context, db *bun.DB, repoID uint64, edges []cirrus.Edge) error {
	for _, edge := range edges {
		n := edge.Node
		if n.Status == "COMPLETED" || n.Status == "FAILED" {
			bID, _ := strconv.ParseUint(string(n.ID), 10, 64)
			bld := &cdb.Build{
				ID:            bID,
				RepoID:        repoID,
				Repo:          nil,
				Created:       epochToTime(n.BuildCreatedTimestamp),
				Duration:      uint32(n.DurationInSeconds),
				ClockDuration: uint32(n.ClockDurationInSeconds),
				Branch:        string(n.Branch),
				PullRequest:   uint(n.PullRequest),
				Status:        string(n.Status),
				Commit:        string(n.ChangeIdInRepo),
			}
			_, err := db.NewInsert().Model(bld).Exec(ctx)
			if err != nil {
				return err
			}
			for _, task := range n.Tasks {
				tID, _ := strconv.ParseUint(string(task.ID), 10, 64)
				tsk := &cdb.Task{
					ID:        tID,
					BuildID:   bld.ID,
					Build:     nil,
					Name:      string(task.Name),
					Duration:  uint32(task.DurationInSeconds),
					Status:    string(task.Status),
					Created:   epochToTime(task.CreationTimestamp),
					Scheduled: epochToTime(task.ScheduledTimestamp),
					Executing: epochToTime(task.ExecutingTimestamp),
				}
				if tsk.Status == "COMPLETED" || tsk.Status == "FAILED" {
					_, err = db.NewInsert().Model(tsk).Exec(ctx)
					if err != nil {
						return err
					}
					if err = processArtifacts(ctx, tID, task, db); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func processArtifacts(ctx context.Context, taskID uint64, tsk cirrus.Task, db *bun.DB) error {
	for _, a := range tsk.Artifacts {
		for _, f := range a.Files {
			art := &cdb.Artifact{
				TaskID:   taskID,
				Name:     string(a.Name),
				Type:     string(a.Type),
				Location: string(f.Path),
			}
			if _, err := db.NewInsert().Model(art).Exec(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}

func NewBuildSync(c *http.Client) (Syncer, error) {
	return &cirrusClient{
		c: graphql.NewClient("https://api.cirrus-ci.com/graphql", c),
	}, nil
}
