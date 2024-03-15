package service

import (
	"crypto/sha1"
	"fmt"
	GO_Project "github.com/zhetibayev0410/Go-1-project"
	"github.com/zhetibayev0410/Go-1-project/pkg/repository"
	"time"
)

const (
	salt       = "jb1k2jb14b3j"
	signingKey = "grggqg@ene1"
	token      = 12 * time.Hour
)

type tokenClaims struct {
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user GO_Project.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

//func (s *AuthService) GenerateToken(username, password string) (string, error) {
//	//get user from db
//}

func generatePasswordHash(pasword string) string {
	hash := sha1.New()
	hash.Write([]byte(pasword))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
