package repositories

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"user-service/constants"
	"user-service/db"
	"user-service/models"
	"user-service/utils"
)

type IUserCacheRepository interface {
	Set(key string, val models.User)
	Get(key string) (models.User, error)
}

type RedisUserCacheRepository struct {
	cache *redis.Client
}

func InitRedisUserCacheRepository() *RedisUserCacheRepository {
	redisClient := new(RedisUserCacheRepository)
	redisClient.cache = db.GetRedisCache()
	return redisClient
}

func (r *RedisUserCacheRepository) Get(key string) (user models.User, err error) {
	ctx := context.Background()
	resp, err := r.cache.Get(ctx, key).Bytes()
	if err != nil {
		// Log cache miss
		utils.Logger().Debugf("RedisUserCacheRepository: failed to get value of key: %s,with error: %s", key, err.Error())
		return
	}
	err = json.Unmarshal(resp, &user)
	if err != nil {
		// Log unmarshalling error
		return
	}
	return
}

func (r *RedisUserCacheRepository) Set(key string, user models.User) {
	ctx := context.Background()
	userValue, err := json.Marshal(user)
	if err != nil {
		// Log marshalling error
		utils.Logger().Debugf("RedisUserCacheRepository: failed to marshal user with error: %s", err.Error())
		return
	}
	r.cache.Set(ctx, key, userValue, constants.TTL)
}
