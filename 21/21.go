package main

import "fmt"

// Паттерн Адаптер (Adapter) предназначен для преобразования интерфейса одного класса в интерфейс другого. 
// Благодаря реализации данного паттерна мы можем использовать вместе классы с несовместимыми интерфейсами.



type animal interface {
    speedOfRunning()
}

type mammal struct {
	speedOfRun int
}

func (m *mammal) speedOfRunning() {
    fmt.Printf("Животное бегает со скоростью %d\n", m.speedOfRun)
}

type birdAdapter struct {
	thisBird *bird
}

func (w *birdAdapter) speedOfRunning() {
	w.thisBird.speedOfFlying()
}

type bird struct{
	speedOfFly int
}

func (b *bird) speedOfFlying() {
    fmt.Printf("Птичка летает со скоростью %d\n", b.speedOfFly)
}

type client struct {
}

func (c *client) speedOfActions(com animal) {
    com.speedOfRunning()
}

func main() {
    client := &client{}
    mammal := &mammal{speedOfRun:2}
    client.speedOfActions(mammal)
    bird := &bird{speedOfFly:100}
    birdAdapter := &birdAdapter{
        thisBird: bird,
    }
    client.speedOfActions(birdAdapter)
}