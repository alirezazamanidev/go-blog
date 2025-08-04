package routes

import (
	"github.com/alirezazamanidev/go-blog/app/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func SetUp(router *gin.Engine,db *gorm.DB){

	authController:=controllers.AuthNewController(db);
	authRoute:=router.Group("auth");
	{
		authRoute.POST("/send-otp",authController.SendOtp);
	}
}