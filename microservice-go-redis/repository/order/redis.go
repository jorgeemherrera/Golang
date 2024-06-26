package order

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"microservice-go-redis/model"

	"github.com/redis/go-redis/v9"
)

// Type pointer of redis client
type RedisRepo struct {
	Client *redis.Client
}

// Methods

func orderIDKey(id uint64) string {
	return fmt.Sprintf("order:%d", id)
}

func (r *RedisRepo) Insert(ctx context.Context, order model.Order) error {

	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("Failed to encode order: %w", err)
	}

	key := orderIDKey(order.OrderID)

	// Begin the transaction for us
	txn := r.Client.TxPipeline()

	//SetNX: not exist
	response := txn.SetNX(ctx, key, string(data), 0)

	if err := response.Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("Failed to insert order: %w", err)
	}

	if err := txn.SAdd(ctx, "orders", key).Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("Failed to add order to set: %w", err)
	}

	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("Failed to exec transaction: %w", err)
	}

	return nil
}

var ErrNotExist = errors.New("Order does not exist")

func (r *RedisRepo) FindByID(ctx context.Context, id uint64) (model.Order, error) {
	key := orderIDKey(id)

	value, err := r.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return model.Order{}, ErrNotExist
	} else if err != nil {
		return model.Order{}, fmt.Errorf("get order: %w", err)
	}

	var order model.Order
	err = json.Unmarshal([]byte(value), &order)
	if err != nil {
		return model.Order{}, fmt.Errorf("Failed to decode order: %w", err)
	}
	return order, nil
}
func (r *RedisRepo) DeleteByID(ctx context.Context, id uint64) error {
	key := fmt.Sprintf("order: %d", id)
	fmt.Println("Attempting to delete key:", key)

	deleted, err := r.Client.Del(ctx, key).Result()
	if err != nil {
		fmt.Println("Error deleting key from Redis:", err)
		return err
	}
	if deleted == 0 {
		fmt.Println("Key does not exist:", key)
		return ErrNotExist
	}

	fmt.Println("Key deleted successfully:", key)
	return nil
}

func (r *RedisRepo) Update(ctx context.Context, order model.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("Failed to encode order: %w", err)
	}

	key := orderIDKey(order.OrderID)
	//SetXX: set a value if it already exist
	err = r.Client.SetXX(ctx, key, string(data), 0).Err()
	if errors.Is(err, redis.Nil) {
		return ErrNotExist
	} else if err != nil {
		return fmt.Errorf("set order: %w", err)
	}
	return nil
}

type FindAllPage struct {
	Size   uint
	Offset uint
}

type FindResult struct {
	Orders []model.Order
	Cursor uint64
}

func (r *RedisRepo) FindAll(ctx context.Context, page FindAllPage) (FindResult, error) {
	response := r.Client.SScan(ctx, "orders", uint64(page.Offset), "*", int64(page.Size))

	keys, cursor, err := response.Result()
	if err != nil {
		return FindResult{}, fmt.Errorf("Failed to get order ids: %w", err)
	}

	if len(keys) == 0 {
		return FindResult{
			Orders: []model.Order{},
		}, nil
	}

	xs, err := r.Client.MGet(ctx, keys...).Result()
	if err != nil {
		return FindResult{}, fmt.Errorf("Failed to get orders: %w", err)
	}

	orders := make([]model.Order, len(xs))

	for i, x := range xs {
		x := x.(string)
		var order model.Order

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
