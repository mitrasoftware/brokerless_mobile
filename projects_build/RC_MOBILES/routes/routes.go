package routes

import (
	"github.com/gin-gonic/gin"
	ctr "github.com/mitrasoftware/pureone_backend_go/controllers/add"
	// dtl "github.com/mitrasoftware/brokerless/controllers/remove"
	// upd "github.com/mitrasoftware/brokerless/controllers/update"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.MaxMultipartMemory = 8 << 20

	// api := r.Group("/api")
	// {
	// r.GET("/", ctrl.GetServices)
	// r.GET("/services", ctrl.GetServices)
	// r.GET("/slider_image", ctrl.GetSliderImages)
	// r.GET("/get_plan", ctrl.GetPlanDetails)
	// r.GET("/categoryIcons", ctrl.CategoryIcons)
	r.POST("/list_services", ctr.ListServices)
	// r.POST("/listProperty", controllers.ListProperty)
	// r.POST("/listPG", controllers.ListPGController)
	// r.POST("/fetch_properties", ctrl.FetchProperties)
	// r.POST("/service_request", controllers.ServiceReqest)
	// r.POST("/add_service_provider", controllers.AddServiceProvider)
	// r.POST("/pgs_list", ctrl.FetchPG)

	// r.POST("/can_we_make_call", controllers.CanWeAbleToCall)
	// r.POST("/my_listings", ctrl.MyListings)

	// r.PUT("/updateCredits", upd.UpdateCredits)

	// r.DELETE("/delist_property", dtl.DelistProperty)

	return r
	// }
}

// }
