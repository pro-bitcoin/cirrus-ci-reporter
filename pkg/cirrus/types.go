package cirrus

import "github.com/hasura/go-graphql-client"

type PageInfo struct {
	HasNextPage graphql.Boolean
	StartCursor graphql.String
	EndCursor   graphql.String
}

type ArtifactFile struct {
	Path graphql.String
}

type Artifact struct {
	Type  graphql.String
	Name  graphql.String
	Files []ArtifactFile
}

type Task struct {
	ID                 graphql.String
	Name               graphql.String
	DurationInSeconds  graphql.Int
	Status             graphql.String
	CreationTimestamp  graphql.Float
	ScheduledTimestamp graphql.Float
	ExecutingTimestamp graphql.Float
	Artifacts          []Artifact
}

type Node struct {
	ID                     graphql.String
	BuildCreatedTimestamp  graphql.Float
	DurationInSeconds      graphql.Int
	ClockDurationInSeconds graphql.Int
	Branch                 graphql.String
	PullRequest            graphql.Int
	Tasks                  []Task
	Status                 graphql.String
	ChangeIdInRepo         graphql.String
}

type Edge struct {
	Node Node
}

type Builds struct {
	PageInfo PageInfo
	Edges    []Edge
}

type Repository struct {
	// Id graphql.String
	Name          graphql.String
	Owner         graphql.String
	DefaultBranch graphql.String
	Builds        Builds `graphql:"builds(after: $after)"`
}

type Query struct {
	Repository Repository `graphql:"repository(id: $id)"`
}
