package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/pro-bitcoin/cirrus-reporter/pkg/builds"
	"github.com/pro-bitcoin/cirrus-reporter/pkg/db"
	"github.com/spf13/cobra"
	"github.com/uptrace/bun"
)

func NewBuildsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "builds",
		Short: "sync builds",
		RunE: func(c *cobra.Command, args []string) error {
			return getBuilds(c.Context(), *config)
		},
	}

	return cmd
}

func getBuilds(ctx context.Context, cfg configOptions) error {
	pdb := DBInstance(cfg)
	var repos []db.Repo
	if err := pdb.NewSelect().Model(&repos).Scan(ctx, &repos); err != nil {
		return err
	}
	for _, r := range repos {
		lastBuild := []db.Build{}
		if err := pdb.NewSelect().Model(&lastBuild).Where("repo_id = ?", r.ID).Order("created DESC").Limit(1).Scan(ctx, &lastBuild); err != nil {
			return err
		}
		after := time.Now().UTC().Add(-1 * (time.Hour * 24) * 30)
		if len(lastBuild) > 0 {
			after = lastBuild[0].Created
		}
		fmt.Printf("getting builds for %s after %s\n", r.Name, after)
		if err := getBuildRepo(ctx, r, after, pdb); err != nil {
			return err
		}
	}

	return nil
}

func getBuildRepo(ctx context.Context, r db.Repo, after time.Time, pdb *bun.DB) error {
	b, err := builds.NewBuildSync(HTTPClient(*config))
	if err != nil {
		return err
	}
	edges, err := b.Get(ctx, r.ID, after)
	if err != nil {
		return err
	}

	// nolint:wrapcheck
	return builds.Process(ctx, pdb, r.ID, edges)
}
