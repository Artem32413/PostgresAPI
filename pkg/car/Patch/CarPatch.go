package Patch

import (
	con "apiGO/run/constLog"
	db "apiGO/run/postgres"
	v "apiGO/structFile"

	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func PatchItem(c *gin.Context) { //Patch
	var outstruct v.Car
	id := c.Param("id")
	database, err := db.Connect()

	if err != nil {
		log.Println(con.ErrDB, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDB})
		return
	}

	selectId := fmt.Sprintf(`SELECT * FROM "Cars" WHERE "id" = %s`, id)
	res, err := database.Query(selectId)
	if err != nil {
		log.Println(con.ErrNotConnect, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	if res.Next() {

		err = res.Scan(&outstruct.ID, &outstruct.Brand, &outstruct.Model, &outstruct.Mileage, &outstruct.Owners)
		if err != nil {
			log.Println(con.ErrInternal, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}

	var instruct v.Car
	if err := c.ShouldBindJSON(&instruct); err != nil {
		log.Println(con.ErrInternal, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": con.ErrInvalidData})
		return
	}
	if instruct.Brand != "" {
		outstruct.Brand = instruct.Brand
	}
	if instruct.Model != "" {
		outstruct.Model = instruct.Model
	}
	if instruct.Mileage != 0 {
		outstruct.Mileage = instruct.Mileage
	}
	if instruct.Owners != 0 {
		outstruct.Owners = instruct.Owners
	}
	fmt.Println(outstruct)
	fmt.Println(instruct)
	param := fmt.Sprintf(`UPDATE "Cars" SET "Brand" = '%s' , "Model" = '%s', "Mileage" = '%d', "Owners" = '%d' WHERE "id" = %s`, outstruct.Brand, outstruct.Model, outstruct.Mileage, outstruct.Owners, id)
	_, err = database.Exec(param)
	if err != nil {
		log.Println(con.ErrNotFound, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	c.IndentedJSON(http.StatusOK, outstruct)
	if err := database.Close(); err != nil {
		log.Println(con.ErrDBClose, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
		return
	}
}
