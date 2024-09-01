package main


// SRP Defintion - A Class or module should have only 1 reason to change 

type User1 struct {
	FirstName string
	LastName string
}

func (u *User1) GetFullName() string{
	return u.FirstName +" " + u.LastName
}

// if user1 module changes, than GetFullName() & Save() is also Impacted
func (u *User1) Save() error {
	// Save to database
	return nil
}

// -----------------------------------------------------------------------------------------------

// SRP - Now each Struct has single responsibility 
type User2 struct {
    FirstName string
    LastName  string
}

func (u *User2) GetFullName() string {
    return u.FirstName + " " + u.LastName
}

type UserRepository struct {
    // Database connection or other storage-related fields
}

func (r *UserRepository) Save(u *User2) error {
    // Save user to the database
	return nil
}