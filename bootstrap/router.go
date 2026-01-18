package bootstrap

import (
	imagehandler "github.com/BroMikey/goBlog/internal/handler/image"
	settingshandler "github.com/BroMikey/goBlog/internal/handler/settings"
	imagesvc "github.com/BroMikey/goBlog/internal/service/image"
	settingssvc "github.com/BroMikey/goBlog/internal/service/settings"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitRouter(conf *Config, db *gorm.DB, log *logrus.Logger) *gin.Engine {
	setGinMode(conf.System.Env)

	r := gin.Default()

	api := r.Group("/api/v1")

	settingsService := settingssvc.NewService(db, log)
	settingsHandler := settingshandler.NewHandler(settingsService)
	settingshandler.RegisterRoutes(api, settingsHandler)

	imageService := imagesvc.NewService(db, log)
	imageHandler := imagehandler.NewHandler(imageService)
	imagehandler.RegisterRoutes(api, imageHandler)

	return r
}

func setGinMode(env string) {
	switch env {
	case "dev", "debug":
		gin.SetMode(gin.DebugMode)
	case "prod", "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}
