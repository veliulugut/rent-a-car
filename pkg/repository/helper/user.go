package helper

import (
	"rent-a-car/ent"
	"rent-a-car/pkg/repository/dto"
)

func DbUserToDTO(d *ent.User) *dto.User {
	return &dto.User{
		FirstName:   d.FirstName,
		LastName:    d.LastName,
		UserName:    d.Username,
		Email:       d.Email,
		PhoneNumber: d.PhoneNumber,
		Password:    d.Password,
		CreateAt:    d.CreatedAt,
		UpdateAt:    d.UpdatedAt,
		ID:          d.ID,
	}
}
