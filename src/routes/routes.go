package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/myperri/copner/src/controllers"
)

func Routes() {
	route := gin.Default()
	// route.Run(":3001")
	route.Use(cors.Default())
	route.Use(static.Serve("/", static.LocalFile("public", false)))
	route.POST("/api/login", controllers.Login)
	route.PUT("/api/actualizar", controllers.Actualizar)
	route.PUT("/api/actualizarT", controllers.ActualizarT)
	route.GET("/api/pedidos", controllers.Pedidos)
	route.GET("/api/domicilios", controllers.Domicilios)
	route.GET("/api/domiciliarios", controllers.Domiciliarios)
	route.GET("/api/novedades", controllers.Novedades)
	route.PUT("/api/aggdomiciliarios", controllers.AggDomiciliarios)
	route.PUT("/api/aggdomiciliariosn2", controllers.AggDomiciliariosn2)
	route.PUT("/api/detalles", controllers.Detalles)
	route.PUT("/api/impresora", controllers.Impresora)
	route.PUT("/api/aggnovedad1", controllers.AggNovedad)
	route.PUT("/api/archivopost", controllers.GenerarArchivoPO1)
	route.GET("/api/pqrs", controllers.Pqrs)
	route.PUT("/api/respuestaPqrs", controllers.RespuestaPqrs)
	route.GET("/api/cuadrecajadomi", controllers.CuadreCajaDomi)
	route.GET("/api/cuadrecajapunto", controllers.CuadreCajaPunto)
	route.Run()
}
