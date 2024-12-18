package PatchF

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
	var outstruct v.Flower
	id := c.Param("id")
	database, err := db.Connect()

	if err != nil {
		log.Println(con.ErrDB, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDB})
		return
	}

	selectId := fmt.Sprintf(`SELECT * FROM "Flowers" WHERE "id" = %s`, id)
	res, err := database.Query(selectId)
	if err != nil {
		log.Println(con.ErrNotConnect, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	if res.Next() {

		err = res.Scan(&outstruct.ID, &outstruct.Name, &outstruct.Quantity, &outstruct.Price, &outstruct.ArrivalDate)
		if err != nil {
			log.Println(con.ErrNotConnect, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
		return
	}

	var instruct v.Flower
	if err := c.ShouldBindJSON(&instruct); err != nil {
		log.Println(con.ErrInternal, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
		return
	}
	if instruct.Name != "" {
		outstruct.Name = instruct.Name
	}
	if instruct.Quantity != 0 {
		outstruct.Quantity = instruct.Quantity
	}
	if instruct.Price != 0 {
		outstruct.Price = instruct.Price
	}
	if instruct.ArrivalDate != "" {
		outstruct.ArrivalDate = instruct.ArrivalDate
	}
	fmt.Println(outstruct)
	fmt.Println(instruct)
	param := fmt.Sprintf(`UPDATE "Flowers" SET "Name" = '%s' , "Quantity" = '%d', "Price" = '%f', "ArrivalDate" = '%s' WHERE "id" = %s`, outstruct.Name, outstruct.Quantity, outstruct.Price, outstruct.ArrivalDate, id)
	_, err = database.Exec(param)
	if err != nil {
		log.Println(con.ErrInternal, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
		return
	}
	if err := database.Close(); err != nil {
		log.Println(con.ErrDBClose, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
		return
	}

	c.IndentedJSON(http.StatusOK, outstruct)
}
