package zin

import (
	"github.com/devisettymahidhar315/zin/api"
	"github.com/devisettymahidhar315/zin/multi_cache"
	"github.com/gin-gonic/gin"
)

var length = 2

var cache = multi_cache.NewMultiCache()

func Zin() *gin.Engine {

	//storing the values in the cache
	cache.Set("a", "1", length)
	cache.Set("b", "2", length)

	//intiliazes the restapi with gin
	r := gin.Default()

	r.GET("/:key", api.GetCacheValue)
	r.DELETE("/:key", api.DeleteCacheValue)
	r.POST("/:key/:value", api.SetCacheValue)
	r.GET("/redis/print", api.PrintRedisCache)
	r.GET("/inmemory/print", api.PrintInMemoryCache)
	r.DELETE("/all", api.DeleteAll)

	return r

}
