package dto

import "github.com/google/uuid"

type RequestMailDTO struct {
	Destination string `json:"destination" binding:"required,email"`
	Subject     string `json:"subject" binding:"required"`
	Body        string `json:"body" binding:"required"`
}

type RequestMailResponseDTO struct {
	ID          uuid.UUID `json:"id"`
	Destination string    `json:"destination"`
	Subject     string    `json:"subject"`
}
