package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type User struct {
	ID    string
	Name  string
	Email string
	Phone string
}

// SplitStrategy
type SplitStrategy interface {
	CalculateShare(paidBy string, amount float64, users []string) (map[string]float64, error)
}

// EqualSplit
type EqualSplit struct{}

func (es *EqualSplit) CalculateShare(paidBy string, amount float64, users []string) (map[string]float64, error) {
	n := len(users)
	if n == 0 {
		return nil, errors.New("no users provided")
	}

	share := roundToTwoDecimals(amount / float64(n))
	shares := make(map[string]float64)
	for _, user := range users {
		if user != paidBy {
			shares[user] = share
		}
	}
	return shares, nil
}

// ExactSplit
type ExactSplit struct {
	Amounts []float64
}

func (es *ExactSplit) CalculateShare(paidBy string, amount float64, users []string) (map[string]float64, error) {
	if len(users) != len(es.Amounts) {
		return nil, errors.New("number of users and amounts do not match")
	}

	total := 0.0
	for _, a := range es.Amounts {
		total += a
	}
	if math.Abs(total-amount) > 0.01 {
		return nil, errors.New("sum of amounts does not match total expense")
	}

	shares := make(map[string]float64)
	for i, user := range users {
		if user != paidBy {
			shares[user] = es.Amounts[i]
		}
	}
	return shares, nil
}

// PercentageSplit
type PercentageSplit struct {
	Percentages []float64
}

func (ps *PercentageSplit) CalculateShare(paidBy string, amount float64, users []string) (map[string]float64, error) {
	if len(users) != len(ps.Percentages) {
		return nil, errors.New("number of users and percentages do not match")
	}

	total := 0.0
	for _, p := range ps.Percentages {
		total += p
	}
	if math.Abs(total-100) > 0.01 {
		return nil, errors.New("total percentage must equal 100")
	}

	shares := make(map[string]float64)
	for i, user := range users {
		if user != paidBy {
			shares[user] = roundToTwoDecimals((ps.Percentages[i] / 100) * amount)
		}
	}
	return shares, nil
}

// ExpenseManager
type ExpenseManager struct {
	users    map[string]*User
	balances map[string]map[string]float64
}

func NewExpenseManager() *ExpenseManager {
	return &ExpenseManager{
		users:    make(map[string]*User),
		balances: make(map[string]map[string]float64),
	}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// generateID generates a random string of a specified length.
func generateID(length int) string {
	// Seed the random number generator to produce different output each time.
	rand.Seed(time.Now().UnixNano())
	id := make([]byte, length)
	for i := range id {
		id[i] = charset[rand.Intn(len(charset))]
	}
	return string(id)
}

func (em *ExpenseManager) AddUser(name, email, phone string) {
	userID := generateID(10)
	user := &User{
		ID:    userID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	em.users[userID] = user
}

// AddExpense
func (em *ExpenseManager) AddExpense(paidBy string, amount float64, users []string, strategy SplitStrategy) error {
	shares, err := strategy.CalculateShare(paidBy, amount, users)
	if err != nil {
		return err
	}

	for user, share := range shares {
		if user == paidBy {
			continue
		}
		if em.balances[user] == nil {
			em.balances[user] = make(map[string]float64)
		}
		if em.balances[paidBy] == nil {
			em.balances[paidBy] = make(map[string]float64)
		}
		em.balances[user][paidBy] += share
		em.balances[paidBy][user] -= share
	}
	return nil
}

// ShowBalances
func (em *ExpenseManager) ShowBalances(userID string) {
	if userID == "" {
		for u1, balances := range em.balances {
			for u2, amount := range balances {
				if amount > 0 {
					fmt.Printf("%s owes %s: %.2f\n", u1, u2, amount)
				} else if amount < 0 {
					fmt.Printf("%s owes %s: %.2f\n", u2, u1, -amount)
				}
			}
		}
	} else {
		balances, exists := em.balances[userID]
		if !exists || len(balances) == 0 {
			fmt.Println("No balances")
			return
		}
		for u2, amount := range balances {
			if amount > 0 {
				fmt.Printf("%s owes %s: %.2f\n", userID, u2, amount)
			} else if amount < 0 {
				fmt.Printf("%s owes %s: %.2f\n", u2, userID, -amount)
			}
		}
	}
}

func roundToTwoDecimals(amount float64) float64 {
	return math.Round(amount*100) / 100
}

func main() {

	em := NewExpenseManager()
	em.AddUser("User1", "user1@example.com", "1234567890")
	em.AddUser("User2", "user2@example.com", "0987654321")
	em.AddUser("User3", "user3@example.com", "1122334455")
	em.AddUser("User4", "user4@example.com", "5566778899")

	fmt.Println("Initial Balances:")
	em.ShowBalances("")

	em.AddExpense("u1", 1000, []string{"u1", "u2", "u3", "u4"}, &EqualSplit{})
	fmt.Println("\nAfter First Transaction (Equal Split):")
	em.ShowBalances("")

	em.AddExpense("u1", 1250, []string{"u2", "u3"}, &ExactSplit{Amounts: []float64{370, 880}})
	fmt.Println("\nAfter Second Transaction (Exact Split):")
	em.ShowBalances("")

	em.AddExpense("u4", 1200, []string{"u1", "u2", "u3", "u4"}, &PercentageSplit{Percentages: []float64{40, 20, 20, 20}})
	fmt.Println("\nAfter Third Transaction (Percentage Split):")
	em.ShowBalances("u1")
}
