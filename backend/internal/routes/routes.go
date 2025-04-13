package routes

import (
    "github.com/gin-gonic/gin"
    "fintrack-api/internal/controllers"
)

func SetupRoutes(router *gin.Engine, expenseController *controllers.ExpenseController) {
    router.POST("/expenses", expenseController.CreateExpense)
    router.GET("/expenses", expenseController.GetExpenses)
    router.GET("/report", expenseController.GenerateReport)
    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "ok",
        })
    })
}