package http

import (
	"net/http"
	"path"
	"smplrstapp/internal/dto"
	"smplrstapp/internal/entity"
	"smplrstapp/internal/service"
	"smplrstapp/internal/utils/req"
	"smplrstapp/internal/utils/resp"
	"smplrstapp/internal/utils/time"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

type activityHandler struct {
	service service.ActivityService
}

func NewActivityHandler(service *service.ActivityService) Handler {
	return &activityHandler{
		service: *service,
	}
}

func (h *activityHandler) Register(router *mux.Router) {
	router.HandleFunc(path.Join(apiVersion, addActivityUrl), h.AddActivity).Methods(methodPost)
}

// @Summary      Create New Activity At User
// @Description  create activity for user and return result
// @Tags         activities
// @Accept       json
// @Produce      json
// @Param        data body dto.ActivityCreateDto true "The input activity dto"
// @Success      201  {object}  entity.Activity
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /activities [post]
func (h *activityHandler) AddActivity(writer http.ResponseWriter, request *http.Request) {
	jsonDto, err := req.GetModelFromBodyRequest(request)
	if check(err) {
		resp.RespondWithError(&writer, http.StatusBadRequest, err.Error())
		return
	}

	var valueDto dto.ActivityCreateDto
	err = mapstructure.Decode(jsonDto, &valueDto)
	if check(err) {
		resp.RespondWithError(&writer, http.StatusBadRequest, errInvalidBody)
		return
	}

	trainingDate, err := time.GetTimeFromString(valueDto.TrainingDate)
	if check(err) {
		resp.RespondWithError(&writer, http.StatusBadRequest, err.Error())
		return
	}

	newActivity := entity.CreateActivity(valueDto.UserId, valueDto.Sport, valueDto.Distance, trainingDate, valueDto.Duration)
	result, err := h.service.Create(request.Context(), newActivity)
	if check(err) {
		resp.RespondWithError(&writer, http.StatusBadRequest, err.Error())
		return
	}

	resp.RespondWithJSON(&writer, http.StatusCreated, result)
}
