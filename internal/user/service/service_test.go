package service

import (
	"log"
	"net/http"
	"oauth2-server-go/config"
	"oauth2-server-go/driver"
	userRepo "oauth2-server-go/internal/user/repository"
	"oauth2-server-go/pkg/er"
	"oauth2-server-go/pkg/valider"
	"os"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	remoteBranch := os.Getenv("REMOTE_BRANCH")
	if remoteBranch == "" {
		// load env
		err := godotenv.Load(config.GetBasePath() + "/.env")
		if err != nil {
			log.Panicln(err)
		}
	}

	valider.Init()
}

func TestService_Get(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	ur := userRepo.NewRepository(orm)
	us := NewService(ur)

	// No data
	userId := 10

	// Act
	usr, err := us.Get(userId)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, usr)
	notFoundErr := err.(*er.AppError)
	assert.Equal(t, http.StatusBadRequest, notFoundErr.StatusCode)
	assert.Equal(t, strconv.Itoa(er.ResourceNotFoundError), notFoundErr.Code)

	// Has data
	userId = 1

	// Act
	usr, err = us.Get(userId)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, userId, usr.Id)
}

func TestService_Verify(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	ur := userRepo.NewRepository(orm)
	us := NewService(ur)

	// No data
	account := "account"
	password := "password"

	// Act
	err := us.Verify(account, password)

	// Assert
	assert.NotNil(t, err)
	notFoundErr := err.(*er.AppError)
	assert.Equal(t, http.StatusBadRequest, notFoundErr.StatusCode)
	assert.Equal(t, strconv.Itoa(er.ResourceNotFoundError), notFoundErr.Code)

	// Has data
	account = "user_1"
	password = "A12345678"

	// Act
	err = us.Verify(account, password)

	// Assert
	assert.Nil(t, err)
}
