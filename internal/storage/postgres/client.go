package postgres

import (
	"fmt"
	"net/url"
	"time"

	"github.com/geekswamp/zen/configs"
	"github.com/geekswamp/zen/internal/logger"
	"github.com/geekswamp/zen/internal/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Client struct {
	db     *gorm.DB
	config configs.Config
}

func New(config configs.Config) Postgres {
	return &Client{
		config: config,
	}
}

func NewDefault() Postgres {
	return &Client{
		config: configs.Get(),
	}
}

func (d *Client) Connect() (db *gorm.DB, err error) {
	d.db, err = connect(d.config)
	if err != nil {
		return nil, err
	}

	return d.db, nil
}

func (d *Client) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func buildDsn(config configs.Config) string {
	q := url.Values{}
	p := config.Postgres
	q.Add("sslmode", p.SSLMode)

	u := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(p.User, p.Password),
		Host:     fmt.Sprintf("%s:%d", p.Address, p.Port),
		Path:     p.Name,
		RawQuery: q.Encode(),
	}

	return u.String()
}

func connect(config configs.Config) (*gorm.DB, error) {
	p := config.Postgres
	base := p.Base

	db, err := gorm.Open(postgres.Open(buildDsn(config)), &gorm.Config{
		Logger: gormLog.Default.LogMode(gormLog.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			loc, err := time.LoadLocation(p.Timezone)
			if err != nil {
				log.Fatal(errors.ErrFailedLoadLocal.Error(), logger.Local(err.Error()))
			}

			return time.Now().In(loc)
		},
	})
	if err != nil {
		log.Fatal(errors.ErrFailedToConnectDB.Error(), logger.Postgres(p.Name))
		return nil, errors.ErrFailedToConnectDB
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(base.MaxOpenConn)
	sqlDB.SetMaxIdleConns(base.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Minute * base.ConnMaxLifeTime)
	sqlDB.SetConnMaxIdleTime(base.ConnMaxIdleTime)

	return db, nil
}
