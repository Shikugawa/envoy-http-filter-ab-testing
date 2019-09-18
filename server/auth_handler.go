package server

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func passwordHashing(password *string) string {
	res := md5.Sum([]byte(*password))
	return fmt.Sprintf("%x", res)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		var cred Credentials

		if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		row := DBConnectopn.QueryRow("SELECT * FROM Users WHERE UserName = ?", cred.Username)

		var _id int
		var _user string
		var _md5_hashed_password string

		if err := row.Scan(&_id, &_user, &_md5_hashed_password); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _md5_hashed_password != passwordHashing(&cred.Password) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		sessionId, err := uuid.NewRandom()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		sessionIdString := sessionId.String()
		dur, _ := time.ParseDuration("1h")
		RedisClient.Set(sessionIdString, cred.Username, dur)

		login_resp := LoginResponse{
			UserId:    strconv.Itoa(_id),
			SessionId: sessionIdString,
		}
		res, err := json.Marshal(login_resp)

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
