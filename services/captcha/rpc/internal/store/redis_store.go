package store

import (
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"strings"
)

type RedisStore struct {
	redis  *redis.Redis
	expire int
}

func NewRedisStore(redis *redis.Redis, expire int) *RedisStore {
	// default expire time: 5 minutes
	if expire == -1 || expire == 0 {
		expire = 60 * 5
	}
	return &RedisStore{redis, expire}
}

func (s *RedisStore) Set(id string, value string) {
	err := s.redis.Setex(id, value, s.expire)
	if err != nil {
		logx.Errorf(fmt.Sprintf("RedisStore set %s:%s failed, err: %v", id, value, err))
		return
	}
}

func (s *RedisStore) Get(id string, clear bool) string {
	val, err := s.redis.Get(id)
	if err != nil {
		logx.Errorf(fmt.Sprintf("RedisStore get %s failed, err: %v", id, err))
		return val
	}

	if clear {
		go func() {
			// TODO: need use redis lock?
			//lock := redis.NewRedisLock(s.redis, id)
			//lock.SetExpire(1)
			//_, _ = lock.Acquire()
			//defer func() {
			//	released, err := lock.Release()
			//	if err != nil || !released {
			//		logx.Errorf(fmt.Sprintf("Release Lock %s failed, err: %v", id, err))
			//	}
			//}()
			_, err := s.redis.Del(id)
			if err != nil {
				logx.Errorf(fmt.Sprintf("RedisStore del %s failed, err: %v", id, err))
				return
			}
		}()
	}
	return val
}

func (s *RedisStore) Verify(id, answer string, clear bool) bool {
	ansInRedis := strings.TrimSpace(s.Get(id, clear))
	return ansInRedis == answer
}
