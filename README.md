# web-service-gin
Go with Gin Rest API, Swaggo for swagger,  MongoDB for Database operation


Steps to run locally:

Run go get . from the root to fetch all of the dependencies listed in go.mod file
Run go run main.go to boot up the service

REST API --- The Gin library is choosen one for this project
Swagger Definition --- Swaggo is choosen one
MongoDB - clusture version is choosen one

In order to regenerate below files,  you can run swag init -g main.go
   docs.go
   swagger.json
   swagger.yaml
   
The project has setup for running debugging with Delve to support setup debugger and conditional debugger   
