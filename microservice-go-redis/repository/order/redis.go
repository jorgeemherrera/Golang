package order

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	Client *redis.Client
}

// Methods
func (r *RedisRepo) Insert(ctx context.Context) {

}
