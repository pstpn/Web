package user

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	httputils "course/internal/controller/v1/http/utils"
	"course/internal/model"
	"course/internal/service"
	"course/internal/service/dto"
	"course/pkg/logger"
)

const multiFormSizeDefault = 10000000

type ProfileController struct {
	l               logger.Interface
	infoCardService service.InfoCardService
	documentService service.DocumentService
	fieldService    service.FieldService
	authService     service.AuthService
	photoService    service.PhotoService
}

func NewProfileController(
	l logger.Interface,
	infoCardService service.InfoCardService,
	documentService service.DocumentService,
	fieldService service.FieldService,
	authService service.AuthService,
	photoService service.PhotoService,
) *ProfileController {
	return &ProfileController{
		l:               l,
		infoCardService: infoCardService,
		documentService: documentService,
		fieldService:    fieldService,
		authService:     authService,
		photoService:    photoService,
	}
}

type fillProfileRequest struct {
	DocumentSerialNumber string `json:"serialNumber"`
	DocumentType         string `json:"documentType"`
	DocumentFields       string `json:"documentFields"`
}

// https://restfulapi.net/resource-naming/#:~:text=than%20one%20archetype.-,2.1.1.%20document,-A%20document%20resource

// FillProfile godoc
//
//	@Summary		Заполнение профиля
//	@Description	Метод для заполнения профиля
//	@Tags			employee
//	@Accept			mpfd
//	@Accept			json
//	@Param			profileData	formData	fillProfileRequest				true	"Заполнения профиля пользователя"
//	@Param			image		formData	file							true	"Заполнения профиля пользователя"
//	@Success		201			{string}	string							"Профиль успешно создан"
//	@Failure		400			{object}	http.StatusBadRequest			"Некорректное тело запроса"
//	@Failure		401			{object}	http.StatusUnauthorized			"Авторизация неуспешна"
//	@Failure		500			{object}	http.StatusInternalServerError	"Внутренняя ошибка заполнения профиля"
//	@Security		BearerAuth
//	@Router			/profile [post]
func (p *ProfileController) FillProfile(c *gin.Context) {
	payload, err := httputils.VerifyAccessToken(c, p.l, p.authService)
	if err != nil {
		return
	}
	infoCardID, err := payload.GetInfoCardID()
	if err != nil {
		p.l.Errorf("failed to parse infoCard id from payload: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	err = c.Request.ParseMultipartForm(multiFormSizeDefault)
	if err != nil {
		p.l.Errorf("failed to parse multipart form: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect request"})
		return
	}

	err = c.Request.ParseForm()
	if err != nil {
		p.l.Errorf("failed to read profile data: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect profile data"})
		return
	}
	data := c.Request.Form

	req := fillProfileRequest{
		DocumentSerialNumber: data.Get("serialNumber"),
		DocumentType:         data.Get("documentType"),
		DocumentFields:       data.Get("documentFields"),
	}
	if req.DocumentFields == "" || req.DocumentType == "" || req.DocumentSerialNumber == "" {
		p.l.Errorf("failed to decode profile data: empty field")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect profile data"})
		return
	}

	document, err := p.documentService.CreateDocument(c.Request.Context(), &dto.CreateDocumentRequest{
		SerialNumber: req.DocumentSerialNumber,
		InfoCardID:   infoCardID,
		DocumentType: model.ToDocumentTypeFromString(req.DocumentType).Int(),
	})
	if err != nil {
		p.l.Errorf("failed to create document: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create document"})
		return
	}

	for _, newField := range strings.Split(req.DocumentFields, ";") {
		parsedField := strings.Split(newField, ",")
		_, err = p.fieldService.CreateDocumentField(c.Request.Context(), &dto.CreateDocumentFieldRequest{
			DocumentID: document.ID.Int(),
			Value:      parsedField[1],
			Type:       model.ToFieldTypeFromString(parsedField[0]).Int(),
		})
		if err != nil {
			p.l.Errorf("failed to create document field: %s", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create document field"})
			return
		}
	}

	f, err := c.FormFile("image")
	if err != nil {
		p.l.Errorf("failed to get image from form file: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can`t get image from form file"})
		return
	}
	photo, err := f.Open()
	if err != nil {
		p.l.Errorf("failed to open form file: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Can`t open form file"})
		return
	}
	defer photo.Close()

	photoData, err := io.ReadAll(photo)
	if err != nil {
		p.l.Errorf("failed to read photo data: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can`t read photo data"})
		return
	}

	_, err = p.photoService.CreatePhoto(c.Request.Context(), &dto.CreatePhotoRequest{
		DocumentID: document.ID.Int(),
		Data:       photoData,
	})
	if err != nil {
		p.l.Errorf("failed to create photo: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	c.Status(http.StatusCreated)
}

type getProfileResponse struct {
	IsConfirmed    bool                   `json:"isConfirmed"`
	CreatedAt      string                 `json:"createdAt"`
	DocumentType   string                 `json:"documentType"`
	SerialNumber   string                 `json:"serialNumber"`
	DocumentFields []*model.KeyValueField `json:"documentFields"`
}

// GetProfile godoc
//
//	@Summary		Получение профиля
//	@Description	Метод для получения профиля
//	@Tags			employee
//	@Success		200	{object}	getProfileResponse				"Профиль успешно получен"
//	@Failure		400	{object}	http.StatusBadRequest			"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusUnauthorized			"Авторизация неуспешна"
//	@Failure		404	{object}	http.StatusNotFound				"Профиль не найден"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка заполнения профиля"
//	@Security		BearerAuth
//	@Router			/profile [get]
func (p *ProfileController) GetProfile(c *gin.Context) {
	payload, err := httputils.VerifyAccessToken(c, p.l, p.authService)
	if err != nil {
		return
	}
	infoCardID, err := payload.GetInfoCardID()
	if err != nil {
		p.l.Errorf("failed to parse infoCard id from payload: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	infoCard, err := p.infoCardService.GetInfoCard(c.Request.Context(), &dto.GetInfoCardByIDRequest{InfoCardID: infoCardID})
	if err != nil {
		p.l.Errorf("failed to get infoCard: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get info card"})
		return
	}

	document, err := p.documentService.GetDocumentByInfoCard(c.Request.Context(), &dto.GetDocumentByInfoCardIDRequest{
		InfoCardID: infoCardID,
	})
	if err != nil {
		p.l.Errorf("failed to get document by infoCard ID: %s", err.Error())

		status := http.StatusInternalServerError
		if errors.Is(err, pgx.ErrNoRows) {
			status = http.StatusNotFound
		}
		c.AbortWithStatusJSON(status, gin.H{"error": "Failed to get info card document"})
		return
	}

	documentFields, err := p.fieldService.ListDocumentFields(c.Request.Context(), &dto.ListDocumentFieldsRequest{
		DocumentID: document.ID.Int(),
	})
	if err != nil {
		p.l.Errorf("failed to list document fields: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list document fields"})
		return
	}

	c.JSON(http.StatusOK, getProfileResponse{
		IsConfirmed:    infoCard.IsConfirmed,
		CreatedAt:      infoCard.CreatedDate.String(),
		DocumentType:   document.Type.String(),
		SerialNumber:   document.SerialNumber,
		DocumentFields: model.ModelToKeyValue(documentFields),
	})
}

// GetEmployeePhoto godoc
//
//	@Summary		Получение фотографии своего профиля
//	@Description	Метод для получения фотографии своего профиля
//	@Tags			employee
//	@Produce		jpeg
//	@Produce		json
//	@Success		200	{string}	string							"Фотография успешно получена"
//	@Failure		400	{object}	http.StatusBadRequest			"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusUnauthorized			"Авторизация неуспешна"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка получения фотографии"
//	@Security		BearerAuth
//	@Router			/employee-photo [get]
func (p *ProfileController) GetEmployeePhoto(c *gin.Context) {
	payload, err := httputils.VerifyAccessToken(c, p.l, p.authService)
	if err != nil {
		return
	}
	infoCardID, err := payload.GetInfoCardID()
	if err != nil {
		p.l.Errorf("failed to parse infoCard id from payload: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	document, err := p.documentService.GetDocumentByInfoCard(c.Request.Context(), &dto.GetDocumentByInfoCardIDRequest{
		InfoCardID: infoCardID,
	})
	if err != nil {
		p.l.Errorf("failed to get document by infoCard ID: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get info card document"})
		return
	}

	photoData, err := p.photoService.GetPhoto(c.Request.Context(), &dto.GetPhotoRequest{
		DocumentID: document.ID.Int(),
	})
	if err != nil {
		p.l.Errorf("failed to get employee photo: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get employee photo"})
		return
	}

	c.Data(http.StatusOK, "image/jpeg", photoData.Data)
	c.JSON(http.StatusOK, "OK")
}
