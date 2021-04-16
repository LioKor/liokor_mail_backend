package usecase

import (
	"golang.org/x/crypto/bcrypt"
	"liokor_mail/internal/pkg/common"
	"liokor_mail/internal/pkg/user"
	"liokor_mail/internal/pkg/user/validators"
	"log"
	"os"
	"strings"
	"time"
)

type UserUseCase struct {
	Repository user.UserRepository
	Config     common.Config
}

func (uc *UserUseCase) Login(credentials user.Credentials) error {
	loginUser, err := uc.Repository.GetUserByUsername(credentials.Username)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(loginUser.HashPassword), []byte(credentials.Password))

	if err != nil {
		return user.InvalidUserError{"Invalid credentials"}
	}

	return nil
}

func (uc *UserUseCase) Logout(sessionToken string) error {
	return uc.Repository.RemoveSession(sessionToken)
}

func (uc *UserUseCase) CreateSession(username string) (user.SessionToken, error) {
	sessionToken := user.SessionToken{
		common.GenerateRandomString(),
		time.Now().Add(10 * 24 * time.Hour),
	}

	err := uc.Repository.CreateSession(user.Session{
		username,
		sessionToken.Value,
		sessionToken.Expiration,
	})

	if err != nil {
		return user.SessionToken{}, err
	}

	return sessionToken, nil
}

func (uc *UserUseCase) GetUserBySessionToken(sessionToken string) (user.User, error) {
	session, err := uc.Repository.GetSessionBySessionToken(sessionToken)
	if err != nil {
		return user.User{}, err
	}

	if session.Expiration.Before(time.Now()) {
		return user.User{}, user.InvalidSessionError{"session token expired"}
	}

	sessionUser, err := uc.Repository.GetUserByUsername(session.Username)
	if err != nil {
		return user.User{}, err
	}

	return sessionUser, nil
}

func (uc *UserUseCase) SignUp(newUser user.UserSignUp) error {
	if !validators.ValidateUsername(newUser.Username) {
		return user.InvalidUsernameError{"invalid username"}
	}
	if !validators.ValidatePassword(newUser.Password) {
		return user.WeakPasswordError{"password is too weak"}
	}

	hashPSWD, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = uc.Repository.CreateUser(user.User{
		0,
		newUser.Username,
		string(hashPSWD),
		newUser.AvatarURL,
		newUser.FullName,
		newUser.ReserveEmail,
		time.Now().String(),
		false,
	})
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCase) UpdateUser(username string, newData user.User) (user.User, error) {
	sessionUser, err := uc.Repository.GetUserByUsername(username)
	if err != nil {
		return user.User{}, err
	}
	if newData.AvatarURL != sessionUser.AvatarURL {
		if strings.HasPrefix(newData.AvatarURL, "data:") {
			avatarFileName := common.GenerateRandomString()
			pathToAvatar, err := common.DataURLToFile(uc.Config.AvatarStoragePath+avatarFileName, newData.AvatarURL, 500)
			if err != nil {
				log.Println(err.Error())
				return sessionUser, user.InvalidImageError{"invalid image"}
			}
			if len(sessionUser.AvatarURL) > 0 {
				_ = os.Remove(sessionUser.AvatarURL)
			}
			sessionUser.AvatarURL = pathToAvatar
		} else {
			return sessionUser, user.InvalidImageError{"invalid image"}
		}
	}
	if newData.FullName != sessionUser.FullName {
		sessionUser.FullName = newData.FullName
	}
	if newData.ReserveEmail != sessionUser.ReserveEmail {
		sessionUser.ReserveEmail = newData.ReserveEmail
	}

	sessionUser, err = uc.Repository.UpdateUser(username, sessionUser)
	if err != nil {
		return user.User{}, err
	}
	return sessionUser, nil
}

func (uc *UserUseCase) GetUserByUsername(username string) (user.User, error) {
	requestedUser, err := uc.Repository.GetUserByUsername(username)
	if err != nil {
		return user.User{}, err
	}

	return requestedUser, nil
}

func (uc *UserUseCase) ChangePassword(sessionUser user.User, changePSWD user.ChangePassword) error {
	if !validators.ValidatePassword(changePSWD.NewPassword) {
		return user.InvalidUserError{"invalid password"}
	}

	err := bcrypt.CompareHashAndPassword([]byte(sessionUser.HashPassword), []byte(changePSWD.OldPassword))

	if err != nil {
		return user.InvalidUserError{"Invalid password"}
	}

	hashPSWD, err := bcrypt.GenerateFromPassword([]byte(changePSWD.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return uc.Repository.ChangePassword(sessionUser.Username, string(hashPSWD))

}
