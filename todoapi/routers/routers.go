package routers

import (
	"todoapi/controllers"
	"github.com/gin-gonic/gin"
	"todoapi/database"

)

func Setup(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		v1.POST("/items", controllers.CreateItem(database.DBGorm))           // create item
		v1.GET("/items", controllers.GetListOfItems(database.DBGorm))        // list items
		v1.GET("/items/:id", controllers.ReadItemById(database.DBGorm))      // get an item by ID
		v1.PUT("/items/:id", controllers.EditItemById(database.DBGorm))      // edit an item by ID
		v1.DELETE("/items/:id", controllers.DeleteItemById(database.DBGorm)) // delete an item by ID
	}
}
