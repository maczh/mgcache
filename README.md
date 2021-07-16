# MgCache 基于github.com/muesli/cache2go封装的本地缓存和redis缓存

## 本地缓存说明
+ 应用内本地缓存，用于缓存数据量不大，可以设置过期时间
+ 无需Redis等外部KV库
+ 多协程安全读写
## Redis缓存说明
+ 需要在mgconfig中配置redis
+ 可以直接缓存对象数据，自动进行序列化和反序列化


## 安装

go get -u github.com/maczh/mgcache

## 本地缓存范例
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
## Redis缓存范例
> yaml配置
```yaml
go:
  config:
    used: redis
```
> 使用范例
```go
    import "github.com/maczh/mgcache"
    ...
    //添加缓存
    mgcache.PutCache("mycache","test",value, 1*time.Minute)
    ...
    //获取缓存
    mgcache.GetCache("mycache","test",&value)
    ...
    //判断缓存是否存在与删除key
    if mgcache.ExistsCache("mycache","test") {
        mgcache.DeleteCache("mycache","test")
    }
    ...
    //清空缓存
    mgcache.ClearCache("mycache")
```