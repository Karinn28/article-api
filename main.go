package main

import (
	"planetku/config"
	"planetku/controller"
	"planetku/repository"
	"planetku/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	// Repo
	articleRepository repository.ArticleRepository = repository.NewArticleRepository(db)


	// Service
	
	articleService service.ArticleService = service.NewArticleService(articleRepository)
	

	// Controller
	articleController controller.ArticleContorller = controller.NewArticleController(articleService)
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "false")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Use(CORSMiddleware())


	
	articleRoutes := r.Group("api")
	{
		// Untuk menampilkan semua data
		articleRoutes.GET("/getAllArticles", articleController.GetAllArticle)
		// Untuk menampilkan artikel terbaru
		articleRoutes.GET("/getLatestArticles", articleController.GetLatestArticles)
		// Untuk output klasifikasi
		articleRoutes.GET("/getArticleWithKey/:title", articleController.GetArticleByKey)
		// Show Article
		articleRoutes.GET("/showArticle/:slug", articleController.ShowArticle)
	}

	r.Run(":8081")
}