package abstraction

// GetFunc - function creates new value for key
type GetFunc func(key string) (interface{}, error)

// CacheGetter - cache interface
type CacheGetter interface {
	Get(key string, f GetFunc) (value interface{}, err error)
}
