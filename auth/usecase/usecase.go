package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/DarkSoul94/task_tracker_server/pkg/logger"
	"github.com/DarkSoul94/task_tracker_server/user_manager"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/spf13/viper"
)

// Usecase ...
type Usecase struct {
	userManager    user_manager.UserManagerUC
	secret         string
	signingKey     []byte
	expireDuration time.Duration
}

type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

// NewUsecase ...
func NewUsecase(
	userManager user_manager.UserManagerUC,
	secret string,
	signingKey []byte,
	tokenTTLSeconds time.Duration) *Usecase {
	return &Usecase{
		userManager:    userManager,
		secret:         secret,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTLSeconds,
	}
}

func (u *Usecase) LDAPSignIn(email, password string) (models.User, string, error) {
	var (
		user  models.User
		token string
	)
	lUser, ok := u.ldapAuthenticate(email, password)
	if !ok {
		return models.User{}, "", ErrLoginFailed
	}

	user, err := u.userManager.GetUserByEmail(email)
	if err != nil {

		user = models.User{
			Email:      email,
			Name:       lUser.Name,
			Department: lUser.Department,
		}
		u.userManager.CreateUser(&user)
		user, err = u.userManager.GetUserByEmail(email)
		if err != nil {
			return models.User{}, "", err
		}
	} else {
		if user.Name != lUser.Name || user.Department != lUser.Department {
			user.Name = lUser.Name
			user.Department = lUser.Department
			u.userManager.CreateUser(&user)
		}
	}
	token, err = u.GenerateToken(&user)
	if err != nil {
		return models.User{}, "", err
	}
	return user, token, nil
}

func (u *Usecase) GenerateToken(user *models.User) (string, error) {
	var (
		token    *jwt.Token
		strToken string
		err      error
	)

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(u.expireDuration)),
		},
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err = token.SignedString(u.signingKey)
	if err != nil {
		logger.LogError(ErrCreateToken.Error(), "auth/usecase", user.Name, err)
		return "", ErrCreateToken
	}

	return strToken, nil
}

func (u *Usecase) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return u.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, nil
}

func (u *Usecase) ldapAuthenticate(email, password string) (ldapUser, bool) {
	ldap := NewLdapAuthenticator(
		viper.GetString("app.auth.ldap.server"),
		viper.GetString("app.auth.ldap.baseDN"),
		viper.GetString("app.auth.ldap.filterDN"),
	)
	user, err := ldap.Auth(email, password)
	if err != nil {
		return user, false
	}
	return user, true
}
