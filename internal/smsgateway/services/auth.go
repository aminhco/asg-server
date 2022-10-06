package services

import (
	"fmt"

	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/models"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/repositories"
	"bitbucket.org/soft-c/gohelpers/pkg/crypto"
	"github.com/jaevor/go-nanoid"
)

type AuthService struct {
	users   *repositories.UsersRepository
	devices *repositories.DevicesRepository

	idgen func() string
}

func (s *AuthService) RegisterUser(login, password string) (models.User, error) {
	user := models.User{
		ID: login,
	}

	var err error
	if user.PasswordHash, err = crypto.MakeBCryptHash(password); err != nil {
		return user, err
	}

	if err = s.users.Insert(&user); err != nil {
		return user, fmt.Errorf("can't create user")
	}

	return user, nil
}

func (s *AuthService) RegisterDevice(userID, name, pushToken string) (models.Device, error) {
	device := models.Device{
		ID:        s.idgen(),
		Name:      name,
		AuthToken: s.idgen(),
		PushToken: pushToken,
		UserID:    userID,
	}

	return device, s.devices.Insert(&device)
}

func (s *AuthService) AuthorizeDevice(token string) (models.Device, error) {
	return s.devices.GetByToken(token)
}

func NewAuthService(users *repositories.UsersRepository, devices *repositories.DevicesRepository) *AuthService {
	idgen, _ := nanoid.Standard(21)

	return &AuthService{
		users:   users,
		devices: devices,
		idgen:   idgen,
	}
}
