package Handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	Dto "nutecth/Dto/Result"
	userDto "nutecth/Dto/User"
	"nutecth/Models"
	jwtToken "nutecth/Pkg/jwt"
	"nutecth/Repositories"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handleruser struct {
	UserRepository Repositories.UserRepository
}

func HandlerUser(UserRepository Repositories.UserRepository) *handleruser {
	return &handleruser{UserRepository}
}

func (h *handleruser) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var user Models.User

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := Dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := Dto.SuccessResult{Code: http.StatusOK, Data: user}
	json.NewEncoder(w).Encode(response)
}

func (h *handleruser) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")

	request := new(userDto.UserRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := Dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user := Models.User{
		Username: request.Username,
		Password: request.Password,
	}

	user, err := h.UserRepository.Login(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := Dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Password != user.Password {
		w.WriteHeader(http.StatusBadRequest)
		response := Dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong username or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
		return
	}

	loginResponse := userDto.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}

	w.Header().Set("Content-Type", "application/json")
	response := Dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handleruser) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	user, err := h.UserRepository.GetUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := Dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	CheckAuthResponse := userDto.CheckAuthResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	w.Header().Set("Content-Type", "application/json")
	response := Dto.SuccessResult{Code: http.StatusOK, Data: CheckAuthResponse}
	json.NewEncoder(w).Encode(response)
}
