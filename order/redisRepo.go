package order

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	Client *redis.Client
}

func orderIDKey(id uint64) string {
	return fmt.Sprintf("order:%d", id)
}

func (r *RedisRepo) Ping(ctx context.Context) error {
	return r.Client.Ping(ctx).Err()
}

func (r *RedisRepo) Close(ctx context.Context) error {
	return r.Client.Close()
}

func (r *RedisRepo) Create(ctx context.Context, order OrderModel) error {
	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("Failed to encode order: %w", err)
	}

	key := orderIDKey(order.OrderID)

	txn := r.Client.TxPipeline()

	err = txn.SetNX(ctx, key, string(data), 0).Err()
	if err != nil {
		txn.Discard()
		return fmt.Errorf("Failed to create order: %w", err)
	}

	err = txn.SAdd(ctx, "orders", key).Err()
	if err != nil {
		txn.Discard()
		return fmt.Errorf("Failed to add to orders: %w", err)
	}

	_, err = txn.Exec(ctx)
	if err != nil {
		return fmt.Errorf("Failed to exec: %w", err)
	}

	return nil
}

var ErrNotExists = errors.New("Order does not exist")

func (r *RedisRepo) FindByID(ctx context.Context, id uint64) (OrderModel, error) {
	key := orderIDKey(id)

	value, err := r.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return OrderModel{}, ErrNotExists
	} else if err != nil {
		return OrderModel{}, fmt.Errorf("get order: %w", err)
	}

	var order OrderModel
	err = json.Unmarshal([]byte(value), &order)
	if err != nil {
		return OrderModel{}, fmt.Errorf("Failed to decode order json: %w", err)
	}

	return order, nil
}

func (r *RedisRepo) DeleteByID(ctx context.Context, id uint64) error {
	key := orderIDKey(id)

	txn := r.Client.TxPipeline()

	_, err := txn.Del(ctx, key).Result()

	if errors.Is(err, redis.Nil) {
		txn.Discard()
		return ErrNotExists
	} else if err != nil {
		txn.Discard()
		return fmt.Errorf("Failed to delete order: %w", err)
	}

	err = txn.SRem(ctx, "orders", key).Err()
	if err != nil {
		txn.Discard()
		return fmt.Errorf("Failed to remove from orders set: %w", err)
	}

	_, err = txn.Exec(ctx)
	if err != nil {
		return fmt.Errorf("Failed to exec %w", err)
	}

	return nil
}

func (r *RedisRepo) UpdateByID(ctx context.Context, order OrderModel) error {
	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("Failed to encode order: %w", err)
	}

	key := orderIDKey(order.OrderID)

	err = r.Client.SetXX(ctx, key, string(data), 0).Err()
	if errors.Is(err, redis.Nil) {
		return ErrNotExists
	} else if err != nil {
		return fmt.Errorf("Failed to delete order: %w", err)
	}

	return nil
}

func (r *RedisRepo) List(ctx context.Context, page FindAllPage) (FindResult, error) {
	res := r.Client.SScan(ctx, "orders", page.Offset, "*", int64(page.Size))

	keys, cursor, err := res.Result()
	if err != nil {
		return FindResult{}, fmt.Errorf("Failed to get order ids: %w", err)
	}

	if len(keys) == 0 {
		return FindResult{
			Orders: []OrderModel{},
		}, nil
	}

	xs, err := r.Client.MGet(ctx, keys...).Result()
	if err != nil {
		return FindResult{}, fmt.Errorf("Failed to get orders: %w", err)
	}

	orders := make([]OrderModel, len(xs))

	for i, x := range xs {
		x := x.(string)
		var order OrderModel

		err := json.Unmarshal([]byte(x), &order)
		if err != nil {
			return FindResult{}, fmt.Errorf("Failed to decode order json: %w", err)
		}
		orders[i] = order
	}

	return FindResult{
		Orders: orders,
		Cursor: cursor,
	}, nil
}
