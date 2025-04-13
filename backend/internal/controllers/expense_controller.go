package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fintrack-api/internal/models"
	"fintrack-api/internal/services"
)

type ExpenseController struct {
    service *services.ExpenseService
}

func NewExpenseController(service *services.ExpenseService) *ExpenseController {
    return &ExpenseController{service: service}
}

func (ec *ExpenseController) CreateExpense(c *gin.Context) {
	var expense models.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ec.service.SaveExpense(&expense); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save expense"})
		return
	}
	c.JSON(http.StatusCreated, expense)
}

func (ec *ExpenseController) GetExpenses(c *gin.Context) {
	expenses, err := ec.service.GetAllExpenses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve expenses"})
		return
	}
	c.JSON(http.StatusOK, expenses)
}

func (ec *ExpenseController) GenerateReport(c *gin.Context) {
	report, err := ec.service.GenerateReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate report"})
		return
	}
	c.JSON(http.StatusOK, report)
}