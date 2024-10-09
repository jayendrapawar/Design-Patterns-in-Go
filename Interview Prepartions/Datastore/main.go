package main

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Row struct {
	ID string 
	Values map[string]string
}

type Table struct {
	Name    string
	Columns []string
	Rows    map[string]*Row
	mu      sync.RWMutex
}

type Database struct {
	Tables map[string]*Table
	mu sync.RWMutex
}


func NewDatabase() *Database {
	return &Database{
		Tables: make(map[string]*Table),
	}
}

// createTable 
func (db *Database) CreateTable(tableName string, cols []string) error {
	db.mu.Lock()
	defer db.mu.Unlock()


	if _ , exists := db.Tables[tableName]; exists{
		return errors.New("table already exists")
	}

	table := &Table{
		Name: tableName,
		Columns: cols,
		Rows: make(map[string]*Row),
	}

	db.Tables[tableName] = table
	return nil
}

func generateID() string {
	rowID := time.Now().UnixNano()
	return strconv.FormatInt(rowID, 10)
}

// InsertRow 
func (db *Database) InsertRow(tableName string, values []string) (string, error) {
	db.mu.RLock()
	table, exists := db.Tables[tableName]
	db.mu.RUnlock()


	if !exists {
		return "",errors.New("Table doesnt Exists")
	}

	if len(values) != len(table.Columns) {
		return "",errors.New("Table doesnt Exists")
	}

	row := &Row{
		ID: generateID(),
		Values: make(map[string]string),
	}

	for i, col := range table.Columns {
		row.Values[col] = values[i]
	}

	table.mu.Lock()
	table.Rows[row.ID] = row
	table.mu.Unlock()

	return row.ID,nil

}

// ReadRow

func (db *Database) ReadRow(tableName, rowID string) (*Row, error) {
	db.mu.RLock()
	table, exists := db.Tables[tableName]
	db.mu.RUnlock()

	if !exists {
		return nil, errors.New("table does not exist")
	}

	table.mu.RLock()
	row, exists := table.Rows[rowID]
	table.mu.RUnlock()

	if !exists {
		return nil, errors.New("row not found")
	}

	return row, nil
}

// update row

	func (db *Database) UpdateRow(tableName string, rowID string, updates map[string]string) (*Row, error){
	db.mu.RLock()
	table, exists := db.Tables[tableName]
	db.mu.RUnlock()

	if !exists {
		return nil, errors.New("table does not exist")
	}

	table.mu.RLock()
	row, exists := table.Rows[rowID]
	table.mu.RUnlock()

	if !exists {
		return nil, errors.New("row not found")
	}

	for col, newValue := range updates {
		if _, exists := row.Values[col]; exists{
			row.Values[col] = newValue
		}
	}

	return db.ReadRow(tableName,rowID)
}


func main(){
	db := NewDatabase()
	
	err := db.CreateTable("Users", []string{"Name", "Email"})
	if err != nil {
		fmt.Println(err)
	}

	rowID, err := db.InsertRow("Users", []string{"Jayendra","rock@example.com"}); 
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Row Inserted: ", rowID)

	row, err := db.ReadRow("Users", rowID)
	if err != nil {
		fmt.Println("Error reading row:", err)
	}
	fmt.Println("Row:", row)


	row, err = db.UpdateRow("Users", rowID, map[string]string{"Email": "john.doe@example.com"})
	if err != nil {
		fmt.Println("Error updating row:", err)
	}
	fmt.Println("Updated Row :", row)

}