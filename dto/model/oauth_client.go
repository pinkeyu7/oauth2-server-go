package model

import "time"

type OauthClient struct {
	Id        string    `xorm:"not null default '' comment('id') VARCHAR(255)" json:"id"`
	Secret    string    `xorm:"not null default '' comment('secret') VARCHAR(255)" json:"secret"`
	Domain    string    `xorm:"not null default '' comment('domain') VARCHAR(255)" json:"domain"`
	Data      string    `xorm:"not null default '' comment('data') TEXT" json:"data"`
	CreatedAt time.Time `xorm:"not null created DATETIME" json:"created_at"`
	UpdatedAt time.Time `xorm:"not null updated DATETIME" json:"updated_at"`
}
