package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"lesson5homework/pkg/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	DB_USER     = "gopher"
	DB_PASSWORD = "P@ssw0rd"
	DB_NAME     = "snippets"
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
)

func main() {
	store, ctx, closePool, err := initPGStore()
	if err != nil {
		log.Fatalf("Error occured while initializing pool: %s", err)
		return
	}
	defer closePool()

	prefix := "CAP"
	limit := 3
	snippets, err := store.FindSnippetByHeader(ctx, prefix, limit)
	if err != nil {
		log.Fatalf("Error occured while finding by header: %s", err)
		return
	}

	if len(snippets) == 0 {
		fmt.Println("No snippets found")
	} else {
		fmt.Println("Found", len(snippets), "snippets:")
		for _, sn := range snippets {
			fmt.Println(sn.Header)
		}
	}
}

func initPGStore() (*storage.PG, context.Context, func(), error) {
	ctx := context.Background()

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USER,
		DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, nil, nil, err
	}
	cfg.MaxConns = 30
	cfg.MinConns = 30
	cfg.HealthCheckPeriod = 1 * time.Minute // Частота проверки работоспособности соединения
	cfg.MaxConnLifetime = 1 * time.Hour     // Сколько времени будет жить соединение
	// Время жизни неиспользуемого соединения, если запросов не поступало,
	// то соединение закроется
	cfg.MaxConnIdleTime = 30 * time.Minute
	// Ограничение по времени на весь процесс установки соединения и аутентификации
	cfg.ConnConfig.ConnectTimeout = 1 * time.Second
	// Лимиты в net.Dialer позволяют достичь предсказуемого поведения в случае обрыва сети
	cfg.ConnConfig.DialFunc = (&net.Dialer{
		KeepAlive: cfg.HealthCheckPeriod,
		// Таймаут на установку соединения гарантирует что не будет зависаний
		// при попытке установить соединение
		Timeout: cfg.ConnConfig.ConnectTimeout,
	}).DialContext

	dbpool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, nil, nil, err
	}
	closeDefer := func() {
		defer dbpool.Close()
	}

	store := storage.NewPG(dbpool)

	return store, ctx, closeDefer, nil
}
