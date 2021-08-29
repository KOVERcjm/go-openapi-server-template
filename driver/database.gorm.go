package driver

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils"
	. "kovercheng/middleware"
	tables "kovercheng/model/table"
	"time"
)

type logger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
}

func newLogger() *logger {
	return &logger{
		SkipErrRecordNotFound: true,
	}
}

func (l *logger) LogMode(gormLogger.LogLevel) gormLogger.Interface {
	return l
}

func (l *logger) Info(ctx context.Context, s string, args ...interface{}) {
	Logger.WithContext(ctx).Infof(s, args)
}

func (l *logger) Warn(ctx context.Context, s string, args ...interface{}) {
	Logger.WithContext(ctx).Warnf(s, args)
}

func (l *logger) Error(ctx context.Context, s string, args ...interface{}) {
	Logger.WithContext(ctx).Errorf(s, args)
}

func (l *logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := logrus.Fields{}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[logrus.ErrorKey] = err
		Logger.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		Logger.WithContext(ctx).WithFields(fields).Warnf("%s [%s]", sql, elapsed)
		return
	}
	Logger.WithContext(ctx).WithFields(fields).Debugf("%s [%s]", sql, elapsed)
}

func newGormConnection(connectUrl string) error {
	client, _ := gorm.Open(postgres.Open(connectUrl), &gorm.Config{
		Logger: newLogger(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})
	sqlDB, _ := client.DB()
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	if err := sqlDB.Ping(); err != nil {
		_ = closeGormConnection(client)
		Logger.Fatal(err.Error())
		return fmt.Errorf("%s", "cannot ping to Postgres database")
	}

	if err := client.AutoMigrate(&tables.Test{}); err != nil {
		_ = closeGormConnection(client)
		Logger.Fatal(err.Error())
		return fmt.Errorf("%s", "cannot migrate Postgres schema")
	}

	database.Postgres = client
	return nil
}

func closeGormConnection(client *gorm.DB) error {
	sqlDB, _ := client.DB()
	err := sqlDB.Close()
	if err != nil {
		Logger.Fatal(err.Error())
		return fmt.Errorf("%s", "cannot close Postgres connection")
	}
	return nil
}
