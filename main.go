package main
import (   

    "github.com/gin-gonic/gin"
    //ginSwagger "github.com/swaggo/gin-swagger"
   // swaggerFiles "github.com/swaggo/files"
    //httpSwagger "github.com/swaggo/http-swagger"
    _ "example/web-service-gin/docs"   	
    "example/web-service-gin/configs"
    "example/web-service-gin/routes"
)


// @title API
// @version 1.0
// @description This is a sample serice for managing albums
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
    
    router := gin.Default()    
   
    //routes
    routes.AulbumRoute(router) //add this
    //run database
    configs.ConnectDB()
    router.Run("localhost:8080")
}