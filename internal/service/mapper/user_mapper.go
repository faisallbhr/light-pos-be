package mapper

import (
	"github.com/faisallbhr/light-pos-be/internal/dto"
	"github.com/faisallbhr/light-pos-be/internal/entities"
)

func ToUserResponse(user *entities.User) *dto.UserResponse {
	var roles []dto.RoleResponse
	for _, r := range user.Roles {
		roles = append(roles, dto.RoleResponse{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Roles: roles,
	}
}

func ToUserResponses(users []*entities.User) []*dto.UserResponse {
	res := make([]*dto.UserResponse, 0, len(users))
	for _, user := range users {
		res = append(res, ToUserResponse(user))
	}
	return res
}
