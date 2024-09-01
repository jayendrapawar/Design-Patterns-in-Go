package main

// Interfaces should be such, that client should not implement
// unnecessary functions they do not need

type RestaurantEmployee1 interface {
    washDishes() string
    serveCustomers() string
    cookFood() string
}

type Waiter1 struct {
    food string
}

func (w *Waiter1) washDishes() string {
	// Not my Job
    return ""
}

func (w *Waiter1) serveCustomers() string {
	// yes my job, here is my implemenation 
	return w.food
}

func (w *Waiter1) cookFood() string{
	// not my Job
	return ""
}

// --------------------------------------------------------------------


type Waiter interface {
	serveCustomers() string
	takeOrder() string
}

type Chef interface {
	cookFood()
}

type frontdesk struct {
	food string
}

func (f *frontdesk) serveCustomers() string {
	// this is my job
	return f.food
}

func (f *frontdesk) takeOrder() string {
	// this is my job
	return f.food
}

type backdesk struct{
	food string
}
 
func (f *backdesk) cookFood() string {
	// this is my job
	return f.food
}
