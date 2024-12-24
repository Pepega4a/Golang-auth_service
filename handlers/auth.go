package handlers

import (
	"auth_service/models"
	"auth_service/services"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateTokens(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "Не указан user_id", http.StatusBadRequest)
			return
		}

		ip := r.RemoteAddr

		accessToken, err := services.GenerateAccessToken(userID, ip)
		if err != nil {
			http.Error(w, "Не удалось создать Access Token", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		refreshToken, refreshTokenHash, err := services.GenerateRefreshToken(ip)
		if err != nil {
			http.Error(w, "Не удалось создать Refresh Token", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		err = models.StoreRefreshToken(db, userID, refreshTokenHash, ip)
		if err != nil {
			http.Error(w, "Ошибка сохранения токена", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		response := TokenPair{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func RefreshTokens(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tokens TokenPair

		if err := json.NewDecoder(r.Body).Decode(&tokens); err != nil {
			http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
			log.Println(err)
			return
		}

		if err := services.HasEmptyFields(tokens); err {
			http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
			log.Println(err)
			return
		}

		claims, err := services.ValidateAccessToken(tokens.AccessToken)
		if err != nil {
			http.Error(w, "Неверный Access Token", http.StatusUnauthorized)
			log.Println(err)
			return
		}

		storedToken, err := models.GetStoredRefreshToken(db, claims.UserID)
		if err != nil {
			http.Error(w, "Ошибка верификации Refresh Token", http.StatusUnauthorized)
			log.Println(err)
			return
		}

		err = services.VerifyRefreshToken(storedToken.Hash, tokens.RefreshToken)
		if err != nil {
			http.Error(w, "Неверный Refresh Token", http.StatusUnauthorized)
			log.Println(err)
			return
		}

		newAccessToken, err := services.GenerateAccessToken(claims.UserID, r.RemoteAddr)
		if err != nil {
			http.Error(w, "Ошибка генерации Access Token", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		newRefreshToken, newRefreshTokenHash, err := services.GenerateRefreshToken(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Ошибка генерации Refresh Token", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		err = models.UpdateRefreshToken(db, claims.UserID, newRefreshTokenHash, r.RemoteAddr)
		if err != nil {
			http.Error(w, "Ошибка обновления Refresh Token", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		response := TokenPair{
			AccessToken:  newAccessToken,
			RefreshToken: newRefreshToken,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
