package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	config "github.com/mitrasoftware/pureone_backend_go/config"
	"github.com/mitrasoftware/pureone_backend_go/routes"
)

var ginLambda *ginadapter.GinLambdaV2

func init() {
	config.LoadEnvVariables()

	config.ConnectDB()
	config.SyncDatabase()

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.MaxMultipartMemory = 8 << 20

	// r.GET("/", ctrl.GetServices)
	// r.GET("/services", ctrl.GetServices)
	// r.GET("/slider_image", ctrl.GetSliderImages)
	// r.GET("/get_plan", ctrl.GetPlanDetails)
	// r.GET("/categoryIcons", ctrl.CategoryIcons)

	// r.POST("/listProperty", controllers.ListProperty)
	// r.POST("/listPG", controllers.ListPGController)
	// r.POST("/fetch_properties", ctrl.FetchProperties)
	// r.POST("/service_request", controllers.ServiceReqest)
	// r.POST("/add_service_provider", controllers.AddServiceProvider)
	// r.POST("/fetch_service_provider", ctrl.FetchServiceProvider)
	// r.POST("/pgs_list", ctrl.FetchPG)
	// r.POST("/live_search", ctrl.LiveSearch)
	// r.POST("/can_we_make_call", controllers.CanWeAbleToCall)
	// r.POST("/my_listings", ctrl.MyListings)
	// r.POST("/searched_items_product", ctrl.SearchedItemsProduct)

	// r.PUT("/updateCredits", upd.UpdateCredits)
	// r.PUT("/take_plan", upd.TakePlan)

	// r.DELETE("/delist_property", dtl.DelistProperty)

	r = routes.SetupRoutes() // routes.SetupRoutes()

	ginLambda = ginadapter.NewV2(r) //

	r.Run(":3000")

}
func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// Ensure the request path is properly mapped
	if req.RequestContext.HTTP.Path != "" {
		req.RequestContext.HTTP.Path = "/" + req.RequestContext.HTTP.Path
	}
	return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
