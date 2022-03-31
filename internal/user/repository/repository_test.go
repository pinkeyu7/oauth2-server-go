package repository

import (
	"fmt"
	"log"
	"oauth2-server-go/config"
	"oauth2-server-go/driver"
	"oauth2-server-go/dto/model"
	"oauth2-server-go/pkg/valider"
	"os"
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

func TestRepository_Insert(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	ur := NewRepository(orm)

	m := model.User{
		Account:  "test_account",
		Phone:    "test_phone",
		Email:    "test_email",
		Password: "test_password",
		Name:     "test_name",
	}

	// Act
	err := ur.Insert(&m)

	// Assert
	assert.Nil(t, err)

	// Teardown
	_, _ = orm.ID(m.Id).Delete(&model.User{})
}

func TestRepository_Find(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	ur := NewRepository(orm)

	// Act
	testCases := []struct {
		Limit     int
		Offset    int
		WantCount int
	}{
		{
			2,
			0,
			1,
		},
		{
			10,
			10,
			0,
		},
	}
	// Act
	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Find User,Offset:%d,Limit:%d", tc.Offset, tc.Limit), func(t *testing.T) {
			data, err := ur.Find(tc.Offset, tc.Limit)
			assert.Nil(t, err)
			assert.Len(t, data, tc.WantCount)
		})
	}
}

func TestRepository_FindOne(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	ur := NewRepository(orm)

	// No data
	// Act
	res, err := ur.FindOne(&model.User{Id: 100})

	// Assert
	assert.Nil(t, err)
	assert.Nil(t, res)

	// Has data
	// Act
	res, err = ur.FindOne(&model.User{Id: 1})

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 1, res.Id)
}

func TestRepository_Count(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	ur := NewRepository(orm)

	// Act
	count, err := ur.Count()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 1, count)
}

func TestRepository_Update(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	ur := NewRepository(orm)

	user := model.User{Id: 1}
	_, _ = orm.Get(&user)

	m := model.User{
		Account:  "test_account",
		Phone:    "test_phone",
		Email:    "test_email",
		Password: "test_password",
		Name:     "test_name",
	}

	// Act
	err := ur.Update(&m)

	// Assert
	assert.Nil(t, err)

	// Teardown
	_, _ = orm.ID(user.Id).Update(&user)
}

func TestRepository_Delete(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	ur := NewRepository(orm)

	user := model.User{
		Account:  "test_account",
		Phone:    "test_phone",
		Email:    "test_email",
		Password: "test_password",
		Name:     "test_name",
	}
	_, _ = orm.Insert(&user)

	// Act
	err := ur.Delete(user.Id)

	// Assert
	assert.Nil(t, err)
}
