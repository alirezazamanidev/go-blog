package routes

import (
	"github.com/alirezazamanidev/go-blog/app/controllers"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)


func SetUp(router *gin.Engine,db *gorm.DB,rdb *redis.Client){

	authController:=controllers.AuthNewController(db,rdb);
	authRoute:=router.Group("auth");
	{
		authRoute.POST("/send-otp",authController.SendOtp);
		authRoute.POST("/check-otp",authController.CheckOtp);
	}
}