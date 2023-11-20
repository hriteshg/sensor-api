package cache

import (
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"sensor-api/pkg/model"
	"time"

	"github.com/gin-gonic/gin"
)

type RedisCache struct {
	Client *redis.Client
}

func NewRedisCache(cacheUrl string) RedisCache {
	return RedisCache{
		Client: redis.NewClient(&redis.Options{
			Addr: cacheUrl,
		})}
}

func (r *RedisCache) GetFromCache(c *gin.Context, groupName string) (*model.SensorGroupAggregate, error) {
	var result model.SensorGroupAggregate
	val, err := r.Client.Get(c, groupName).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(val), &result)
	return &result, err
}

func (r *RedisCache) SetInCache(c *gin.Context, key string, value interface{}) error {
	status := r.Client.Set(c, key, value, 10*time.Second)
	return status.Err()
}
