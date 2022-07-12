package service

type CaptchaCacheInterface interface {
	Get(key string) string
	Set(key string, val string, expiresInSeconds int)
	Delete(key string)
	Exists(key string) bool
	GetType() string
	Increment(key string, val int) int
}
