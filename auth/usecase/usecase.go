package usecase

import (
	"time"

	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/DarkSoul94/task_tracker_server/pkg/logger"
	"github.com/DarkSoul94/task_tracker_server/user_manager"
	"github.com/dgrijalva/jwt-go/v4"
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

func (u *Usecase) SignUp(user *models.LoginUser) (models.User, error) {
	var (
		hash string
		err  error
	)
	hash, err = user.GetPassHash()
	if err != nil {
		return models.User{}, err
	}

	return u.userManager.CreateUser(user.Name, hash)
}

func (u *Usecase) SignIn(inpUser *models.LoginUser) (string, models.User, error) {
	var (
		user     models.User
		token    *jwt.Token
		strToken string
		err      error
	)

	user, err = u.userManager.GetUserByName(inpUser.Name)
	if err != nil {
		return strToken, models.User{}, err
	}

	if inpUser.VerifyPass(user.PassHash) {
		claims := AuthClaims{
			User: &user,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: jwt.At(time.Now().Add(u.expireDuration)),
			},
		}
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		strToken, err = token.SignedString(u.signingKey)
		if err != nil {
			logger.LogError(ErrCreateToken.Error(), "auth/usecase", user.Name, err)
			return "", models.User{}, ErrCreateToken
		}
		return strToken, user, nil
	} else {
		return strToken, models.User{}, ErrLoginFailed
	}
}
