package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nyunja/30budget/backend/internal/api/handlers"
	"github.com/nyunja/30budget/backend/internal/config"
	"go.uber.org/zap"
)

// SetupRoutes initializes all API routes.
func SetupRoutes(r *chi.Mux, dbPool *pgxpool.Pool, cfg *config.Config, logger *zap.Logger) {
	// Initialize handlers
	userHandler := handlers.NewUserHandler(dbPool, cfg, logger)
	categoryHandler := handlers.NewCategoryHandler(dbPool, cfg, logger)
	transactionHandler := handlers.NewTransactionHandler(dbPool, cfg, logger)
	notificationHandler := handlers.NewNotificationHandler(dbPool, cfg, logger)
	budgetTemplateHandler := handlers.NewBudgetTemplateHandler(dbPool, cfg, logger)

	r.Route("/api/v1", func(r chi.Router) {
		// Example route
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Welcome to the 30Budget API!"}`))
		})

		// User routes
		r.Route("/users", func(r chi.Router) {
			r.Post("/", userHandler.CreateUser)
			r.Get("/{userID}", userHandler.GetUserByID)
			r.Put("/{userID}", userHandler.UpdateUser)
			r.Delete("/{userID}", userHandler.DeleteUser)
		})

		// Category routes
		r.Route("/users/{userID}/categories", func(r chi.Router) {
			r.Post("/", categoryHandler.CreateCategory)
			r.Get("/", categoryHandler.ListCategoriesByUserID)
			r.Get("/{categoryID}", categoryHandler.GetCategoryByID)
			r.Put("/{categoryID}", categoryHandler.UpdateCategory)
			r.Delete("/{categoryID}", categoryHandler.DeleteCategory)
		})

		// Transaction routes
		r.Route("/users/{userID}/transactions", func(r chi.Router) {
			r.Post("/", transactionHandler.CreateTransaction)
			r.Get("/", transactionHandler.ListTransactionsByUserID)
			r.Get("/{transactionID}", transactionHandler.GetTransactionByID)
			r.Put("/{transactionID}", transactionHandler.UpdateTransaction)
			r.Delete("/{transactionID}", transactionHandler.DeleteTransaction)
		})

		// Notification routes
		r.Route("/users/{userID}/notifications", func(r chi.Router) {
			r.Post("/", notificationHandler.CreateNotification)
			r.Get("/", notificationHandler.ListNotificationsByUserID)
			r.Get("/{notificationID}", notificationHandler.GetNotificationByID)
			r.Put("/{notificationID}", notificationHandler.UpdateNotification)
			r.Delete("/{notificationID}", notificationHandler.DeleteNotification)
		})

		// Budget Template routes
		r.Route("/users/{userID}/budget-templates", func(r chi.Router) {
			r.Post("/", budgetTemplateHandler.CreateBudgetTemplate)
			r.Get("/", budgetTemplateHandler.ListBudgetTemplatesByUserID)
			r.Get("/{templateID}", budgetTemplateHandler.GetBudgetTemplateByID)
			r.Put("/{templateID}", budgetTemplateHandler.UpdateBudgetTemplate)
			r.Delete("/{templateID}", budgetTemplateHandler.DeleteBudgetTemplate)
		})
	})
}
