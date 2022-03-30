package seed

import (
	"oauth2-server-go/dto/model"

	"github.com/brianvoe/gofakeit/v4"
	"xorm.io/xorm"
)

func CreateUser(engine *xorm.Engine, account, name, email, phone string) error {
	defaultPassword := "0eb683eacea7957d8b4140ed837f1ee7fce60ba74e48839a51d6b2085938b49b"

	con := model.User{
		Account:  account,
		Phone:    phone,
		Email:    email,
		Password: defaultPassword,
		Name:     name,
	}

	_, err := engine.Insert(&con)
	return err
}

func AllUser() []Seed {
	return []Seed{
		{
			Name: "Create User - 1",
			Run: func(engine *xorm.Engine) error {
				err := CreateUser(engine, "user_1", gofakeit.Name(), gofakeit.Email(), gofakeit.Phone())
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
}
