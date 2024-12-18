package GetIDFu

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

func GetFurnituresByID(c *gin.Context) { //GetID
	slFurnitures := []v.Furniture{}
	id := c.Param("id")
	database, err := db.Connect()

	if err != nil {
		log.Println(con.ErrDB, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDB})
		return
	}

	query := fmt.Sprintf(`SELECT * FROM "Furnitures" WHERE "id" = %s`, id)
	res, err := database.Query(query)
	if err != nil {
		log.Println(con.ErrNotConnect, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	if res.Next() {
		strFurniture := v.Furniture{}

		err = res.Scan(&strFurniture.ID, &strFurniture.Name, &strFurniture.Manufacturer, &strFurniture.Height, &strFurniture.Width, &strFurniture.Length)
		if err != nil {
			log.Println(con.ErrInternal, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
			return
		}
		slFurnitures = append(slFurnitures, strFurniture)
		c.IndentedJSON(http.StatusOK, slFurnitures)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
	}
	if err := database.Close(); err != nil {
		log.Println(con.ErrDBClose, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
		return
	}
}
