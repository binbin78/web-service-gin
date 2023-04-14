package routes

import ("github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
	"example/web-service-gin/controllers"

)


func AulbumRoute(router *gin.Engine)  {
    //All routes related to album comes here

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
				"data": "Hello from Gin-gonic & mongoDB",
		})
})	

router.GET("/albums", controllers.GetAlbums)
router.POST("/albums", controllers.CreateAlbum)
router.GET("/albums/:id", controllers.GetAlbumByID)
router.DELETE("/albums/:id", controllers.DeleteAlbumByID)
router.PUT("/albums/:id", controllers.EditAlbumByID)
url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

