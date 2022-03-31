package service

import (
	"errors"
	"log"
	"strings"
	"sync"

	"github.com/MakMoinee/appInviteService/internal/appInviteService/models"
	"github.com/MakMoinee/appInviteService/internal/appInviteService/repository/mysqllocal"
	"github.com/MakMoinee/go-mith/pkg/encrypt"
)

// SendGenerationOfToken - sends the generation of token based on the user credential
func SendGenerationOfToken(user models.User, localsql mysqllocal.TokenIntf) (models.Token, error) {
	log.Println("Inside SendGenerationOfToken()")
	token := models.Token{}

	if isValid, errs := validateUserRequest(user); !isValid && errs != nil {
		return token, errs
	}

	userFromDb, err := localsql.GetUser(user)
	if err != nil {
		return token, err
	}

	if len(userFromDb.Username) > 0 && len(userFromDb.Password) > 0 {
		passwordMatched := encrypt.CheckPasswordHash(user.Password, userFromDb.Password)
		if !passwordMatched {
			err = errors.New("password doesn't match")
		}
		return token, err
	}

	// do a concurrent call
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		token, err = localsql.GenerateToken(user)
	}()
	wg.Wait()

	return token, err
}

// validateUserRequest - validates basic user info
func validateUserRequest(user models.User) (bool, error) {
	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)

	if len(user.Username) == 0 || len(user.Password) == 0 {
		return false, errors.New("missing required parameters")
	}
	return true, nil
}
