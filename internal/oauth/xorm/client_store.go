package xorm

import (
	"context"
	"fmt"
	"oauth2-server-go/dto/model"
	"time"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	jsoniter "github.com/json-iterator/go"
	"xorm.io/xorm"
)

// NewClientStore creates xorm mysql store instance
func NewClientStore(orm *xorm.EngineGroup, options ...ClientStoreOption) (*ClientStore, error) {
	store := &ClientStore{
		orm:          orm,
		tableName:    "oauth2_client",
		maxLifetime:  time.Hour * 2,
		maxOpenConns: 50,
		maxIdleConns: 25,
	}

	for _, o := range options {
		o(store)
	}

	var err error
	if !store.initTableDisabled {
		err = store.initTable()
	}

	if err != nil {
		return store, err
	}

	store.orm.SetMaxOpenConns(store.maxOpenConns)
	store.orm.SetMaxIdleConns(store.maxIdleConns)
	store.orm.SetConnMaxLifetime(store.maxLifetime)

	return store, err
}

type ClientStore struct {
	orm               *xorm.EngineGroup
	tableName         string
	initTableDisabled bool
	maxLifetime       time.Duration
	maxOpenConns      int
	maxIdleConns      int
}

func (s *ClientStore) initTable() error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id VARCHAR(255) NOT NULL PRIMARY KEY,
		secret VARCHAR(255) NOT NULL,
		domain VARCHAR(255) NOT NULL,
		data TEXT NOT NULL
	  );
`, s.tableName)

	_, err := s.orm.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *ClientStore) toClientInfo(data string) (oauth2.ClientInfo, error) {
	var cm models.Client
	err := jsoniter.Unmarshal([]byte(data), &cm)
	return &cm, err
}

// GetByID retrieves and returns client information by id
func (s *ClientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	if id == "" {
		return nil, nil
	}

	item := model.OauthClient{
		Id: id,
	}

	_, err := s.orm.Table(s.tableName).Get(&item)
	if err != nil {
		return nil, err
	}

	return s.toClientInfo(item.Data)
}

// Create creates and stores the new client information
func (s *ClientStore) Create(info oauth2.ClientInfo) error {
	data, err := jsoniter.Marshal(info)
	if err != nil {
		return err
	}

	_, err = s.orm.Exec(fmt.Sprintf("INSERT INTO %s (id, secret, domain, data) VALUES (?,?,?,?)", s.tableName),
		info.GetID(),
		info.GetSecret(),
		info.GetDomain(),
		string(data),
	)
	if err != nil {
		return err
	}

	return nil
}
