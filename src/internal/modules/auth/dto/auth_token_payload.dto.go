package dto

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthTokenPayloadDto struct {
	UserID     uuid.UUID `validate:"required,uuid" json:"userId"`
	Subscriber uuid.UUID `validate:"required,uuid" json:"subscriber"`
	jwt.RegisteredClaims
}
