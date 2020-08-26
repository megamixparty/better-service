package auth

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"customer/api/middlewares"

	"golang.org/x/crypto/bcrypt"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	u := new(LoginRequest)
	json.NewDecoder(r.Body).Decode(&u)
	res := LoginResponse{}
	id, hashedPassword, err := loginUser(middlewares.GetDB(r.Context()), u.User, u.Pass)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		res.Error = "Invalid user or password"
		b, _ := json.Marshal(res)
		w.Write(b)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Pass))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		res.Error = "Invalid user or password"
		b, _ := json.Marshal(res)
		w.Write(b)
		return
	}

	token := GenerateToken(32)
	middlewares.GetRedis(r.Context()).Set(token, id, 5*time.Minute)
	res.Token = token
	b, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	u := new(RegisterRequest)
	json.NewDecoder(r.Body).Decode(&u)
	res := RegisterResponse{}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Pass), bcrypt.MinCost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = "Failed when register, please try again"
		b, _ := json.Marshal(res)
		w.Write(b)
		return
	}

	err = registerUser(middlewares.GetDB(r.Context()), u.User, string(hashedPassword))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = "Failed when register, please try again"
		b, _ := json.Marshal(res)
		w.Write(b)
		return
	}

	res.Message = "Account has been created"
	b, _ := json.Marshal(res)
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func GenerateToken(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
