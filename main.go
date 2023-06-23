package main

// TODO :
// ENCRYPT AND BCRYPT MASIH GA JELAS NARONYA DIMANA
// SUB untuk JWT masih nyimpen struct user
// ERROR UNTUK : REGISTER DOUBLE, AKSES TANPA JWT (KE /users)
import (
	"fmt"
	"ntika/auth"
	"ntika/handler"
	"ntika/item"
	"ntika/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	DB_USER := os.Getenv("USER")
	DB_PASS := os.Getenv("PASS")
	DB := os.Getenv("DB")
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASS, DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("GAGAl")
	}

	db.AutoMigrate(&auth.User{})
	db.AutoMigrate(&item.Item{})
	db.AutoMigrate(&item.Orders{})
	db.AutoMigrate(&item.OrderItem{})

	userRepository := auth.NewRepo(db)
	userService := auth.NewService(userRepository)
	userHandler := handler.NewHandler(userService)

	itemRepository := item.NewRepo(db)
	itemService := item.NewService(itemRepository)
	itemHandler := handler.NewHandlerItems(itemService, userService)

	r := gin.Default()
	// r.Use(CORSMiddleware())

	v1 := r.Group("/v1").Use(CORSMiddleware())
	// ex := r.Group("/ex").Use(CORSMiddleware2())
	// api := r.Group("/api").Use(CORSMiddleware())
	// v1.Use(CORSMiddleware())

	// dev only
	v1.GET("/users", middleware.RequireAuth, userHandler.GetAllUsers)

	// AUTH
	v1.POST("/sign-up", userHandler.CreateUser)
	v1.POST("/sign-in", userHandler.SignIn)

	// v1.Static("/image", "./static/")

	// GET ALL PRODUCTS
	v1.GET("/items", itemHandler.Catalog)
	// CREATE PRODUCT
	v1.POST("/item", itemHandler.Create)
	// GET USER PROFILE
	v1.GET("/user", middleware.RequireAuth, userHandler.UserProfile)

	v1.POST("/order", middleware.RequireAuth, itemHandler.Order)

	// TODO : harus admin
	v1.PUT("/admin/acc", middleware.RequireAuth, itemHandler.ACC)
	v1.PUT("/admin/update_order", middleware.RequireAuth, itemHandler.UpdateOrder)
	// v1.GET("/user", middleware.RequireAuth, userHandler.Call)
	// v1.POST("/cange_address", userHandler.UpdateAddress)

	// =================================================================================
	// =================================================================================

	// v1.GET("/song/:id", songHandler.GetSongByID)

	// v1.OPTIONS("/song", CORSMiddleware())

	// v1.PUT("/song/:id", songHandler.UpdateSong)

	// v1.DELETE("/song/:id", songHandler.DeleteSong)
	// v1.OPTIONS("/song/:id", CORSMiddleware())
	// // EX-API RELATED
	// api.GET("/yt/:id", YT_TN)

	// // TAG-RELATED===============================================
	// v1.GET("/songs/:tag", tagHandler.GetSongByTag)
	// v1.GET("/tags/:song", tagHandler.GetTagsBySong)
	// v1.GET("/tags", tagHandler.GetAllTags)
	// v1.GET("/tag/:id", tagHandler.GetTagByID)

	// v1.POST("/tag", tagHandler.PostTag)
	// v1.OPTIONS("/tag", CORSMiddleware())
	// // PUT NOT YET?
	// v1.PUT("/tag/:id", tagHandler.UpdateTag)

	// v1.DELETE("/tag/:id", tagHandler.DeleteTag)
	// v1.OPTIONS("/tag/:id", CORSMiddleware())

	// v1.POST("/filtertag", CORSMiddleware(), tagHandler.FilterTag)
	// v1.OPTIONS("/filtertag", CORSMiddleware())

	// // SONG-TAG-RELATED===========================================
	// v1.POST("/songtag", tagHandler.AddTag)
	// v1.OPTIONS("/songtag", CORSMiddleware())

	// v1.DELETE("/songtag/:id", tagHandler.DeleteSongTag)
	// v1.OPTIONS("/songtag/:id", CORSMiddleware())

	// // DEV ONLY ======================================

	// v1.GET("/songtags", tagHandler.GetAllSongTags)

	// //  Extension =================================================
	// ex.GET("/songs", songHandler.GetAllSongs)
	// ex.POST("/song", songHandler.PostSongsHandler)
	// ex.OPTIONS("/song", CORSMiddleware2())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func CORSMiddleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://www.youtube.com")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
