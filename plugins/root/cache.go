package root

import (
	"time"

	"github.com/astaxie/beego"
	_cache "github.com/astaxie/beego/cache"
)

// CachePut put
func CachePut(k string, v interface{}, t time.Duration) {
	if err := cache.Put(k, v, t); err != nil {
		beego.Error(err)
	}
}

// CacheGet get
func CacheGet(k string) interface{} {
	return cache.Get(k)
}

var cache _cache.Cache

func init() {
	var err error
	if cache, err = _cache.NewCache(
		beego.AppConfig.String("cacheadapter"),
		beego.AppConfig.String("cachesource"),
	); err != nil {
		beego.Error(err)
	}
}
