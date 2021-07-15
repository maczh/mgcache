package mgcache

import (
	jsoniter "github.com/json-iterator/go"
	config "github.com/maczh/mgconfig"
	"strings"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func PutCache(cacheName, cacheKey string, value interface{}, ttl time.Duration) {
	config.Redis.Set(cacheName+":"+cacheKey, toJSON(value), ttl)
}

func DeleteCache(cacheName, cacheKey string) {
	config.Redis.Del(cacheName + ":" + cacheKey)
}

func GetCache(cacheName, cacheKey string, result interface{}) {
	r := config.Redis.Get(cacheName + ":" + cacheKey).Val()
	fromJSON(r, &result)
}

func ClearCache(cacheName string) {
	keys := config.Redis.Keys(cacheName + ":*").Val()
	if keys != nil && len(keys) > 0 {
		config.Redis.Del(keys...)
	}
}

func ExistsCache(cacheName, cacheKey string) bool {
	exists, _ := config.Redis.Exists(cacheName + ":" + cacheKey).Result()
	return exists == 1
}

func toJSON(o interface{}) string {
	j, err := json.Marshal(o)
	if err != nil {
		return "{}"
	} else {
		js := string(j)
		js = strings.Replace(js, "\\u003c", "<", -1)
		js = strings.Replace(js, "\\u003e", ">", -1)
		js = strings.Replace(js, "\\u0026", "&", -1)
		return js
	}
}

func fromJSON(j string, o interface{}) *interface{} {
	err := json.Unmarshal([]byte(j), &o)
	if err != nil {
		return nil
	} else {
		return &o
	}
}
