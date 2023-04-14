package controllers

import (
	"context"
    "example/web-service-gin/configs"
    "example/web-service-gin/models"
    "example/web-service-gin/responses"
    "net/http"
    "time"
	_ "example/web-service-gin/docs"   	
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var albumCollection *mongo.Collection = configs.GetCollection(configs.DB, "albums")
var validate = validator.New()

// CreateAlbum godoc
// @Summary Create a new Album
// @Description Create a new Album with the input paylod
// @Tags albums
// @Param album-model body models.Album true "Album details"
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /albums [post]
func CreateAlbum (c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var album models.Album
        defer cancel()

        //validate the request body
        if err := c.BindJSON(&album); err != nil {
			 c.JSON(http.StatusBadRequest, responses.AlbumReponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //use the validator library to validate required fields
        if validationErr := validate.Struct(&album); validationErr != nil {
             c.JSON(http.StatusBadRequest, responses.AlbumReponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }

        newAlbum := models.Album{
            ID:       primitive.NewObjectID(),
            Title:    album.Title,
            Artist:   album.Artist,
            Price:    album.Price,
        }

        result, err := albumCollection.InsertOne(ctx, newAlbum)
        if err != nil {
			 c.JSON(http.StatusInternalServerError, responses.AlbumReponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

         c.JSON(http.StatusCreated, responses.AlbumReponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
    }
// GetAlbums godoc
// @Summary Get details of all albums
// @Description Get details of all albums
// @Tags albums
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /albums [get]
func GetAlbums(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var albums []models.Album
	defer cancel()

	//find all items from data base
	results, err := albumCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.AlbumReponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleAlbum models.Album
		if err = results.Decode(&singleAlbum); err != nil {
			c.JSON(http.StatusInternalServerError, responses.AlbumReponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}

		albums = append(albums, singleAlbum)
	}
     
	c.JSON(http.StatusOK,
		responses.AlbumReponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": albums}},
	)
}
// @Summary Get details for a given albumId
// @Description Get details of album corresponding to the input albumId
// @Tags albums
// @Accept  */*
// @Produce  json
// @Param id path string true "ID of the album"
// @Success 200 {object} map[string]interface{}
// @Router /albums/{id} [get]
func GetAlbumByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	albumId := c.Param("id")
	var album models.Album
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(albumId)

	//find the item from data base
	err := albumCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&album)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.AlbumReponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusOK, responses.AlbumReponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": album}})
}
//	@Summary		Delete an album
//	@Description	Delete by album ID
//	@Tags			album
//	@Accept			*/*
//	@Produce		json
//	@Param			id	path		string	true	"album ID"	
//	@Success		204	{object} map[string]interface{}
//	@Router			/albums/{id} [delete]
func DeleteAlbumByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	albumId := c.Param("id")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(albumId)
    //delete from database
	result, err := albumCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.AlbumReponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusNotFound,
			responses.AlbumReponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Album with specified ID not found!"}},
		)
		return
	}

	c.JSON(http.StatusOK,
		responses.AlbumReponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Album successfully deleted!"}},
	)
}
//	@Summary	Add a new albume to the store
//	@Tags		albums
//	@Accept		*/*
//	@Produce	json
//	@Param		message	body		models.Album		true	"Album Data"
//	@Success	204	{object} map[string]interface{}
//	@Router		/albums/album/ [put]
func EditAlbumByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	albumId := c.Param("id")
	var album models.Album
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(albumId)

	//validate the request body
	if err := c.BindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, responses.AlbumReponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&album); validationErr != nil {
		c.JSON(http.StatusBadRequest, responses.AlbumReponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	update := bson.M{"artist": album.Artist, "title": album.Title, "price": album.Price}
	//update data base
	result, err := albumCollection.UpdateOne(ctx, bson.M{"id": albumId}, bson.M{"$set": update})
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.AlbumReponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	//get updated album details
	var updatedAlbum models.Album
	if result.MatchedCount == 1 {
		err := albumCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedAlbum)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AlbumReponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
	}

	c.JSON(http.StatusOK, responses.AlbumReponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedAlbum}})
}
