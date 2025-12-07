package main

import "fmt"

type Mover interface {
	Move()
	Speed() float64
}

type Car struct {
	name         string
	currentSpeed float64
}

func (c Car) Move() {
	fmt.Println(c.name, "正在行驶...")
}

func (c Car) Speed() float64 {
	return c.currentSpeed
}

func Statr(m Mover) {
	fmt.Printf("准备启动设备，最大速度限制为 %.2f km/h\n", m.Speed())
	m.Move()
}

func main() {
	myCar := Car{name: "Xiaomi SU7", currentSpeed: 1000.5}
	Statr(myCar)
}
