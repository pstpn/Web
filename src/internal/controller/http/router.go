package http

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"course/docs"
	"course/internal/controller/http/admin"
	"course/internal/controller/http/user"
	"course/internal/service"
	"course/pkg/logger"
)

type Controller struct {
	handler     *gin.Engine
	routerGroup *gin.RouterGroup
}

func NewRouter(handler *gin.Engine) *Controller {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(cors.Default())

	// Swagger v1 settings
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := handler.Group("/api/v1")
	{
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		v1.GET("/healthcheck", healthCheck)
	}

	return &Controller{
		handler:     handler,
		routerGroup: v1,
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

func (c *Controller) SetAuthRoute(l logger.Interface, authService service.AuthService) {
	a := NewAuthController(l, authService)

	c.routerGroup.POST("/register", a.Register)
	c.routerGroup.POST("/login", a.Login)
	c.routerGroup.POST("/refresh", a.RefreshTokens)
}

func (c *Controller) SetInfoCardRoute(
	l logger.Interface,
	infoCardService service.InfoCardService,
	documentService service.DocumentService,
	fieldService service.FieldService,
	checkpointService service.CheckpointService,
	photoService service.PhotoService,
	authService service.AuthService,
) {
	i := admin.NewInfoCardController(
		l,
		infoCardService,
		documentService,
		fieldService,
		checkpointService,
		photoService,
		authService,
	)

	c.routerGroup.GET("/infocards", i.ListFullInfoCards)
	c.routerGroup.GET("/infocards/:id", i.GetFullInfoCard)
	c.routerGroup.PATCH("/infocards/:id", i.ConfirmEmployeeInfoCard)
	c.routerGroup.GET("infocard-photos/:id", i.GetEmployeeInfoCardPhoto)
}

func (c *Controller) SetProfileRoute(
	l logger.Interface,
	infoCardService service.InfoCardService,
	documentService service.DocumentService,
	fieldService service.FieldService,
	authService service.AuthService,
	photoService service.PhotoService,
) {
	p := user.NewProfileController(l, infoCardService, documentService, fieldService, authService, photoService)

	// https://restfulapi.net/resource-naming/#:~:text=than%20one%20archetype.-,2.1.1.%20document,-A%20document%20resource
	c.routerGroup.POST("/profile", p.FillProfile)
	c.routerGroup.GET("/profile", p.GetProfile)
	c.routerGroup.GET("/employee-photo", p.GetEmployeePhoto)
}

func (c *Controller) SetPassageRoute(
	l logger.Interface,
	documentService service.DocumentService,
	checkpointService service.CheckpointService,
	authService service.AuthService,
) {
	p := admin.NewPassageController(l, documentService, checkpointService, authService)

	c.routerGroup.POST("/passages", p.CreatePassage)
}
