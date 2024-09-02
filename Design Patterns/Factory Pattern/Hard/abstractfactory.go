package main

import "fmt"

// ISportsFactory is the abstract factory interface
// It declares the methods for creating abstract products (Shoe and Shirt)
type ISportsFactory interface {
    makeShoe() IShoe
    makeShirt() IShirt
}

// GetSportsFactory returns a concrete factory based on the brand name provided
// If an unrecognized brand is passed, it returns an error
func GetSportsFactory(brand string) (ISportsFactory, error) {
    if brand == "adidas" {
        return &Adidas{}, nil
    }

    if brand == "nike" {
        return &Nike{}, nil
    }

    return nil, fmt.Errorf("Wrong brand type passed")
}

// Adidas is a concrete factory that implements the ISportsFactory interface
// It provides implementations to create Adidas branded shoes and shirts
type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
    return &AdidasShoe{
        Shoe: Shoe{
            logo: "adidas",
            size: 14,
        },
    }
}

func (a *Adidas) makeShirt() IShirt {
    return &AdidasShirt{
        Shirt: Shirt{
            logo: "adidas",
            size: 14,
        },
    }
}

// Nike is another concrete factory that implements the ISportsFactory interface
// It provides implementations to create Nike branded shoes and shirts
type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
    return &NikeShoe{
        Shoe: Shoe{
            logo: "nike",
            size: 14,
        },
    }
}

func (n *Nike) makeShirt() IShirt {
    return &NikeShirt{
        Shirt: Shirt{
            logo: "nike",
            size: 14,
        },
    }
}

// IShoe is the abstract product interface for shoes
// It declares methods to set and get shoe properties
type IShoe interface {
    setLogo(logo string)
    setSize(size int)
    getLogo() string
    getSize() int
}

// Shoe is a concrete implementation of the IShoe interface
// It contains fields for shoe properties such as logo and size
type Shoe struct {
    logo string
    size int
}

func (s *Shoe) setLogo(logo string) {
    s.logo = logo
}

func (s *Shoe) getLogo() string {
    return s.logo
}

func (s *Shoe) setSize(size int) {
    s.size = size
}

func (s *Shoe) getSize() int {
    return s.size
}

// AdidasShoe is a concrete product that embeds the Shoe struct
// It represents a specific brand (Adidas) of shoe
type AdidasShoe struct {
    Shoe
}

// NikeShoe is another concrete product that embeds the Shoe struct
// It represents a specific brand (Nike) of shoe
type NikeShoe struct {
    Shoe
}

// IShirt is the abstract product interface for shirts
// It declares methods to set and get shirt properties
type IShirt interface {
    setLogo(logo string)
    setSize(size int)
    getLogo() string
    getSize() int
}

// Shirt is a concrete implementation of the IShirt interface
// It contains fields for shirt properties such as logo and size
type Shirt struct {
    logo string
    size int
}

func (s *Shirt) setLogo(logo string) {
    s.logo = logo
}

func (s *Shirt) getLogo() string {
    return s.logo
}

func (s *Shirt) setSize(size int) {
    s.size = size
}

func (s *Shirt) getSize() int {
    return s.size
}

// AdidasShirt is a concrete product that embeds the Shirt struct
// It represents a specific brand (Adidas) of shirt
type AdidasShirt struct {
    Shirt
}

// NikeShirt is another concrete product that embeds the Shirt struct
// It represents a specific brand (Nike) of shirt
type NikeShirt struct {
    Shirt
}

func main() {
    // Create a factory for Adidas products
    adidasFactory, _ := GetSportsFactory("adidas")
    // Create a factory for Nike products
    nikeFactory, _ := GetSportsFactory("nike")

    // Use the Nike factory to create Nike branded products
    nikeShoe := nikeFactory.makeShoe()
    nikeShirt := nikeFactory.makeShirt()

    // Use the Adidas factory to create Adidas branded products
    adidasShoe := adidasFactory.makeShoe()
    adidasShirt := adidasFactory.makeShirt()

    // Print details of the Nike products
    printShoeDetails(nikeShoe)
    printShirtDetails(nikeShirt)

    // Print details of the Adidas products
    printShoeDetails(adidasShoe)
    printShirtDetails(adidasShirt)
}

// Helper function to print details of a shoe
func printShoeDetails(s IShoe) {
    fmt.Printf("Shoe Logo: %s\n", s.getLogo())
    fmt.Printf("Shoe Size: %d\n", s.getSize())
}

// Helper function to print details of a shirt
func printShirtDetails(s IShirt) {
    fmt.Printf("Shirt Logo: %s\n", s.getLogo())
    fmt.Printf("Shirt Size: %d\n", s.getSize())
}
