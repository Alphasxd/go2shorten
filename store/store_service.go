package store

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// StorageService 定义一个结构体，用来存储redis的客户端
type StorageService struct {
	redisClient *redis.Client
}

// 定义一个全局变量，用来存储redis的客户端和上下文
var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

// CacheDuration 定义缓存期限，实际场景中不需要定义，应该使用 LRU 算法
const CacheDuration = 6 * time.Hour

// InitializeStore 初始化 StoreService
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis-11708.c292.ap-southeast-1-1.ec2.cloud.redislabs.com:11708",
		Password: "G6Njk0oCWMqiyQ3FHGUtuTiNpnvmk9j2",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
	fmt.Printf("Connected to Redis: %v\n", pong)
	storeService.redisClient = redisClient
	return storeService
}

// SaveUrlMapping 存储原始URL和短链接的映射
func SaveUrlMapping(shortURL string, originalURL string, userID string) {
	err := storeService.redisClient.Set(ctx, shortURL, originalURL, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortURL, originalURL))
	}
}

// RetrieveInitialUrl 根据短链接获取原始URL
func RetrieveInitialUrl(shortURL string) string {
	result, err := storeService.redisClient.Get(ctx, shortURL).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortURL))
	}
	return result
}
