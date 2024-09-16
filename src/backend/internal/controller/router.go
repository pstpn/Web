package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docsv1 "course/docs/v1"
	docsv2 "course/docs/v2"
	graphqlv1 "course/internal/controller/v1/graphql"
	routesv1 "course/internal/controller/v1/http"
	httputils "course/internal/controller/v1/http/utils"
	graphqlv2 "course/internal/controller/v2/graphql"
	routesv2 "course/internal/controller/v2/http"
	"course/internal/service"
	"course/pkg/logger"
)

type Controller struct {
	handler      *gin.Engine
	routerGroups map[string]*gin.RouterGroup
}

func NewRouter(handler *gin.Engine) *Controller {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Disable CORS
	handler.OPTIONS("/*any", httputils.DisableCors)

	// GraphQL API
	graphqlv1.Handle(handler)
	graphqlv2.Handle(handler)

	// Swagger settings
	docsv1.SwaggerInfov1.BasePath = "/api/v1"
	docsv2.SwaggerInfov2.BasePath = "/api/v2"

	v1 := handler.Group("/api/v1")
	{
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v1")))
		v1.GET("/healthcheck", healthCheck)
	}
	v2 := handler.Group("/api/v2")
	{
		v2.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v2")))
		v2.GET("/healthcheck", healthCheck)
	}

	return &Controller{
		handler: handler,
		routerGroups: map[string]*gin.RouterGroup{
			"v1": v1,
			"v2": v2,
		},
	}
}

// healthCheck godoc
//
//	@Summary		Проверка здоровья
//	@Description	Проверка на жизнеспособность
//	@Tags			system
//	@Success		200	{string} string "Сервис жив"
//	@Failure		404	"Сервис мертв"
//	@Router			/healthcheck [get]
func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, time.Now().String())
}

func (c *Controller) SetV1Routes(
	l logger.Interface,
	infoCardService service.InfoCardService,
	documentService service.DocumentService,
	fieldService service.FieldService,
	checkpointService service.CheckpointService,
	photoService service.PhotoService,
	authService service.AuthService,
) {
	routesv1.SetRoutes(
		c.routerGroups["v1"],
		l,
		infoCardService,
		documentService,
		fieldService,
		checkpointService,
		photoService,
		authService,
	)
}

func (c *Controller) SetV2Routes(
	l logger.Interface,
	infoCardService service.InfoCardService,
	documentService service.DocumentService,
	fieldService service.FieldService,
	checkpointService service.CheckpointService,
	photoService service.PhotoService,
	authService service.AuthService,
) {
	routesv2.SetRoutes(
		c.routerGroups["v2"],
		l,
		infoCardService,
		documentService,
		fieldService,
		checkpointService,
		photoService,
		authService,
	)
}
