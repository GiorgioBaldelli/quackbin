package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"quackbin/internal/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HandlePaste(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handling paste creation request from IP: %s", GetClientIP(r))
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var paste models.Paste
		err := json.NewDecoder(r.Body).Decode(&paste)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		paste.ID = uuid.New().String()

		var passwordHash []byte
		if paste.IsPrivate && paste.Password != "" {
			passwordHash, err = bcrypt.GenerateFromPassword([]byte(paste.Password), bcrypt.DefaultCost)
			if err != nil {
				http.Error(w, "Error hashing password", http.StatusInternalServerError)
				return
			}
		}

		_, err = db.Exec("INSERT INTO pastes (id, content, is_private, password_hash) VALUES (?, ?, ?, ?)",
			paste.ID, paste.Content, paste.IsPrivate, passwordHash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		paste.Password = "" // Don't send the password back to the client
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(paste)
	}
}

func GetPaste(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handling paste access request from IP: %s", GetClientIP(r))
		if r.Method != http.MethodGet && r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Path[len("/api/paste/"):]

		var paste models.Paste
		var passwordHash []byte
		err := db.QueryRow("SELECT id, content, is_private, password_hash FROM pastes WHERE id = ?", id).
			Scan(&paste.ID, &paste.Content, &paste.IsPrivate, &passwordHash)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Paste not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if paste.IsPrivate {
			if r.Method == http.MethodGet {
				http.Error(w, "Password required", http.StatusUnauthorized)
				return
			}

			var providedPassword struct {
				Password string `json:"password"`
			}
			err = json.NewDecoder(r.Body).Decode(&providedPassword)
			if err != nil {
				http.Error(w, "Invalid password format", http.StatusBadRequest)
				return
			}

			err = bcrypt.CompareHashAndPassword(passwordHash, []byte(providedPassword.Password))
			if err != nil {
				http.Error(w, "Incorrect password", http.StatusUnauthorized)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(paste)
	}
}
