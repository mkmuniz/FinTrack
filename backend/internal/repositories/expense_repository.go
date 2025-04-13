package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"fintrack-api/internal/models"
)

type ExpenseRepository struct {
	collection *mongo.Collection
}

func NewExpenseRepository(db *mongo.Database, collectionName string) *ExpenseRepository {
	return &ExpenseRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *ExpenseRepository) SaveExpense(expense *models.Expense) error {
	_, err := r.collection.InsertOne(context.Background(), expense)
	return err
}

func (r *ExpenseRepository) FindAllExpenses() ([]models.Expense, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var expenses []models.Expense
	for cursor.Next(context.Background()) {
		var expense models.Expense
		if err := cursor.Decode(&expense); err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func (r *ExpenseRepository) GetExpenseByID(id string) (*models.Expense, error) {
	var expense models.Expense
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&expense)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &expense, nil
}

func (r *ExpenseRepository) FindExpensesByDateRange(ctx context.Context, startDate, endDate time.Time) ([]models.Expense, error) {
	// Simulação de busca no banco de dados
	return []models.Expense{
		{ID: "1", Amount: 100.0, Category: "Food", Date: "2025-04-01"},
		{ID: "2", Amount: 200.0, Category: "Transport", Date: "2025-04-02"},
	}, nil
}