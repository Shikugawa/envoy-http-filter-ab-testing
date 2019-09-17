package server

import (
	"database/sql"

	"github.com/gomodule/redigo/redis"
)

type Credentials struct {
	Username string `json:username`
	Password string `json:password`
}

type WelcomeResponse struct {
	Color string `json:color`
}

var DBConnectopn *sql.DB
var RedisConnectopn redis.Conn
