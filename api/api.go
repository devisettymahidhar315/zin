package api

import (
	"net/http"

	"github.com/devisettymahidhar315/zin/multi_cache"

	"github.com/gin-gonic/gin"
)

const length = 2

var cache = multi_cache.NewMultiCache()

// Endpoint to retrieve a value by key
func GetCacheValue(ctx *gin.Context) {
	k := ctx.Param("key")
	ctx.JSON(http.StatusOK, cache.Get(k))

}

// Endpoint to delete a value by key
func DeleteCacheValue(ctx *gin.Context) {
	//storing the key value
	k := ctx.Param("key")
	//calling the delete method
	cache.Del(k)
}

// Endpoint to set a key-value pair
func SetCacheValue(ctx *gin.Context) {
	//storing the key and value pair
	k := ctx.Param("key")
	v := ctx.Param("value")
	//calling the set methods and sending the key,value and length
	cache.Set(k, v, length)
}

// Endpoint to print the in-memory cache contents
func PrintInMemoryCache(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, cache.Print_in_mem())
}

// Endpoint to print the redis cache contents
func PrintRedisCache(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, cache.Print_redis())
}

// Endpoint to delete entire data
func DeleteAll(ctx *gin.Context) {
	cache.Del_ALL()
}

// func InitializeRoutes() *gin.Engine {

// 	//storing the values in the cache
// 	cache.Set("a", "1", length)
// 	cache.Set("b", "2", length)

// 	//intiliazes the restapi with gin
// 	r := gin.Default()

// 	r.GET("/:key", GetCacheValue)
// 	r.DELETE("/:key", DeleteCacheValue)
// 	r.POST("/:key/:value", SetCacheValue)
// 	r.GET("/redis/print", PrintRedisCache)
// 	r.GET("/inmemory/print", PrintInMemoryCache)
// 	r.DELETE("/all", DeleteAll)

// 	return r
// }
