package users_transport_http

import (
	"net/http"

	"github.com/pkpal-uhobp/todo-app/internal/core/domain"
	core_logger "github.com/pkpal-uhobp/todo-app/internal/core/logger"
	core_http_request "github.com/pkpal-uhobp/todo-app/internal/core/transport/http/request"
	core_http_response "github.com/pkpal-uhobp/todo-app/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	Fullname    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

type CreateUserResponse struct {
	Id          int     `json:"id"`
	Version     int     `json:"version"`
	Fullname    string  `json:"full_name"`
	Phonenumber *string `json:"phone_number"`
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	log.Debug("invoice CreateUser handler")
	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "Failed to decode and validate HTTP request")
		return
	}
	userDomain := domainFromDTO(request)
	userDomain, err := h.usersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to create user")
		return
	}
	response := dtoFromDomain(userDomain)
	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.Fullname, dto.PhoneNumber)
}

func dtoFromDomain(user domain.User) CreateUserResponse {
	return CreateUserResponse{
		Id:          user.ID,
		Version:     user.Version,
		Fullname:    user.FullName,
		Phonenumber: user.PhoneNumber,
	}
}
