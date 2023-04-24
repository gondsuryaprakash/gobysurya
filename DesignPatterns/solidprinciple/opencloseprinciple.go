package solidprinciple

// Open Closse principle state that the open for extention but close for modification

type Cache interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

type RedisCache struct{}

func (r *RedisCache) Get(key string) (str string, err error) {
	// Retrieve value from Redis
	return
}

func (r *RedisCache) Set(key, value string) (err error) {
	// Set value in Redis
	return
}

type MemcachedCache struct{}

func (m *MemcachedCache) Get(key string) (str string, err error) {
	// Retrieve value from Memcached
	return
}

func (m *MemcachedCache) Set(key, value string) (err error) {
	// Set value in Memcached
	return
}
