package services

import (
	"context"
	"time"
	"fintrack-api/internal/models"
	"fintrack-api/internal/repositories"
)

type ExpenseService struct {
	repo *repositories.ExpenseRepository
}

func NewExpenseService(repo *repositories.ExpenseRepository) *ExpenseService {
	return &ExpenseService{repo: repo}
}

func (s *ExpenseService) CalculateTotalExpenses(ctx context.Context, startDate, endDate time.Time) (float64, error) {
	expenses, err := s.repo.FindExpensesByDateRange(ctx, startDate, endDate)
	if err != nil {
		return 0, err
	}

	total := 0.0
	for _, expense := range expenses {
		total += expense.Amount
	}
	return total, nil
}

func (s *ExpenseService) GetHighestExpenseCategory(ctx context.Context, startDate, endDate time.Time) (string, error) {
	expenses, err := s.repo.FindExpensesByDateRange(ctx, startDate, endDate)
	if err != nil {
		return "", err
	}

	categoryTotals := make(map[string]float64)
	for _, expense := range expenses {
		categoryTotals[expense.Category] += expense.Amount
	}

	var highestCategory string
	var highestAmount float64
	for category, amount := range categoryTotals {
		if amount > highestAmount {
			highestAmount = amount
			highestCategory = category
		}
	}
	return highestCategory, nil
}

func (s *ExpenseService) GetLowestExpenseCategory(ctx context.Context, startDate, endDate time.Time) (string, error) {
	expenses, err := s.repo.FindExpensesByDateRange(ctx, startDate, endDate)
	if err != nil {
		return "", err
	}

	categoryTotals := make(map[string]float64)
	for _, expense := range expenses {
		categoryTotals[expense.Category] += expense.Amount
	}

	var lowestCategory string
	var lowestAmount float64
	for category, amount := range categoryTotals {
		if lowestAmount == 0 || amount < lowestAmount {
			lowestAmount = amount
			lowestCategory = category
		}
	}
	return lowestCategory, nil
}

func (s *ExpenseService) TrackMonthlyTrends(ctx context.Context, year int) (map[int]float64, error) {
	trends := make(map[int]float64)
	for month := 1; month <= 12; month++ {
		startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
		endDate := startDate.AddDate(0, 1, 0)
		total, err := s.CalculateTotalExpenses(ctx, startDate, endDate)
		if err != nil {
			return nil, err
		}
		trends[month] = total
	}
	return trends, nil
}

func (s *ExpenseService) SaveExpense(expense *models.Expense) error {
	return s.repo.SaveExpense(expense)
}

func (s *ExpenseService) GetAllExpenses() ([]models.Expense, error) {
	return s.repo.FindAllExpenses()
}

func (s *ExpenseService) GenerateReport() (map[string]interface{}, error) {
	expenses, err := s.repo.FindAllExpenses()
	if err != nil {
		return nil, err
	}

	total := 0.0
	categoryTotals := make(map[string]float64)
	for _, expense := range expenses {
		total += expense.Amount
		categoryTotals[expense.Category] += expense.Amount
	}

	report := map[string]interface{}{
		"total":           total,
		"category_totals": categoryTotals,
	}
	return report, nil
}