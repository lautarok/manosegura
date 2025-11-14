package services

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/internal/modules/credentials/domain"
	"github.com/lautarok/manosegura/src/internal/modules/credentials/interfaces"
)

type CredentialsService struct {
	credentialsRepository interfaces.CredentialsRepository
}

func NewCredentialsService(credentialsRepository interfaces.CredentialsRepository) *CredentialsService {
	return &CredentialsService{
		credentialsRepository: credentialsRepository,
	}
}

func (credentialsService *CredentialsService) FindOneByUserId(userId uuid.UUID) (*domain.Credential, error) {
	return credentialsService.credentialsRepository.FindOneByUserId(userId)
}

func (credentialsService *CredentialsService) FindOneByUsername(username string) (*domain.Credential, error) {
	return credentialsService.credentialsRepository.FindOneByUsername(username)
}

func (credentialsService *CredentialsService) EmailExists(email string) (bool, error) {
	return credentialsService.credentialsRepository.EmailExists(email)
}

func (credentialsService *CredentialsService) UsernameExists(username string) (bool, error) {
	return credentialsService.credentialsRepository.UsernameExists(username)
}
