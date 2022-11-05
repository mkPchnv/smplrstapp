package http

import (
	"net/http"
	"path"
	"strconv"

	"smplrstapp/internal/dto"
	"smplrstapp/internal/entity"
	"smplrstapp/internal/service"
	"smplrstapp/internal/utils/req"
	"smplrstapp/internal/utils/resp"

	"github.com/gorilla/mux"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service *service.UserService) Handler {
	return &userHandler{
		service: *service,
	}
}

func (h *userHandler) Register(router *mux.Router) {
	router.HandleFunc(path.Join(apiVersion, getUserUrl), h.GetUserById).Methods(methodGet)
	router.HandleFunc(path.Join(apiVersion, addUserUrl), h.AddUser).Methods(methodPost)
	router.HandleFunc(path.Join(apiVersion, getUsersUrl), h.GetAllUsers).Methods(methodGet)
}

// @Summary      Get User by id
// @Description  get user by identificator
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  entity.User
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /users/{id} [get]
func (h *userHandler) GetUserById(writer http.ResponseWriter, request *http.Request) {
	userId, err := req.GetParamFromRequest(request, idParam, requestWithoutParams)
	if isError := check(err); isError {
		resp.RespondWithError(&writer, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.GetById(request.Context(), userId)
	if isError := check(err); isError {
		resp.RespondWithError(&writer, http.StatusNotFound, notFoundUser)
		return
	}

	resp.RespondWithJSON(&writer, http.StatusOK, user)
}

// @Summary      Create New User
// @Description  create user and return result
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        data body dto.UserCreateDto true "The input user dto"
// @Success      201  {object}  entity.User
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /users [post]
func (h *userHandler) AddUser(writer http.ResponseWriter, request *http.Request) {
	valueDto := &dto.UserCreateDto{}
	err := req.GetModelFromBodyRequest(request, valueDto)
	if check(err) {
		resp.RespondWithError(&writer, http.StatusBadRequest, err.Error())
		return
	}

	newUser := entity.CreateUser(valueDto.FirstName, valueDto.LastName, valueDto.Age, valueDto.Email)
	result, err := h.service.Create(request.Context(), newUser)
	if check(err) {
		resp.RespondWithError(&writer, http.StatusBadRequest, err.Error())
		return
	}

	resp.RespondWithJSON(&writer, http.StatusCreated, result)
}

// @Summary      Get All Users
// @Description  get all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        limit   query   int   false   "Limit"    10
// @Param        offset  query   int   false   "Offset"   0
// @Success      200  {array}  entity.User
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /users [get]
func (h *userHandler) GetAllUsers(writer http.ResponseWriter, request *http.Request) {
	var limit, offset = 10, 0
	limits, _ := req.GetQueryFromRequest(request, limitQueryParam, emptyMessage)
	if len(limits) == 1 {
		limit, _ = strconv.Atoi(limits[0])
	}

	offsets, _ := req.GetQueryFromRequest(request, offsetQueryParam, emptyMessage)
	if len(offsets) == 1 {
		offset, _ = strconv.Atoi(offsets[0])
	}

	result, err := h.service.GetAll(request.Context(), limit, offset)
	if check(err) {
		resp.RespondWithError(&writer, http.StatusBadRequest, err.Error())
		return
	}

	resp.RespondWithJSON(&writer, http.StatusOK, result)
}
