# MgCache 本地缓存，基于github.com/muesli/cache2go封装

## 说明
+ 应用内本地缓存，用于缓存数据量不大，可以设置过期时间
+ 无需Redis等外部KV库
+ 多协程安全读写

## 安装

go get -u github.com/maczh/mgcache

## 范例
```go
    import "github.com/maczh/mgcache"
    ...
    //添加缓存
    mgcache.OnGetCache("mycache").Add("test",value, 1*time.Minute)
    ...
    //获取缓存
    v,cached := mgcache.OnGetCache("mycache").Value("test")
    if cached {
        logs.Debug("Got test cache:{}",v)
    }
    ...
    //判断缓存是否存在与删除key
    if mgcache.OnGetCache("mycache").IsExist("test") {
        mgcache.OnGetCache("mycache").Delete("test")
    }
    ...
    //清空指定缓存表的所有缓存，释放内存空间
    mgcache.OnGetCache("mycache").Clear()

```