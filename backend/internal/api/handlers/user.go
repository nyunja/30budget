package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nyunja/30budget/backend/internal/config"
	"go.uber.org/zap"
)

type UserHandler struct {
	dbPool *pgxpool.Pool
	config *config.Config
	logger *zap.Logger
}

func NewUserHandler(dbPool *pgxpool.Pool, cfg *config.Config, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		dbPool: dbPool,
		config: cfg,
		logger: logger,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "CreateUser not implemented"})
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "GetUserByID not implemented"})
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "UpdateUser not implemented"})
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "DeleteUser not implemented"})
}
