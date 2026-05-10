package users_transport_http

import "github.com/pkpal-uhobp/todo-app/internal/core/domain"

type UserDTOResponse struct {
	Id          int     `json:"id"`
	Version     int     `json:"version"`
	Fullname    string  `json:"full_name"`
	Phonenumber *string `json:"phone_number"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		Id:          user.ID,
		Version:     user.Version,
		Fullname:    user.FullName,
		Phonenumber: user.PhoneNumber,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))
	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}
	return usersDTO
}
