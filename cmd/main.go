package main

import (
	"github.com/DamThiLanAnh/Ecommerce-projcet/controllers"
	"github.com/DamThiLanAnh/Ecommerce-projcet/database"
	"github.com/DamThiLanAnh/Ecommerce-projcet/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Product"), database.UserData(database.Client, "User"))

	router := gin.New()
	router.Use(gin.Logger())

	router.User(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
