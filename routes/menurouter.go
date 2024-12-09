

package routes

import (
       "github.com/gin-gonic/gin"
	   controller "golangrestaurantmanagement/controllers"



)

func MenuRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/menus", controller.GetMenus())
	incomingRoutes.GET("/menus/:menu_id",controller.GetMenus())
	incomingRoutes.POST("/menus", controller.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id", controller.UpdateMenu( ))



}