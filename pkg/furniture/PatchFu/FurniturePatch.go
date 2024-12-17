package PatchFu

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
	var outstruct v.Furniture
	id := c.Param("id")
	database, err := db.Connect()

	if err != nil {
		log.Println(con.ErrDB, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDB})
		return
	}
	if err := database.Close(); err != nil {
		log.Println(con.ErrDBClose, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
		return
	}
	selectId := fmt.Sprintf(`SELECT * FROM "Furnitures" WHERE "id" = %s`, id)
	res, err := database.Query(selectId)
	if err != nil {
		log.Println(con.ErrNotConnect, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	if res.Next() {

		err = res.Scan(&outstruct.ID, &outstruct.Name, &outstruct.Manufacturer, &outstruct.Height, &outstruct.Width, &outstruct.Length)
		if err != nil {
			log.Println(con.ErrInternal, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
		return
	}

	var instruct v.Furniture
	if err := c.ShouldBindJSON(&instruct); err != nil {
		log.Println(con.ErrNotConnect, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	if instruct.Name != "" {
		outstruct.Name = instruct.Name
	}
	if instruct.Manufacturer != "" {
		outstruct.Manufacturer = instruct.Manufacturer
	}
	if instruct.Height != 0 {
		outstruct.Height = instruct.Height
	}
	if instruct.Width != 0 {
		outstruct.Width = instruct.Width
	}
	if instruct.Length != 0 {
		outstruct.Length = instruct.Length
	}
	fmt.Println(outstruct)
	fmt.Println(instruct)
	param := fmt.Sprintf(`UPDATE "Furnitures" SET "Name" = '%s' , "Manufacturer" = '%s', "Height" = '%d', "Width" = '%d', "Length" = '%d' WHERE "id" = %s`, outstruct.Name, outstruct.Manufacturer, outstruct.Height, outstruct.Width, outstruct.Length, id)
	_, err = database.Exec(param)
	if err != nil {
		log.Println(con.ErrInternal, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
		return
	}
	c.IndentedJSON(http.StatusOK, outstruct)

}
