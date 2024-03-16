package db

import (
	"database/sql"
	"fmt"
	"github.com/geejjoo/library_repository/internal/config"
	"github.com/geejjoo/library_repository/internal/db/adapter"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

func NewSqlDB(dbConf config.DB, logger *zap.Logger) (*sqlx.DB, *adapter.SQLAdapter, error) {
	var dsn string
	var err error
	var dbRaw *sql.DB

	switch dbConf.Driver {
	case "postgres":
		//os.Getenv("APP_NAME")+"db"
		//dbConf.Host
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("APP_NAME")+"db", dbConf.Port, dbConf.User, dbConf.Password, dbConf.Name)
	case "mysql":
		cfg := mysql.NewConfig()
		cfg.Net = dbConf.Net
		cfg.Addr = dbConf.Host
		cfg.User = dbConf.User
		cfg.Passwd = dbConf.Password
		cfg.DBName = dbConf.Name
		cfg.ParseTime = true
		cfg.Timeout = time.Duration(dbConf.Timeout) * time.Second
		dsn = cfg.FormatDSN()
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	timeoutExceeded := time.After(time.Second * time.Duration(dbConf.Timeout))
	for {
		select {
		case <-timeoutExceeded:
			return nil, nil, fmt.Errorf("db connection failed after %d timeout %s", dbConf.Timeout, err)
		case <-ticker.C:
			dbRaw, err = sql.Open(dbConf.Driver, dsn)
			if err != nil {
				log.Println(err)
				return nil, nil, err
			}
			err = dbRaw.Ping()
			if err == nil {
				db := sqlx.NewDb(dbRaw, dbConf.Driver)
				sqlAdapter := adapter.NewSQLAdapter(db, dbConf)
				return db, sqlAdapter, nil
			}
			logger.Error("failed to connect to the database", zap.String("dsn", dsn), zap.Error(err))
		}

	}

}
