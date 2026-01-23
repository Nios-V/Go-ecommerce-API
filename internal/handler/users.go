package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Nios-V/Go-ecommerce-API/internal/handler/dto"
	"github.com/Nios-V/Go-ecommerce-API/internal/response"
	"github.com/Nios-V/Go-ecommerce-API/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService *service.UserService
	validate    *validator.Validate
}

func NewUserHandler(s *service.UserService, v *validator.Validate) *UserHandler {
	return &UserHandler{
		userService: s,
		validate:    v,
	}
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get user details by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID" format(uuid)
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	user, err := h.userService.GetByID(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to retrieve user")
		return
	}

	res := dto.ToUserResponse(user)
	response.JSON(w, http.StatusOK, res)
}

// Login godoc
// @Summary User login
// @Description Authenticate user and return a token
// @Tags Users
// @Accept json
// @Produce json
// @Param credentials body dto.LoginRequest true "User credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} response.StandardResponse
// @Failure 401 {object} response.StandardResponse
// @Router /users/login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		if errors.Is(err, io.EOF) {
			response.Error(w, http.StatusBadRequest, "Empty body")
			return
		}
		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	defer r.Body.Close()

	if err := h.validate.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, "Validation Error: "+err.Error())
		return
	}

	user, err := h.userService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"token": user})
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "User data"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		if errors.Is(err, io.EOF) {
			response.Error(w, http.StatusBadRequest, "Empty body")
			return
		}
		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	defer r.Body.Close()
	if err := h.validate.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, "Validation Error: "+err.Error())
		return
	}

	user := *req.CreateUserToModel()

	err = h.userService.Create(r.Context(), &user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update an existing user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID" format(uuid)
// @Param user body dto.UpdateUserRequest true "User data"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} response.StandardResponse
// @Failure 500 {object} response.StandardResponse
// @Router /users/{id} [put]
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	user, err := h.userService.GetByID(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "User not found")
		return
	}

	var req dto.UpdateUserRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		if errors.Is(err, io.EOF) {
			response.Error(w, http.StatusBadRequest, "Empty body")
			return
		}
		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	defer r.Body.Close()
	if err := h.validate.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, "Validation Error: "+err.Error())
		return
	}

	req.UpdateUserToModel(user)

	err = h.userService.Update(r.Context(), user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	response.JSON(w, http.StatusOK, user)
}
