package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "fintrack-api/config"
    "fintrack-api/internal/controllers"
    "fintrack-api/internal/repositories"
    "fintrack-api/internal/routes"
    "fintrack-api/internal/services"
)

func main() {
    db := config.GetDatabase()

    expenseRepo := repositories.NewExpenseRepository(db, "expenses")
    expenseService := services.NewExpenseService(expenseRepo)
    expenseController := controllers.NewExpenseController(expenseService)

    router := gin.Default()
    routes.SetupRoutes(router, expenseController)

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Could not start server: %v", err)
    }
}