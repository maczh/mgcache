package mgcache

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/maczh/mgconfig"
	"strings"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func PutCache(cacheName, cacheKey string, value interface{}, ttl time.Duration) {
	redis,err := mgconfig.GetRedisConnection()
	if err != nil {
		return
	}
	redis.Set(cacheName+":"+cacheKey, toJSON(value), ttl)
}

func DeleteCache(cacheName, cacheKey string) {
	redis,err := mgconfig.GetRedisConnection()
	if err != nil {
		return
	}
	redis.Del(cacheName + ":" + cacheKey)
}

func GetCache(cacheName, cacheKey string, result interface{}) {
	redis,err := mgconfig.GetRedisConnection()
	if err != nil {
		return
	}
	r := redis.Get(cacheName + ":" + cacheKey).Val()
	fromJSON(r, &result)
}

func ClearCache(cacheName string) {
	redis,err := mgconfig.GetRedisConnection()
	if err != nil {
		return
	}
	keys := redis.Keys(cacheName + ":*").Val()
	if keys != nil && len(keys) > 0 {
		redis.Del(keys...)
	}
}

func ExistsCache(cacheName, cacheKey string) bool {
	redis,err := mgconfig.GetRedisConnection()
	if err != nil {
		return false
	}
	exists, _ := redis.Exists(cacheName + ":" + cacheKey).Result()
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
