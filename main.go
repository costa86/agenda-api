package main

import (
	"agenda/helpers"
	"agenda/routes"

	"agenda/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	r := gin.Default()

	r.POST("/people", routes.CreatePerson)
	r.GET("/people", routes.GetPeople)
	r.GET("/people/:id", routes.GetPerson)
	r.GET("/slots", routes.GetSlots)
	r.DELETE("/people/:id", routes.DeletePerson)
	r.PUT("/people/:id", routes.UpdatePerson)

	r.Run(":" + helpers.Port)
}
