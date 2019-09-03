package cache

import (
	"time"

	cache "github.com/pmylund/go-cache"
)

var (
	mc     = cache.New(2*time.Hour, 600*time.Second)
	ts int = 600
)

func GetCache(key string) interface{} {
	if c, ok := mc.Get(key); ok && c != nil {
		return c
	}
	return nil
}

func SetCache(key string, val interface{}) bool {
	mc.Set(key, val, time.Duration(ts)*time.Second)
	return true
}
