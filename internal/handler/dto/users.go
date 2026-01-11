package dto

import (
	"github.com/Nios-V/Go-ecommerce-API/internal/models"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	Email     string    `json:"email"`
	// Roles     []RoleResponse    `json:"roles"`
	// Addresses []AddressResponse `json:"addresses"`
}

func ToUserResponse(u *models.User) *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Email:     u.Email,
	}
}

func ToUserListResponse(users []models.User) []*UserResponse {
	res := make([]*UserResponse, len(users))
	for i, u := range users {
		res[i] = ToUserResponse(&u)
	}
	return res
}

type CreateUserRequest struct {
	Firstname string `json:"first_name" validate:"required,min=2,max=30"`
	Lastname  string `json:"last_name" validate:"required,min=2,max=30"`
	Email     string `json:"email" validate:"required,email,max=50"`
	Password  string `json:"password" validate:"required,min=6,max=100"`
}

func (r *CreateUserRequest) CreateUserToModel() *models.User {
	return &models.User{
		Firstname: r.Firstname,
		Lastname:  r.Lastname,
		Email:     r.Email,
		Password:  r.Password,
	}
}

type UpdateUserRequest struct {
	Firstname *string `json:"first_name" validate:"omitempty,min=2,max=30"`
	Lastname  *string `json:"last_name" validate:"omitempty,min=2,max=30"`
	Email     *string `json:"email" validate:"omitempty,email,max=50"`
}

func (r *UpdateUserRequest) UpdateUserToModel(user *models.User) {
	if r.Firstname != nil {
		user.Firstname = *r.Firstname
	}
	if r.Lastname != nil {
		user.Lastname = *r.Lastname
	}
	if r.Email != nil {
		user.Email = *r.Email
	}
}

type UpdatePasswordRequest struct {
	Password string `json:"password" validate:"required,min=6,max=100"`
}

func (r *UpdatePasswordRequest) UpdatePasswordToModel(user *models.User) {
	user.Password = r.Password
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}
