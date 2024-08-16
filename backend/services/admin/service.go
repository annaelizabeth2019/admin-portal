package admin

import (
	"admin-portal/model"
	"admin-portal/repository"
	"crypto/rsa"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Service is the service interface exported by this package.
type Service interface {
	Get(userID int64) (model.User, error)
	Authenticate(email, password string) (model.User, error)
	LogIn(c *gin.Context, u model.User) error
	LogOut(c *gin.Context, userID int64) error
	UpdatePassword(password string, userID int64) error
}

type service struct {
	jwtPublicKey   *rsa.PublicKey
	jwtPrivateKey  *rsa.PrivateKey
	userRepository repository.User
	authRepository repository.Auth
}

func NewService(db *gorm.DB, publicKey *rsa.PublicKey, privateKey *rsa.PrivateKey) Service {
	return &service{
		jwtPublicKey:   publicKey,
		jwtPrivateKey:  privateKey,
		userRepository: repository.NewUserRepo(db),
		authRepository: repository.NewAuthRepo(db),
	}
}

func (s *service) Get(userID int64) (model.User, error) {
	user, err := s.userRepository.SelectOne(userID)
	if err != nil {
		return user, err
	}
	roles, err := s.userRepository.GetRolesForUser(userID)
	if err != nil {
		return user, err
	}
	user.Roles = make(map[model.AdminRole]model.Role, 0)
	for _, role := range roles {
		user.Roles[role.Title] = role
	}
	return user, err
}

func (s *service) Authenticate(email, password string) (u model.User, err error) {
	userAuth, err := s.authRepository.SelectByEmail(email)
	if err != nil {
		return u, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userAuth.Password), []byte(password))
	if err != nil {
		return u, err
	}

	return s.userRepository.SelectOne(userAuth.ID)
}

func (s *service) LogIn(c *gin.Context, u model.User) error {
	return nil
}

func (s *service) LogOut(c *gin.Context, userID int64) error {
	return nil
}

func (s *service) UpdatePassword(password string, userID int64) error {
	h, err := hash(password)
	if err != nil {
		return err
	}
	return s.authRepository.SetPassword(h, userID)
}
