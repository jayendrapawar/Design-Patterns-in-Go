package main

import (
	"errors"
	"fmt"
	"sync"
)

type KeyValueStore interface {
	AddData(int)
	DeleteData(int) error
	GetData(int) bool
}

type Set struct {
	Key map[int]struct{}
}

func GetNewSet() *Set {
	return &Set{
		Key: make(map[int]struct{}),
	}
}

func (s *Set) AddData(val int) {
	// store in map with empty struct as value
	s.Key[val] = struct{}{}
	fmt.Println("Value is stored:", val)
}

func (s *Set) DeleteData(val int) error {
	// check if the key exists
	_, exists := s.Key[val]
	if !exists {
		return errors.New("value not exists, can't be deleted")
	}
	// delete the key
	delete(s.Key, val)
	fmt.Println("Value is deleted:", val)
	return nil
}

func (s *Set) GetData(val int) bool {
	_, exists := s.Key[val]
	return exists
}

func main() {
	kvs := GetNewSet()
	var wg sync.WaitGroup

	// add and delete elements concurrently
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			kvs.AddData(i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			err := kvs.DeleteData(i)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	// wait for all goroutines to finish
	wg.Wait()

	// check if a specific value exists
	exists := kvs.GetData(1)
	fmt.Println("Does exist in set?", exists)
}
