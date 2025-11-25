package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nyunja/30budget/backend/internal/config"
	"go.uber.org/zap"
)

type TransactionHandler struct {
	dbPool *pgxpool.Pool
	config *config.Config
	logger *zap.Logger
}

func NewTransactionHandler(dbPool *pgxpool.Pool, cfg *config.Config, logger *zap.Logger) *TransactionHandler {
	return &TransactionHandler{
		dbPool: dbPool,
		config: cfg,
		logger: logger,
	}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "CreateTransaction not implemented"})
}

func (h *TransactionHandler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "GetTransactionByID not implemented"})
}

func (h *TransactionHandler) ListTransactionsByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "ListTransactionsByUserID not implemented"})
}

func (h *TransactionHandler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "UpdateTransaction not implemented"})
}

func (h *TransactionHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "DeleteTransaction not implemented"})
}
