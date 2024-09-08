package main

import "math/rand"

type Dice struct {
	diceCnt int
	min     int
	max     int
}

func NewDice(diceCnt int) Dice {
	return Dice{
		diceCnt: diceCnt,
		min:     1,
		max:     6,
	}
}

func (d *Dice) RollDice() int {
	totalSum := 0
	for i := 0; i < d.diceCnt; i++ {
		totalSum += rand.Intn(d.max-d.min+1) + d.min
	}
	return totalSum
}
