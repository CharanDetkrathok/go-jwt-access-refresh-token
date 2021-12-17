package databaseConnection

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type connection struct{}

func NewDatabaseConnection() *connection {

	return &connection{}

}

func (c *connection) RedisConnection() *redis.Client {

	return redis.NewClient(&redis.Options{
		Addr:     viper.GetString("rdb.address"),
		Password: viper.GetString("rdb.password"),
		DB:       viper.GetInt("rdb.db"),
	})

}

func (c *connection) OracleConnection() (*sqlx.DB, error) {

	// PRODUCTION DATABASE CONNECTION 
	dns := fmt.Sprintf("%v", viper.GetString("db.connection"))
	driver := viper.GetString("db.openDriver")

	return sqlx.Open(driver, dns)

}
