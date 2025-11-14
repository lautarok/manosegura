package service

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/domain"
	"github.com/lautarok/manosegura/src/repository"
)

type CredentialsService struct {
	credentialsRepository repository.CredentialsRepository
}

func NewCredentialsService(credentialsRepository repository.CredentialsRepository) *CredentialsService {
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
