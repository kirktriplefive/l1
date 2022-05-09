package main

import "fmt"


type Human struct {//структура Human
	age int
	name string
	job string
}

func (h *Human) AboutMe() {
	fmt.Printf("My name is %s. I am %d years old. My job is %s.\n", h.name, h.age, h.job)//Метод
}

type Action struct {
	favouriteAction string
	Human//встраиваем тип Human в action
}

var (
    human = Human{
        age: 30,
		name: "Human",
		job: "GO-DEV",
    }

    action = Action{
        favouriteAction: "Playing football",
        Human: human,
    }
)

func main() {
	action.AboutMe()
}