//go:build integration
// +build integration

package storage_test

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
	"lesson5app/pkg/emailhint/storage"
	"os"
	"testing"
)

func TestIntegrationSearch(t *testing.T) {
	ctx := context.Background()
	dbpool := connect(ctx)
	defer dbpool.Close()

	tests := []struct {
		name    string
		store   *storage.PG
		ctx     context.Context
		prefix  string
		limit   int
		prepare func(*pgxpool.Pool)
		check   func(*testing.T, []storage.EmailSearchHint, error)
	}{
		{
			name:   "success",
			store:  storage.NewPG(dbpool),
			ctx:    context.Background(),
			prefix: "alexandra",
			limit:  5,
			prepare: func(dbpool *pgxpool.Pool) {
				dbpool.Exec(context.Background(),
					`insert into employees (first_name, last_name, phone, email, salary, manager_id, department_id, position)
				values
					('Alex', 'Smith', '123456', 'alex@gopher_corp.com', 100000, 1, 1, 1),
					('Alexandra', 'Smith', '123456', 'alexandra@gopher_corp.com', 100000, 1, 1, 1);`)
			},
			check: func(t *testing.T, hints []storage.EmailSearchHint, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, hints)
			},
		},
	}

	for _, tt := range tests {
		//tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(dbpool)
			hints, err := tt.store.Search(tt.ctx, tt.prefix, tt.limit)
			tt.check(t, hints, err)
		})
	}
}

func connect(ctx context.Context) *pgxpool.Pool {
	dbpool, err := pgxpool.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	return dbpool
}
