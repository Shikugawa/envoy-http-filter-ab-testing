package blue

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Shikugawa/envoy-http-filter-ab-testing/server"
	"github.com/gomodule/redigo/redis"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	resp := &server.WelcomeResponse{
		Color: "blue",
	}
	byte_resp, _ := json.Marshal(resp)
	w.Write(byte_resp)
}

func main() {
	conn, err := redis.DialURL("redis://localhost")
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	server.RedisConnectopn = conn
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/EnvoyABTesting")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	server.DBConnectopn = db
	http.HandleFunc("/welcome", welcomeHandler)
	http.HandleFunc("/login", server.LoginHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
