package cmd

import (
	"database/sql"
	"net/http"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func DBInstance(config configOptions) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.DSN)))
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(config.Debug),
		bundebug.WithEnabled(config.Debug),
	))
	return db
}

type authToken struct {
	Token string
}

// nolint: wrapcheck
func (a *authToken) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Set("Authorization", "Bearer "+a.Token)

	return http.DefaultClient.Do(request)
}

func HTTPClient(config configOptions) *http.Client {
	return &http.Client{
		Transport: &authToken{
			Token: config.Token,
		},
	}
}
