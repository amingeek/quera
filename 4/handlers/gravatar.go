package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/mail"
	"strings"
)

type GravatarResponse struct {
	Ok          bool   `json:"ok"`
	GravatarUrl string `json:"gravatar_url"`
}

type ErrorResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

var (
	BaseUrl = "https://www.gravatar.com/avatar/"
)

func IsEmailValid(email string) bool {
	if email == "" {
		return false
	}
	_, err := mail.ParseAddress(email)
	return err == nil
}

func HashMail(email string) string {
	email = strings.ToLower(email)
	h := md5.New()
	h.Write([]byte(email))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
func RequestGravatar(email string) string {
	if !IsEmailValid(email) {
		return ""
	}
	hash := HashMail(email)
	return BaseUrl + hash
}

func HandleGravatarRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	email := r.URL.Query().Get("email")
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Ok:      false,
			Message: "No email address provided",
		})
		return
	}
	if !IsEmailValid(email) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Ok:      false,
			Message: "Invalid email address",
		})
		return
	}

	gravatarUrl := RequestGravatar(email)
	json.NewEncoder(w).Encode(GravatarResponse{
		Ok:          true,
		GravatarUrl: gravatarUrl,
	})
}
