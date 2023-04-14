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
