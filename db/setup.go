package db

import (
	"context"
	"fmt"
	"log"

	"github.com/alirezazamanidev/go-blog/app/configs"
	"github.com/alirezazamanidev/go-blog/db/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetUp(config *configs.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DbUser, // نام کاربری
		config.DbPass, // رمز عبور
		config.DbHost, // هاست ← این قسمت درست شد
		config.DbPort, // پورت
		config.DbName, // نام دیتابیس
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	db.AutoMigrate(
		&models.User{},
	)
	log.Println("✅ Database connected successfully")
	return db
}
func SetUpRedis(config *configs.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("❌ Failed to connect to Redis: %v", err)
	}
	log.Println("✅ Redis connected successfully")
	return rdb
}
