package abstraction

// GetFunc - function creates new value for key
type GetFunc func(key string) (interface{}, error)

// CacheOperator - cache interface
type CacheOperator interface {
	Get(key string) (value interface{}, found bool)
	Set(key string, value interface{})
}
