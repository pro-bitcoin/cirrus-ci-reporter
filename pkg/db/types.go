package db

import "time"

type Repo struct {
	ID            uint64 `bun:"id,pk"`
	Owner         string `bun:",notnull"`
	Name          string `bun:",notnull"`
	DefaultBranch string `bun:",notnull"`
	URL           string `bun:",notnull"`
}

type Build struct {
	ID            uint64    `bun:"id,pk"`
	RepoID        uint64    `bun:",notnull"`
	Repo          *Repo     `bun:"rel:belongs-to,join:repo_id=id"`
	Created       time.Time `bun:",notnull"`
	Duration      uint32
	ClockDuration uint32
	Branch        string `bun:",notnull"`
	PullRequest   uint   `bun:",notnull"`
	Status        string
}

type Task struct {
	ID        uint64    `bun:"id,pk"`
	BuildID   uint64    `bun:",notnull"`
	Build     *Build    `bun:"rel:belongs-to,join:build_id=id"`
	Name      string    `bun:",notnull"`
	Duration  uint32    `bun:",notnull"`
	Status    string    `bun:",notnull"`
	Created   time.Time `bun:",notnull"`
	Scheduled time.Time `bun:",notnull"`
	Executing time.Time `bun:",notnull"`
}

type Artifact struct {
	ID       uint64 `bun:"id,pk,unique:grp"`
	TaskID   uint64 `bun:",notnull"`
	Task     *Task  `bun:"rel:belongs-to,join:task_id=id"`
	Type     string `bun:",notnull"`
	Location string `bun:",unique:grp"`
}
