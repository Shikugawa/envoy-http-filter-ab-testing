package server

import (
	"database/sql"

	"github.com/go-redis/redis"
)

type Credentials struct {
	Username string `json:username`
	Password string `json:password`
}

type WelcomeResponse struct {
	Color string `json:color`
}

type WelcomeRequest struct {
	SessionId string `json:session_id`
}

var DBConnectopn *sql.DB
var RedisClient redis.Client
