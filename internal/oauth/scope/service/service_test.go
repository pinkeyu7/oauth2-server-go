package service

import (
	"net/http"
	"oauth2-server-go/config"
	"oauth2-server-go/driver"
	clientRepo "oauth2-server-go/internal/oauth/client/repository"
	scopeRepo "oauth2-server-go/internal/oauth/scope/repository"
	"oauth2-server-go/pkg/er"
	"oauth2-server-go/pkg/valider"
	"os"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	config.InitEnv()
	valider.Init()
}

func TestService_GetScope(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	re, _ := driver.NewRedis()

	ocr := clientRepo.NewRepository(orm)
	osc := scopeRepo.NewRedis(re)
	osr := scopeRepo.NewRepository(orm, osc)
	oss := NewService(osr, osc, ocr)

	// Act
	_, err := oss.GetScope("/v1/users", "DELETE")
	assert.NotNil(t, err)
	notFoundErr := err.(*er.AppError)
	assert.Equal(t, http.StatusBadRequest, notFoundErr.StatusCode)
	assert.Equal(t, strconv.Itoa(er.ResourceNotFoundError), notFoundErr.Code)

	// Act
	scope, err := oss.GetScope("/v1/users", "GET")
	assert.Nil(t, err)
	assert.Equal(t, "user.profile_get", scope)
}

func TestService_VerifyScope(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	re, _ := driver.NewRedis()

	ocr := clientRepo.NewRepository(orm)
	osc := scopeRepo.NewRedis(re)
	osr := scopeRepo.NewRepository(orm, osc)
	oss := NewService(osr, osc, ocr)

	clientId := "address-book-go"
	scope := "user.profile_get"

	isPass, err := oss.VerifyScope(clientId, scope)
	assert.Nil(t, err)
	assert.Equal(t, true, isPass)
}
