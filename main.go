package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

type Base struct {
	Brand           string
	Year            int
	BootVolume      float32
	IsEngineStarted bool
	IsWindowOpend   bool
	BootFill        float32
}

type Auto struct {
	Base      Base
	ClassName string
}

type Truck struct {
	Base       Base
	WheelCount int
}

func auto() {
	auto := models.Auto{
		Base:      models.Base{Brand: "Opel", Year: 2012, BootVolume: 30.5, BootFill: 20.8, IsWindowOpend: false},
		ClassName: "Lift"}
	truck := models.Truck{
		Base:       models.Base{Brand: "Volvo", Year: 2005, BootVolume: 120, BootFill: 80, IsEngineStarted: true},
		WheelCount: 14}
	fmt.Println(auto)
	fmt.Println(truck)
	Auto.Base.IsEngineStarted = true
	truck.Base.IsWindowOpend = true
	fmt.Println(auto, reflect.TypeOf(auto))
	fmt.Println(truck, reflect.TypeOf(truck))
}
func qPush(res int) (err int) {

	if len(Queue)-1 > QueuePos {
		QueuePos++
		Queue[QueuePos] = res
		return 0
	}
	return 1

}

func qPop() (res int, err int) {

	if QueuePos > 0 {
		res = Queue[QueuePos]
		QueuePos--
		return res, 0
	}
	return 0, 1

}

func main() {
	auto()

	for i := 1; i < 22; i++ {

		if err := qPush(i); err > 0 {
			log.Println("Error - переполнение очереди", err)
			break
		}
	}
	fmt.Println(Queue)

	for i := 1; i < 25; i++ {

		res, err := qPop()
		if err > 0 {
			log.Println("Error - окончание очереди", err)
			break
		}

		fmt.Println("Элемент ", i, "=", res)
	}

	phoneBook := make(map[string]int)

	phoneBook["Alex"] = 79035484225
	phoneBook["Anna"] = 79651115737
	phoneBook["Sergey"] = 79047225789
	phoneBook["Lena"] = 79037574453

	fmt.Println(phoneBook)

	serialized, err := json.Marshal(phoneBook)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(serialized))

	err = ioutil.WriteFile("phoneBook", serialized, 0644)
	if err != nil {
		log.Fatal(err)
	}

	delete(phoneBook, "Sergey")
	fmt.Println("Удалили запись Sergey")
	fmt.Println(phoneBook)

	content, err := ioutil.ReadFile("phoneBook")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
	fmt.Println(phoneBook)

	phone := &phoneBook
	if err := json.Unmarshal(content, phone); err != nil {
		log.Fatal(err)
	}
	fmt.Println(phoneBook)
}
