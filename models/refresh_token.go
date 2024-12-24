package models

import (
	"database/sql"
)

func StoreRefreshToken(db *sql.DB, userID, hash, ip string) error {
	_, err := db.Exec("INSERT INTO refresh_tokens (user_id, token_hash, ip) VALUES ($1, $2, $3)", userID, hash, ip)
	return err
}

func GetStoredRefreshToken(db *sql.DB, userID string) (RefreshToken, error) {
	var token RefreshToken
	err := db.QueryRow("SELECT token_hash, ip FROM refresh_tokens WHERE user_id = $1", userID).Scan(&token.Hash, &token.IP)
	return token, err
}

func UpdateRefreshToken(db *sql.DB, userID, newHash, ip string) error {
	_, err := db.Exec("UPDATE refresh_tokens SET token_hash = $1, ip = $2 WHERE user_id = $3", newHash, ip, userID)
	return err
}

type RefreshToken struct {
	Hash string
	IP   string
}
