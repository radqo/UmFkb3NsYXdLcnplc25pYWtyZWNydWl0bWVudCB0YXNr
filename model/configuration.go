package model

// AppConfiguration - application configuration
type AppConfiguration struct {
	Port                  string `json:"port"`
	CacheTimeoutInSeconds int    `json:"cacheTimeoutInSeconds"`
	ClientTimeoutInSeconds int    `json:"clientTimeoutInSeconds"`
}