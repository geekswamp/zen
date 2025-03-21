package postgres

import (
	"fmt"
	"net/url"
	"time"

	"github.com/geekswamp/zen/configs"
	"github.com/geekswamp/zen/internal/errors"
	"github.com/geekswamp/zen/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
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
	var logMode gormLog.LogLevel

	p := config.Postgres
	base := p.Base

	if config.App.Mode == "debug" {
		logMode = gormLog.Info
	} else {
		logMode = gormLog.Silent
	}

	db, err := gorm.Open(postgres.Open(buildDsn(config)), &gorm.Config{
		Logger: gormLog.Default.LogMode(logMode),
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
