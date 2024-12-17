package GetFu

import (
	con "apiGO/run/constLog"
	db "apiGO/run/postgres"
	v "apiGO/structFile"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func GetFurnitures(c *gin.Context) { //Get
	slFurniture := []v.Furniture{}
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
	res, err := database.Query(`SELECT * FROM "Furnitures"`)
	if err != nil {
		log.Println(con.ErrNotConnect, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	for res.Next() {
		strFurniture := v.Furniture{}
		err = res.Scan(&strFurniture.ID, &strFurniture.Name, &strFurniture.Manufacturer, &strFurniture.Height, &strFurniture.Width, &strFurniture.Length)
		if err != nil {
			log.Println(con.ErrInternal, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
			return
		}
		slFurniture = append(slFurniture, strFurniture)
	}

	c.IndentedJSON(http.StatusOK, slFurniture)
}
