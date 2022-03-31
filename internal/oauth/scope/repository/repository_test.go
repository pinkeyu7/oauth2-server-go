package repository

import (
	"oauth2-server-go/config"
	"oauth2-server-go/driver"
	"oauth2-server-go/dto/model"
	"oauth2-server-go/pkg/valider"
	"os"
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

func TestRepository_Find(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	re, _ := driver.NewRedis()
	osc := NewRedis(re)
	osr := NewRepository(orm, osc)

	// Act
	scopes, err := osr.Find()

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, scopes)
	assert.Len(t, scopes, 4)
}

func TestRepository_FindOne(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	re, _ := driver.NewRedis()
	osc := NewRedis(re)
	osr := NewRepository(orm, osc)

	// Act
	scope, err := osr.FindOne(&model.OauthScope{
		Path:   "/v1/users",
		Method: "GET",
	})

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, scope)
}

func TestRepository_FindScope(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	re, _ := driver.NewRedis()
	osc := NewRedis(re)
	osr := NewRepository(orm, osc)

	// Act
	scopes, err := osr.FindScope()

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, scopes)
	assert.Len(t, scopes, 4)
}
