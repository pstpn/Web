package http

import (
	"github.com/gin-gonic/gin"

	"course/internal/controller/v2/http/admin"
	"course/internal/controller/v2/http/auth"
	"course/internal/controller/v2/http/user"
	"course/internal/service"
	"course/pkg/logger"
)

func setAuthRoute(handler *gin.RouterGroup, l logger.Interface, authService service.AuthService) {
	a := auth.NewAuthController(l, authService)

	handler.POST("/register", a.Register)
	handler.POST("/login", a.Login)
	handler.POST("/refresh", a.RefreshTokens)
}

func setInfoCardRoute(
	handler *gin.RouterGroup,
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

	handler.GET("/infocards", i.ListFullInfoCards)
	handler.GET("/infocards/:id", i.GetFullInfoCard)
	handler.PATCH("/infocards/:id", i.ConfirmEmployeeInfoCard)
	handler.GET("infocard-photos/:id", i.GetEmployeeInfoCardPhoto)
}

func setProfileRoute(
	handler *gin.RouterGroup,
	l logger.Interface,
	infoCardService service.InfoCardService,
	documentService service.DocumentService,
	fieldService service.FieldService,
	authService service.AuthService,
	photoService service.PhotoService,
) {
	p := user.NewProfileController(l, infoCardService, documentService, fieldService, authService, photoService)

	// https://restfulapi.net/resource-naming/#:~:text=than%20one%20archetype.-,2.1.1.%20document,-A%20document%20resource
	handler.POST("/profile", p.FillProfile)
	handler.GET("/profile", p.GetProfile)
	handler.GET("/employee-photo", p.GetEmployeePhoto)
}

func setPassageRoute(
	handler *gin.RouterGroup,
	l logger.Interface,
	documentService service.DocumentService,
	checkpointService service.CheckpointService,
	authService service.AuthService,
) {
	p := admin.NewPassageController(l, documentService, checkpointService, authService)

	handler.POST("/passages", p.CreatePassage)
}

func SetRoutes(
	handler *gin.RouterGroup,
	l logger.Interface,
	infoCardService service.InfoCardService,
	documentService service.DocumentService,
	fieldService service.FieldService,
	checkpointService service.CheckpointService,
	photoService service.PhotoService,
	authService service.AuthService,
) {
	setAuthRoute(handler, l, authService)
	setInfoCardRoute(
		handler,
		l,
		infoCardService,
		documentService,
		fieldService,
		checkpointService,
		photoService,
		authService,
	)
	setProfileRoute(
		handler,
		l,
		infoCardService,
		documentService,
		fieldService,
		authService,
		photoService,
	)
	setPassageRoute(
		handler,
		l,
		documentService,
		checkpointService,
		authService,
	)
}
