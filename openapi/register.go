package openapi

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files/v2"
	"github.com/swaggo/swag/v2"
	"net/http"
	"regexp"
)

//go:embed swagger.json
var docs string

// copy from github.com/swaggo/files/v2/dist/swagger-initializer.js
//
//go:embed swagger-initializer.js
var customSwaggerInitializerJsFile string

// CustomSpec 自定义简易Spec实现
type customSpec struct{}

func (s *customSpec) ReadDoc() string {
	return docs
}

func init() {
	swag.Register("swagger", &customSpec{})
}

func RegisterOpenApiRoute(route *gin.Engine) {
	route.GET("/openapi/*any", openapiRequestHandler())
}

func openapiRequestHandler() gin.HandlerFunc {
	var matcher = regexp.MustCompile(`(.*)(docs\.json|index\.html|index\.css|favicon-16x16\.png|favicon-32x32\.png|/oauth2-redirect\.html|swagger-initializer\.js|swagger-ui\.css|swagger-ui\.css\.map|swagger-ui\.js|swagger-ui\.js\.map|swagger-ui-bundle\.js|swagger-ui-bundle\.js\.map|swagger-ui-standalone-preset\.js|swagger-ui-standalone-preset\.js\.map)[?|.]*`)
	return func(ctx *gin.Context) {
		path := "/"
		matches := matcher.FindStringSubmatch(ctx.Request.RequestURI)
		if len(matches) == 3 {
			path = matches[2]
		}
		if matches != nil && len(matches) != 3 {
			ctx.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		switch path {
		case "docs.json":
			readDoc, err := swag.ReadDoc()
			if err != nil {
				ctx.AbortWithStatusJSON(500, err.Error())
				return
			}
			ctx.Header("Content-Type", "application/json; charset=utf-8")
			ctx.String(http.StatusOK, readDoc)
			return
		case "swagger-initializer.js":
			ctx.Header("Content-Type", "application/javascript")
			ctx.String(http.StatusOK, customSwaggerInitializerJsFile)
			return
		default:
			ctx.Request.URL.Path = path
			http.FileServer(http.FS(swaggerFiles.FS)).ServeHTTP(ctx.Writer, ctx.Request)
		}
	}
}
