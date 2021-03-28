package main

import (
	"ecommerce-app/app/helpers"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	orderHttpDelivery "ecommerce-app/orders/delivery/http"
	orderPostgresRepository "ecommerce-app/orders/repository/postgres"
	orderUsecase "ecommerce-app/orders/usecase"
	productHttpDelivery "ecommerce-app/products/delivery/http"
	productPostgresRepository "ecommerce-app/products/repository/postgres"
	productUsecase "ecommerce-app/products/usecase"
)

func main() {
	_ = godotenv.Load()
	help := helpers.New()
	// 1 month max age
	help.SetLogMaxAge(time.Hour * 24 * 30)
	help.InitLogger()
	db := help.InitPostgres()

	timeoutStr := os.Getenv("TIMEOUT")
	if timeoutStr == "" {
		timeoutStr = "5"
	}
	timeout, _ := strconv.Atoi(timeoutStr)
	timeoutDuration := time.Duration(timeout) * time.Second

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.StandardLogger().Writer()
	router := gin.New()
	router.Use(help.CustomRequestLogger())
	//default root endpoint
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello there")
	})
	router.Use(gin.Recovery())

	//setup the layers
	//product entity
	pr := productPostgresRepository.NewPostgresRepository(db, help)
	pu := productUsecase.NewProductUsecase(timeoutDuration, pr, help)
	productHttpDelivery.NewProductHandler(router, pu)
	//order entity
	or := orderPostgresRepository.NewPostgresRepository(db, help)
	ou := orderUsecase.NewOrderUsecase(timeoutDuration, or, pr, help)
	orderHttpDelivery.NewOrderHandler(router, ou)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}
	router.Run(":" + os.Getenv("PORT"))
}
