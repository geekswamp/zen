package postgres

import (
	"fmt"
	"net/url"
	"time"

	"github.com/geekswamp/zen/configs"
	"github.com/geekswamp/zen/internal/logger"
	"github.com/geekswamp/zen/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	_   Postgres = (*dbPostgres)(nil)
	log          = logger.New()
)

type (
	Postgres interface {
		Close() error
		t()
	}

	dbPostgres struct{ DB *gorm.DB }
)

func New() (Postgres, error) {
	db, err := connect(configs.Get())
	if err != nil {
		return nil, err
	}

	return &dbPostgres{DB: db}, nil
}

func (d *dbPostgres) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (d *dbPostgres) t() {}

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
