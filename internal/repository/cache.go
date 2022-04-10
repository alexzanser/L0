package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	AddOrderTask = `INSERT INTO orders (order_uid, data) VALUES($1, $2);`
	GetOrdersTask = `SELECT * FROM orders;`
)


type Cache struct {
	pool *pgxpool.Pool
}

func NewCache(pgxPool *pgxpool.Pool) Cache {
	return Cache{pool: pgxPool}
}

func (c *Cache) CacheOrder(orderID string, data []byte) error {
	tx, err := c.pool.Begin(context.TODO())

	if err != nil {
		return fmt.Errorf("Can`t initialize transaction %v", err)
	}
	defer func() { _ = tx.Rollback(context.TODO()) }()

	res, err := c.pool.Exec(context.TODO(), AddOrderTask, orderID, data)
	if err != nil {
		return fmt.Errorf("Can`t add data to table %v", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("Error : now rows affected")
	}

	if err := tx.Commit(context.TODO()); err != nil {
		return fmt.Errorf("Cant`t commit transaction, %v", err)
	}
	
	return nil
}

func (c *Cache) GetOrders() (map[string][]byte, error) {
	rows, err := c.pool.Query(context.TODO(), GetOrdersTask)
	if err != nil {
		return nil, fmt.Errorf("Can`t get data from database, %v", err)
	}

	defer rows.Close()

	Orders := make(map[string][]byte, 0)

	for rows.Next() {
		var ID string
		var data []byte

		err := rows.Scan(&ID, &data)
		if err != nil {
			return nil, fmt.Errorf("Can`t scan row in database, %v", err)
		}
		Orders[ID] = data
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Can`t scan row in database %v", err)
	}
	return Orders, nil
}
