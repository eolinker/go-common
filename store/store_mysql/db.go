/*
 * Copyright (c) 2023. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package store_mysql

import (
	"context"
	slog "log"
	"os"
	"time"

	"github.com/eolinker/go-common/autowire"
	"github.com/eolinker/go-common/cftool"
	"github.com/eolinker/go-common/store"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	_ store.IDB = (*storeDB)(nil)
)

type storeDB struct {
	db *gorm.DB
}
type mysqlInit struct {
	config *DBConfig `autowired:""`
}

var _ store.IDB = (*storeDB)(nil)

func init() {
	cftool.Register[DBConfig]("mysql")
	autowire.Autowired(new(mysqlInit))

}

func (m *storeDB) DB(ctx context.Context) *gorm.DB {
	if ctx == nil {
		return m.db.WithContext(context.Background())
	}
	if tx, ok := ctx.Value(store.TxContextKey).(*gorm.DB); ok {
		return tx
	}
	return m.db.WithContext(ctx)
}
func (m *storeDB) IsTxCtx(ctx context.Context) bool {
	if _, ok := ctx.Value(store.TxContextKey).(*gorm.DB); ok {
		return ok
	}
	return false
}

func (m *mysqlInit) OnComplete() {
	m.InitDb()
}
func (m *mysqlInit) InitDb() {
	dialector := mysql.Open(m.config.getDBNS())
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.New(slog.New(os.Stderr, "\r\n", slog.LstdFlags), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		}),
	})
	if err != nil {
		slog.Fatal(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		slog.Fatal(err)
	}
	sqlDb.SetConnMaxLifetime(time.Second * 9)
	sqlDb.SetMaxOpenConns(200)
	sqlDb.SetMaxIdleConns(200)

	autowire.Autowired[store.IDB](&storeDB{db: db})

}
