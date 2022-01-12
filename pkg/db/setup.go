package db

import (
	"context"

	"github.com/uptrace/bun"
)

// nolint:gochecknoglobals
var probitcoin = &Repo{
	ID:            5478867478511616,
	Owner:         "pro-bitcoin",
	Name:          "pro-bitcoin",
	DefaultBranch: "prometheus",
	URL:           "https://github.com/pro-bitcoin/pro-bitcoin",
}

// nolint:gochecknoglobals
var bitcoin = &Repo{
	ID:            6264162542157824,
	Owner:         "bitcoin",
	Name:          "bitcoin",
	DefaultBranch: "master",
	URL:           "https://github.com/bitcoin/bitcoin",
}

// nolint:wsl,wrapcheck
func Setup(db *bun.DB) error {
	ctx := context.Background()
	_, err := db.NewCreateTable().Model((*Repo)(nil)).Exec(ctx)
	if err != nil {
		return err
	}
	if _, err = db.NewCreateTable().Model((*Build)(nil)).Exec(ctx); err != nil {
		return err
	}
	if _, err = db.NewCreateTable().Model((*Task)(nil)).Exec(ctx); err != nil {
		return err
	}
	if _, err = db.NewCreateTable().Model((*Artifact)(nil)).Exec(ctx); err != nil {
		return err
	}
	if _, err = db.NewInsert().Model(probitcoin).Exec(ctx); err != nil {
		return err
	}
	_, err = db.NewInsert().Model(bitcoin).Exec(ctx)

	return err
}
