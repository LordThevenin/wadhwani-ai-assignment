package repositories

type IUserCacheRepository interface {
}

type RedisUserCacheRepository struct {
}

func InitRedisUserCacheRepository() *RedisUserCacheRepository {
	return new(RedisUserCacheRepository)
}
