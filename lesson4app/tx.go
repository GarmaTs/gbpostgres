package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TransactionFunc func(context.Context, pgx.Tx) error

// Нет необходимости отправть BEGIN и COMMIT вручную, т.к. в драйвер заложена
// возможность работать с транзакциями

// inTx создает транзакцию и передает ее для использования в функцию f
// если в функии f происходит ошибка, транзакция откатывается
func inTx(ctx context.Context, dbpool *pgxpool.Pool, f TransactionFunc) error {
	transaction, err := dbpool.Begin(ctx)
	if err != nil {
		return err
	}

	err = f(ctx, transaction)
	if err != nil {
		rbErr := transaction.Rollback(ctx)
		if rbErr != nil {
			log.Print(rbErr)
			return err
		}
	}

	err = transaction.Commit(ctx)
	if err != nil {
		rbErr := transaction.Rollback(ctx)
		if rbErr != nil {
			log.Println(rbErr)
			return err
		}
	}

	return nil
}

type (
	EmployeeID   int
	PositionID   int
	DepartmentID int
)

func update(ctx context.Context, dbpool *pgxpool.Pool, id1, id2 EmployeeID, salary int) error {
	err := inTx(ctx, dbpool, func(ctx context.Context, tx pgx.Tx) error {
		const sql = `update employees set salary = salary+($1) where id = $2;`
		_, err := tx.Exec(ctx, sql, salary, id1)
		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx, sql, -salary, id2)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
