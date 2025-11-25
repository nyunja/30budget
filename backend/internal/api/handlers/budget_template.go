package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nyunja/30budget/backend/internal/config"
	"go.uber.org/zap"
)

type BudgetTemplateHandler struct {
	dbPool *pgxpool.Pool
	config *config.Config
	logger *zap.Logger
}

func NewBudgetTemplateHandler(dbPool *pgxpool.Pool, cfg *config.Config, logger *zap.Logger) *BudgetTemplateHandler {
	return &BudgetTemplateHandler{
		dbPool: dbPool,
		config: cfg,
		logger: logger,
	}
}

func (h *BudgetTemplateHandler) CreateBudgetTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "CreateBudgetTemplate not implemented"})
}

func (h *BudgetTemplateHandler) GetBudgetTemplateByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "GetBudgetTemplateByID not implemented"})
}

func (h *BudgetTemplateHandler) ListBudgetTemplatesByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "ListBudgetTemplatesByUserID not implemented"})
}

func (h *BudgetTemplateHandler) UpdateBudgetTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "UpdateBudgetTemplate not implemented"})
}

func (h *BudgetTemplateHandler) DeleteBudgetTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"message": "DeleteBudgetTemplate not implemented"})
}
