package router
import (
	fuGet "apiGO/pkg/car/get"
	fuGetID "apiGO/pkg/car/getid"
	fuDelete "apiGO/pkg/car/delete"
	fuPost "apiGO/pkg/car/post"
	fuPut "apiGO/pkg/car/put"
	fuPatch "apiGO/pkg/car/patch"

	fuGet2 "apiGO/pkg/flower/getf"
	fuGetID2 "apiGO/pkg/flower/getidf"
	fuDelete2 "apiGO/pkg/flower/deletef"
	fuPost2 "apiGO/pkg/flower/postf"
	fuPut2 "apiGO/pkg/flower/putf"
	fuPatch2 "apiGO/pkg/flower/patchf"

	fuGet3 "apiGO/pkg/furniture/getfu"
	fuGetID3 "apiGO/pkg/furniture/getidfu"
	fuDelete3 "apiGO/pkg/furniture/deletefu"
	fuPost3 "apiGO/pkg/furniture/postfu"
	fuPut3 "apiGO/pkg/furniture/putfu"
	fuPatch3 "apiGO/pkg/furniture/patchfu"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func Run(){
	router := gin.Default()
	//flowers
	router.GET("/flowers", fuGet2.GetFlowers)
	router.GET("/flowers/:id", fuGetID2.GetFlowersByID)
	router.DELETE("/flowers/:id", fuDelete2.DeletedById)
	router.POST("/flowers", fuPost2.PostFlowers)
	router.PUT("/flowers/:id", fuPut2.PutItem)
	router.PATCH("/flowers/:id", fuPatch2.PatchItem)
	//cars
	router.GET("/cars", fuGet.GetCars)
	router.GET("/cars/:id", fuGetID.GetCarsByID)
	router.DELETE("/cars/:id", fuDelete.DeletedById)
	router.POST("/cars", fuPost.PostCars)
	router.PUT("/cars/:id", fuPut.PutItem)
	router.PATCH("/cars/:id", fuPatch.PatchItem)
	//furniture
	router.GET("/furniture", fuGet3.GetFurnitures)
	router.GET("/furniture/:id", fuGetID3.GetFurnituresByID)
	router.DELETE("/furniture/:id", fuDelete3.DeletedById)
	router.POST("/furniture", fuPost3.PostFurnitures)
	router.PUT("/furniture/:id", fuPut3.PutItem)
	router.PATCH("/furniture/:id", fuPatch3.PatchItem)
	router.Run(":8080")
}