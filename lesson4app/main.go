package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Description of pgxpool config
func poolDescr() {
	ctx := context.Background()

	url := "postgres://gopher:P@ssw0rd@localhost:5432/gopher_corp"
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatal(err)
	}

	cfg.MaxConns = 8 // Максимальное кол-во соединений
	cfg.MinConns = 4 // Минимальное кол-во соединений

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
		log.Fatal(err)
	}
	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(ctx, "select 'Hello, World!'").Scan(&greeting)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greeting)
}

type (
	Phone string
	Email string
)

type EmailSearchHint struct {
	Phone Phone
	Email Email
}

func search(ctx context.Context, dbpool *pgxpool.Pool, prefix string, limit int) ([]EmailSearchHint, error) {
	const sql = `select email, phone from employees 
	where email like $1 
	order by email asc
	limit $2;`

	pattern := prefix + "%"
	rows, err := dbpool.Query(ctx, sql, pattern, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}
	defer rows.Close()

	var hints []EmailSearchHint
	for rows.Next() {
		var hint EmailSearchHint
		err = rows.Scan(&hint.Email, &hint.Phone)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		hints = append(hints, hint)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to read response: %w", rows.Err())
	}

	return hints, nil
}

type AttackResults struct {
	Duration         time.Duration
	Threads          int
	QueriesPerformed uint64
}

func attack(ctx context.Context, duration time.Duration, threads int, dbpool *pgxpool.Pool) AttackResults {
	var queries uint64

	attacker := func(stopAt time.Time) {
		for {
			_, err := search(ctx, dbpool, "alex", 5)
			if err != nil {
				log.Fatal(err)
			}
			atomic.AddUint64(&queries, 1)
			if time.Now().After(stopAt) {
				return
			}
		}
	}

	var wg sync.WaitGroup
	wg.Add(threads)

	startAt := time.Now()
	stopAt := startAt.Add(duration)

	for i := 0; i < threads; i++ {
		go func() {
			attacker(stopAt)
			wg.Done()
		}()
	}
	wg.Wait()

	return AttackResults{
		Duration:         time.Now().Sub(startAt),
		Threads:          threads,
		QueriesPerformed: queries,
	}
}

func main() {
	ctx := context.Background()

	url := "postgres://gopher:P@ssw0rd@localhost:5432/gopher_corp"
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatal(err)
	}

	cfg.MaxConns = 20
	cfg.MinConns = 20

	dbpool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	// Пример использования транзакции
	/*
		var employee1, employee2 EmployeeID = 7483465, 7483357
		salary := 20000
		err = update(ctx, dbpool, employee1, employee2, salary)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("salary updated")
	*/

	// Нагрузочное тестирование
	duration := time.Duration(10 * time.Second)
	threads := 40
	fmt.Println("start attack")
	res := attack(ctx, duration, threads, dbpool)

	fmt.Println("duration:", res.Duration)
	fmt.Println("threads:", res.Threads)
	fmt.Println("queries:", res.QueriesPerformed)
	qps := res.QueriesPerformed / uint64(res.Duration.Seconds())
	fmt.Println("QPS:", qps)
}
