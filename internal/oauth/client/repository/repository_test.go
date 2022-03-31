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
	ocr := NewRepository(orm)

	oauthClientId := "test_id"
	m := model.OauthClient{
		Id:           oauthClientId,
		SysAccountId: 1,
		Name:         "test_name",
		Secret:       "test_secret",
		Domain:       "test_domain",
		Scope:        "test_scope",
		IconPath:     "test_icon_path",
		Data:         "",
	}

	// Act
	err := ocr.Insert(&m)

	// Assert
	assert.Nil(t, err)

	// Teardown
	_, _ = orm.Delete(&model.OauthClient{Id: oauthClientId})
}

func TestRepository_Find(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	ocr := NewRepository(orm)

	// Act
	testCases := []struct {
		Limit     int
		Offset    int
		WantCount int
	}{
		{
			1,
			0,
			1,
		},
		{
			10,
			0,
			2,
		},
	}
	// Act
	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Find Oauth Client,Offset:%d,Limit:%d", tc.Offset, tc.Limit), func(t *testing.T) {
			data, err := ocr.Find(tc.Offset, tc.Limit)
			assert.Nil(t, err)
			assert.Len(t, data, tc.WantCount)
		})
	}
}

func TestRepository_FindOne(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	ocr := NewRepository(orm)

	// No data
	// Act
	oauthClientId := "test_not_found"
	res, err := ocr.FindOne(&model.OauthClient{Id: oauthClientId})

	// Assert
	assert.Nil(t, err)
	assert.Nil(t, res)

	// Has data
	// Act
	oauthClientId = "address-book-go"
	res, err = ocr.FindOne(&model.OauthClient{Id: oauthClientId})

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, oauthClientId, res.Id)
}

func TestRepository_Count(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	ocr := NewRepository(orm)

	// Act
	count, err := ocr.Count()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 2, count)
}

func TestRepository_Update(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	ocr := NewRepository(orm)

	oauthClientId := "address-book-go"
	cli := model.OauthClient{Id: oauthClientId}
	_, _ = orm.Get(&cli)

	m := model.OauthClient{
		Id:           oauthClientId,
		SysAccountId: 2,
		Name:         "test_name",
		Secret:       "test_secret",
		Domain:       "test_domain",
		Scope:        "test_scope",
		IconPath:     "test_icon_path",
		Data:         "",
	}

	// Act
	err := ocr.Update(&m)

	// Assert
	assert.Nil(t, err)

	// Teardown
	_, _ = orm.Where("id = ? ", oauthClientId).Update(&cli)
}

func TestRepository_Delete(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	ocr := NewRepository(orm)

	cli := model.OauthClient{
		Id:           "test_id",
		SysAccountId: 1,
		Name:         "test_name",
		Secret:       "test_secret",
		Domain:       "test_domain",
		Scope:        "test_scope",
		IconPath:     "test_icon_path",
		Data:         "",
	}
	_, _ = orm.Insert(&cli)

	// Act
	err := ocr.Delete(cli.Id)

	// Assert
	assert.Nil(t, err)
}
