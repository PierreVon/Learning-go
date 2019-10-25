package main

import "fmt"

type Meal interface {
	whatKindOfFood() string
	whenStart() string
}

type Breakfast struct {
	food string
	time string
}

func (b Breakfast) whatKindOfFood() string {
	return b.food
}

func (b Breakfast) whenStart() string {
	return b.time
}

type Lunch struct {
	food string
	time string
}

func (l Lunch) whatKindOfFood() string {
	return l.food
}

func (l Lunch) whenStart() string {
	return l.time
}

func main() {
	morning := Breakfast{"milk&bread","7am"}
	noon := Lunch{"rice&beef", "12am"}
	var meals = [2]Meal{morning, noon}
	//meals := make([]Meal, 2)
	//meals[0] = morning
	//meals[1] = noon
	for _, meal := range meals{
		fmt.Println(meal.whatKindOfFood() + " " + meal.whenStart())
	}
}