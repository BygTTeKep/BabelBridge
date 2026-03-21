package dtos

import "github.com/google/uuid"

type CreateCompanyResponseDto struct {
	ID    int
	Name  string
	Token uuid.UUID
}
