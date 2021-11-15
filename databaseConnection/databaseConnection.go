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

	/*
		// สำหรับ RUN TEST
		dns := `user="scenter01" password="scenter01new" connectString="10.2.1.98:1571/RUBRAM?expire_timconnect_time=2" timezone="local"`
		driver := "godror"
	*/

	dns := fmt.Sprintf("%v", viper.GetString("db.connection"))
	driver := viper.GetString("db.openDriver")

	return sqlx.Open(driver, dns)

}
