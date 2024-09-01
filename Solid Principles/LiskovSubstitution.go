package main

// if Class B is subtype of Class A, then we should be able to replace 
// object of A with B without breaking the behavior of the Program 


type Bird1 interface{
	fly()
}

type Pigeon1 struct{}

func (p *Pigeon1) fly() string {
	return "pigeon  can fly"
}


type Penguin1 struct{}

// breaks - liskov because pengiun cant fly, narrow down Parent-Bird interface
func (p *Penguin1) fly() string{
	return "pengiun can't fly"
}

// ---------------------------------------------------------------------------------------------

// Parent Struct Bird - Child Flying Bird
type Bird interface {
    MakeSound() string
}

type FlyingBird interface {
    Bird
    Fly() string
}

type Pigeon struct{}

func (p *Pigeon) MakeSound() string {
    return "Coo"
}

func (p *Pigeon) Fly() string {
    return "Pigeon is flying."
}

type Penguin struct{}

func (p *Penguin) MakeSound() string {
    return "Squawk"
}