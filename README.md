# web-service-gin
Go with Gin Rest API, Swaggo for swagger,  MongoDB for Database operation


Steps to run locally:

  1. Run go get . from the root to fetch all of the dependencies listed in go.mod file

  2. Run go run main.go to boot up the service

What main libary and tool are using:

  1. REST API --- The Gin library is choosen one for this project
  2. Swagger Definition --- Swaggo is choosen one
  3. MongoDB - clusture version is choosen one

In order to regenerate below files,  you can run swag init -g main.go
   1. docs.go
   2. swagger.json
   3. swagger.yaml
   
The project has setup for running debugging with Delve to support setup debugger and conditional debugger   

References and links:

   The project follows below blogs to finish above three goals : REST API, SWagger Definition and MongoDB
   
   1. [Build Rest API](https://go.dev/doc/tutorial/web-service-gin)
   1. [Build SwaggerAPI ](https://www.soberkoder.com/swagger-go-api-swaggo/)
   1. [Connect to MongoDB](https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gin-gonic-version-269m)
