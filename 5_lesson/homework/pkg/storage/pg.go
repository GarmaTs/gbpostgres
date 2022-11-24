package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PG struct {
	dbpool *pgxpool.Pool
}

func NewPG(dbpool *pgxpool.Pool) *PG {
	return &PG{dbpool: dbpool}
}

type Snippet struct {
	Id        int
	Header    string
	Content   string
	MainTheme string
}

func (p *PG) FindSnippetByHeader(ctx context.Context, prefix string, limit int) ([]Snippet, error) {
	const sql = `select id, header, content, main_theme
		from snippets
		where header like $1
		order by id
		limit $2`
	pattern := prefix + "%"
	rows, err := p.dbpool.Query(ctx, sql, pattern, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %s", err)
	}
	defer rows.Close()

	var snippets []Snippet
	for rows.Next() {
		var sn Snippet
		err = rows.Scan(&sn.Id, &sn.Header, &sn.Content, &sn.MainTheme)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		snippets = append(snippets, sn)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to read response: %w", rows.Err())
	}

	return snippets, nil
}
