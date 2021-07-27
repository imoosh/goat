package orm

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm/logger"
    "time"

    "goat/ecode"
    "goat/log"
    xtime "goat/time"

    "gorm.io/gorm"
)

// Config mysql config.
type Config struct {
    DSN         string         // data source name.
    Active      int            // pool
    Idle        int            // pool
    IdleTimeout xtime.Duration // connect max life time.
}

type ormLog struct{}

func (l ormLog) Printf(format string, v ...interface{}) {
    log.Infof(format, v...)
}

func init() {
    gorm.ErrRecordNotFound = ecode.NothingFound
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *gorm.DB) {
    db, err := gorm.Open(mysql.Open(c.DSN), nil)
    if err != nil {
        log.Error(err)
        panic(err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        log.Error(err)
        panic(err)
    }
    sqlDB.SetMaxIdleConns(c.Idle)
    sqlDB.SetMaxOpenConns(c.Active)
    sqlDB.SetConnMaxLifetime(time.Duration(c.IdleTimeout) / time.Second)

    // 设置日志
    db.Config.Logger = logger.New(&ormLog{}, logger.Config{
        SlowThreshold: 200 * time.Millisecond,
        LogLevel:      logger.Error,
        Colorful:      false,
    })

    return
}
