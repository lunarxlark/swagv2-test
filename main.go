package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lunarxlark/swaggo-test/controller"
)

//	@title			Swagger Map Example API
//	@version		1.0
//	@termsOfService	http://swagger.io/terms/

// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	r := gin.New()

	c := controller.NewController()

	r.GET("/map", c.GetMap)
	r.StaticFS("/docs", http.Dir("docs"))

	r.Run(":8080")
}
