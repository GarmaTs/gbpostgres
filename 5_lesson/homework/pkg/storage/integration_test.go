//go:build integration
// +build integration

package storage_test

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
	"lesson5homework/pkg/storage"
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
		check   func(*testing.T, []storage.Snippet, error)
	}{
		{
			name:   "random_test_header_found",
			store:  storage.NewPG(dbpool),
			ctx:    context.Background(),
			prefix: "random_test_header",
			limit:  5,
			prepare: func(dbpool *pgxpool.Pool) {
				dbpool.Exec(context.Background(),
					`insert into snippets (main_theme, header, content)
				values
					('test_theme', 'random_test_header','test_content');`)
			},
			check: func(t *testing.T, hints []storage.Snippet, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, hints)
			},
		},
		{
			name:   "no_random_test_header",
			store:  storage.NewPG(dbpool),
			ctx:    context.Background(),
			prefix: "test_header",
			limit:  5,
			prepare: func(dbpool *pgxpool.Pool) {
				dbpool.Exec(context.Background(),
					`delete from snippets where header = 'random_test_header';`)
			},
			check: func(t *testing.T, hints []storage.Snippet, err error) {
				require.NoError(t, err)
				require.Empty(t, hints)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(dbpool)
			hints, err := tt.store.FindSnippetByHeader(tt.ctx, tt.prefix, tt.limit)
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
