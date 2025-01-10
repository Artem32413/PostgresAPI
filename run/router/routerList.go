package router

import (
	fuDelete "apiGO/pkg/car/Delete"
	fuGet "apiGO/pkg/car/Get"
	fuGetID "apiGO/pkg/car/GetID"
	fuPatch "apiGO/pkg/car/Patch"
	fuPost "apiGO/pkg/car/Post"
	fuPut "apiGO/pkg/car/Put"

	fuDelete2 "apiGO/pkg/flower/DeleteF"
	fuGet2 "apiGO/pkg/flower/GetF"
	fuGetID2 "apiGO/pkg/flower/GetIDF"
	fuPatch2 "apiGO/pkg/flower/PatchF"
	fuPost2 "apiGO/pkg/flower/PostF"
	fuPut2 "apiGO/pkg/flower/PutF"

	fuDelete3 "apiGO/pkg/furniture/DeleteFu"
	fuGet3 "apiGO/pkg/furniture/GetFu"
	fuGetID3 "apiGO/pkg/furniture/GetIDFu"
	fuPatch3 "apiGO/pkg/furniture/PatchFu"
	fuPost3 "apiGO/pkg/furniture/PostFu"
	fuPut3 "apiGO/pkg/furniture/PutFu"

	"github.com/gin-gonic/gin"
)

func StartFlowers() {
	router := gin.Default()
	log.Println("Цветок ввввввввввввввввввввввввввввввввввввввввввввввввввввввввввввввввв")
	router.GET("/flowers/", fuGet2.GetFlowers)
	router.GET("/flowers/:id", fuGetID2.GetFlowersByID)
	router.DELETE("/flowers/:id", fuDelete2.DeletedById)
	router.POST("/flowers/", fuPost2.PostFlowers)
	router.PUT("/flowers/:id", fuPut2.PutItem)
	router.PATCH("/flowers/:id", fuPatch2.PatchItem)
}
func StartCars() {
	router := gin.Default()
	log.Println("Машина ввввввввввввввввввввввввввввввввввввввввввввввввввввввввввввввввв")
	router.GET("/cars/", fuGet.GetCars)
	router.GET("/cars/:id", fuGetID.GetCarsByID)
	router.DELETE("/cars/:id", fuDelete.DeletedById)
	router.POST("/cars/", fuPost.PostCars)
	router.PUT("/cars/:id", fuPut.PutItem)
	router.PATCH("/cars/:id", fuPatch.PatchItem)
}
func StartFurniture() {
	router := gin.Default()
	log.Println("Мебель ввввввввввввввввввввввввввввввввввввввввввввввввввввввввввввввввв")
	router.GET("/furniture/", fuGet3.GetFurnitures)
	router.GET("/furniture/:id", fuGetID3.GetFurnituresByID)
	router.DELETE("/furniture/:id", fuDelete3.DeletedById)
	router.POST("/furniture/", fuPost3.PostFurnitures)
	router.PUT("/furniture/:id", fuPut3.PutItem)
	router.PATCH("/furniture/:id", fuPatch3.PatchItem)
}