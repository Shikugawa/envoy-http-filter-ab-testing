package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Shikugawa/envoy-http-filter-ab-testing/server"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var resp *server.WelcomeResponse

	if r.Method == http.MethodPost {
		var welcome_req server.WelcomeRequest
		if err := json.NewDecoder(r.Body).Decode(&welcome_req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(welcome_req.SessionId) != 0 {
			if err := server.RedisClient.Get(welcome_req.SessionId).Err(); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			resp = &server.WelcomeResponse{
				Color: "green",
			}
		} else {
			resp = &server.WelcomeResponse{
				Color: "none",
			}
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	byte_resp, _ := json.Marshal(resp)
	w.Write(byte_resp)
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: server.RedisHost + ":" + server.RedisPort,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	server.RedisClient = *client
	db, err := sql.Open("mysql", server.MySQLUserName+"@tcp("+server.MySQLHost+":"+server.MySQLPort+")/"+server.MySQLDBName)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	server.DBConnectopn = db
	http.HandleFunc("/welcome", welcomeHandler)
	http.HandleFunc("/login", server.LoginHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
