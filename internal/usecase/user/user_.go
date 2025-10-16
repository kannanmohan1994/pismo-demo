package user

import (
	"time"
	"xm/config"
	"xm/internal/entity/models"
	"xm/internal/entity/response"
	"xm/internal/middleware"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (us *usecase) CreateUser(name, password string) (resp *response.UserResponse, err error) {
	us.logger.Infof("Begin Usecase - CreateUser")

	user := &models.User{
		ID:       uuid.New().String(),
		Name:     name,
		Password: password,
	}

	accessToken, err := us.CreateUserAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	_, err = us.user.CreateUser(user)
	if err != nil {
		return nil, err
	}

	us.logger.Infof("End Usecase - CreateUser")
	return &response.UserResponse{
		AccessToken: accessToken,
	}, nil
}
func (us *usecase) LoginUser(name, password string) (resp *response.UserResponse, err error) {
	us.logger.Infof("Begin Usecase - LoginUser")

	user, err := us.user.GetUser(name, password)
	if err != nil {
		return nil, err
	}

	accessToken, err := us.CreateUserAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	us.logger.Infof("End Usecase - LoginUser")
	return &response.UserResponse{
		AccessToken: accessToken,
	}, nil
}

func (us *usecase) CreateUserAccessToken(userId string) (string, error) {
	accessTokenExpiry := config.GetConfig().AccessTokenExpiryDurationSeconds
	expiryDuration := time.Duration(accessTokenExpiry) * time.Second

	return us.createCustomJWTToken(userId, expiryDuration)
}

func (us *usecase) createCustomJWTToken(userId string, expiryDuration time.Duration) (string, error) {
	kid := config.GetConfig().JWKKid

	key, _ := us.middleware.GetKeyByKID(kid)
	currTime := time.Now().UTC()

	claims := middleware.JWTClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    middleware.TokenIssuer,
			Subject:   middleware.TokenSubject,
			Audience:  []string{},
			ExpiresAt: jwt.NewNumericDate(currTime.Add(expiryDuration)),
			NotBefore: jwt.NewNumericDate(currTime),
			IssuedAt:  jwt.NewNumericDate(currTime),
			ID:        middleware.TokenID,
		},
	}

	return us.middleware.CreateJWTTokenWithClaims(claims, key, kid)
}
