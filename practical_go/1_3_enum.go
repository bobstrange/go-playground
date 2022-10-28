package main

import "fmt"

type CarType int

// iota を使って enum 的なものを定義する
const (
	Sedan CarType = iota + 1
	Hatchback
	MPV
	SUV
	Crossover
	Coupe
	Convertible
)

// フラグ的なものも定義できる

type CarOption uint64

const (
	GPS CarOption = 1 << iota
	AWD
	SunRoof
	HeatedSeat
	DriverAssist
)

func main() {
	var t CarType = SUV
	fmt.Println("CarType is: ", t) // "CarType is: 4"

	var o CarOption = SunRoof | HeatedSeat
	// & 論理積
	if o&GPS != 0 {
		fmt.Println("GPS")
	}
	if o&SunRoof != 0 {
		fmt.Println("SunRoof")
	}
	if o&HeatedSeat != 0 {
		fmt.Println("HeatedSeat")
	}
}
