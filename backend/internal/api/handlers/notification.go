package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nyunja/30budget/backend/internal/config"
	"go.uber.org/zap"
)

type NotificationHandler struct {
	dbPool *pgxpool.Pool
	config *config.Config
	logger *zap.Logger
}

func NewNotificationHandler(dbPool *pgxpool.Pool, cfg *config.Config, logger *zap.Logger) *NotificationHandler {
	return &NotificationHandler{
		dbPool: dbPool,
		config: cfg,
		logger: logger,
	}
}

func (h *NotificationHandler) CreateNotification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "CreateNotification not implemented"})
}

func (h *NotificationHandler) GetNotificationByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "GetNotificationByID not implemented"})
}

func (h *NotificationHandler) ListNotificationsByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "ListNotificationsByUserID not implemented"})
}

func (h *NotificationHandler) UpdateNotification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "UpdateNotification not implemented"})
}

func (h *NotificationHandler) DeleteNotification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "DeleteNotification not implemented"})
}
