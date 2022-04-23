package ports

import "time"

type RedisRepository interface {
	Get(key string, dest interface{}) error
	Set(key string, value interface{}, ttl time.Duration) error
	GetBool(key string) bool
	Exists(key string) bool
	Del(key string) error
}
