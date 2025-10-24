package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	config "github.com/mitrasoftware/pureone_backend_go/config"

	_ "github.com/mitrasoftware/pureone_backend_go/docs" // import generated docs
	"github.com/mitrasoftware/pureone_backend_go/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var ginLambda *ginadapter.GinLambdaV2

// @title PureOne Backend API
// @version 1.0
// @description API documentation for PureOne backend
// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func init() {
	config.LoadEnvVariables()

	config.ConnectDB()
	config.SyncDatabase()

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.MaxMultipartMemory = 8 << 20

	r = routes.SetupRoutes() // routes.SetupRoutes()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ginLambda = ginadapter.NewV2(r) //

	r.Run()

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
