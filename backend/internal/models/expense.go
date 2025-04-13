package models

type Expense struct {
    ID       string  `bson:"_id,omitempty"`
    Amount   float64 `bson:"amount"`
    Category string  `bson:"category"`
    Date     string  `bson:"date"`
}