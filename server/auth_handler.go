package server

import (
	"crypto/md5"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var cred Credentials
	if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	rows, err := DBConnectopn.Query("SELECT FROM Users WHERE UserName = ?", cred.Username)
	defer rows.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var _id int
	var _user string
	var _md5_hashed_password string
	for rows.Next() {
		if err := rows.Scan(&_id, &_user, &_md5_hashed_password); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	if &_id == nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	hashed_user_password := md5.Sum([]byte(cred.Password))
	if _md5_hashed_password == string(hashed_user_password[:]) {
		sessionId, err := uuid.NewRandom()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		sessionIdString := sessionId.String()
		RedisConnectopn.Do("SETEX", sessionIdString, "3600", cred.Username) // 1時間後にセッション切る
		http.SetCookie(w, &http.Cookie{
			Name:    "session_id",
			Value:   sessionIdString,
			Expires: time.Now().Add(1 * time.Hour),
		})
		http.SetCookie(w, &http.Cookie{
			Name:    "user_id",
			Value:   string(_id),
			Expires: time.Now().Add(1 * time.Hour),
		})
		w.WriteHeader(http.StatusAccepted)
		return
	}
}
