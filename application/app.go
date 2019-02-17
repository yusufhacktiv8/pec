package application

import (
	"fmt"
	"log"

	"time"

	"github.com/appleboy/gofight"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// Postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"pecrsuh.id/pec/models"
)

// App struct used to store DB connectiona and router
type App struct {
	Router  *gin.Engine
	DB      *gorm.DB
	GoFight *gofight.RequestConfig
}

// InitRoutes with gin config
func (a *App) InitRoutes() {
	a.Router = gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	a.Router.Use(cors.New(config))
}

// InitDatabase with retry connect
func (a *App) InitDatabase(user, password, dbname string) {
	const maxRetryConnect = 5
	connectionString := fmt.Sprintf("user=%s dbname=%s host=localhost port=5432 sslmode=disable", user, dbname)
	var err error
	for i := 0; i < maxRetryConnect; i++ {
		a.DB, err = gorm.Open("postgres", connectionString) // gorm checks Ping on Open
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatal(err)
		fmt.Printf("Err: " + err.Error())
	}

	a.DB.AutoMigrate(&models.Role{})
}
