package routespkg

import (
	"net/http"

	api "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/api"
	corsmiddleware "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/middlewares/cors_middleware"
	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// NewRouter returns a new router.
func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	return NewRouterWithGinEngine(gin.Default(), handleFunctions)
}

// NewRouter add routes to existing gin engine.
func NewRouterWithGinEngine(router *gin.Engine, handleFunctions ApiHandleFunctions) *gin.Engine {
	router.Use(corsmiddleware.CorsMiddleware())
	protected := router.Group("/")
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
		switch route.Name {
		case "GetAllRolesGet":
			protected.GET(route.Pattern, route.HandlerFunc)
		default:
			switch route.Method {
			case http.MethodGet:
				router.GET(route.Pattern, route.HandlerFunc)
			case http.MethodPost:
				router.POST(route.Pattern, route.HandlerFunc)
			case http.MethodPut:
				router.PUT(route.Pattern, route.HandlerFunc)
			case http.MethodPatch:
				router.PATCH(route.Pattern, route.HandlerFunc)
			case http.MethodDelete:
				router.DELETE(route.Pattern, route.HandlerFunc)
			}
		}
	}

	return router
}

// Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandleFunctions struct {

	// Routes for the DefaultAPI part of the API
	DefaultAPI api.DefaultAPI
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{}
}
