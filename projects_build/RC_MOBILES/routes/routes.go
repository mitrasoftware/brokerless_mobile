package routes

import (
	"github.com/gin-gonic/gin"
	auth "github.com/mitrasoftware/pureone_backend_go/auth"
	ctr "github.com/mitrasoftware/pureone_backend_go/controllers/add"
	fetch "github.com/mitrasoftware/pureone_backend_go/controllers/fetch"
	"github.com/mitrasoftware/pureone_backend_go/middleware"
	// dtl "github.com/mitrasoftware/brokerless/controllers/remove"
	// upd "github.com/mitrasoftware/brokerless/controllers/update"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.MaxMultipartMemory = 8 << 20

	// --- Public routes (no JWT needed) ---
	r.POST("/login", auth.Login)

	// --- Protected routes (JWT required) ---

	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())

	auth.GET("/fetch_category_icons", fetch.CategoryIcons)
	auth.POST("/add_services", ctr.AddCategory)

	auth.POST("/add_products", ctr.AddProducts)
	auth.GET("/get_products", fetch.GetProducts)

	return r
	// }
}

// }
