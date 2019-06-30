// blog server main

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"backend/config"
	"backend/model"
	"backend/router"
)

func migrate() {
	var dc = config.AllConfig.Database
	db := model.DB

	options := fmt.Sprintf("ENGINE=InnoDB CHARSET=%s auto_increment=1;", dc.Charset)

	db.Set("gorm:table_options", options).AutoMigrate(&model.Category{})
	db.Set("gorm:table_options", options).AutoMigrate(&model.Article{})
	db.Set("gorm:table_options", options).AutoMigrate(&model.User{})

	//db.AutoMigrate(&model.User{})
	//db.AutoMigrate(&model.Category{})
	//db.AutoMigrate(&model.Article{})

	model.CreateAdminUser()
}

func main() {
	// create table
	migrate()

	// Creates a router without any middleware by default
	app := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	app.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	app.Use(gin.Recovery())

	router.Route(app)

	app.Run(":6060")
}
