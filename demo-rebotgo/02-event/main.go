package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func addEvent() {
	fmt.Println("-------------Please press crtl + shift + q ----------")
	ok := robotgo.AddEvents("q", "ctrl", "shift")
	if ok {
		fmt.Println("add events...")
	}

	fmt.Println("---Please press w ----")
	ok = robotgo.AddEvents("w")
	if ok {
		fmt.Println("add events")
	}

	// start hook
	s := robotgo.Start()
	// end hook
	defer robotgo.End()

	for ev := range s {
		fmt.Println("hook", ev)
	}
}

func addMouse() {
	fmt.Println("--- Please press left mouse button ---")
	ok := robotgo.AddMouse("left")
	if ok {
		fmt.Println("add mouse...")
	}

	fmt.Println("--- Please press left mouse button and move mouse to 100, 100 ---")
	ok = robotgo.AddMouse("left", 100, 100)
	if ok {
		fmt.Println("add mouse and move to 100,100 ...")
	}
}

func add() {
	fmt.Println("--- Please press v ---")
	eve := robotgo.AddEvent("v")

	if eve {
		fmt.Println("--- You press v ---", "v")
	}

	fmt.Println("--- Please press k---")
	keve := robotgo.AddEvent("k")
	if keve {
		fmt.Println("--- You press k---", "k")
	}

	fmt.Println("--- Please press f1 ---")
	feve := robotgo.AddEvent("f1")
	if feve {
		fmt.Println("You press....", "f1")
	}
}

func event() {
	/////////////////////////////////////////////////////////
	//  Global event listener
	/////////////////////////////////////////////////////////

	add()

	fmt.Println("--- Please press left mouse button---")
	mleft := robotgo.AddEvent("mleft")
	if mleft {
		fmt.Println("--- You press left mouse button---", "mleft")
	}

	mright := robotgo.AddEvent("mright")
	if mright {
		fmt.Println("--- You press right mouse button---", "mright")
	}

	// stop addEvent
	robotgo.StopEvent()
}

func main() {
	fmt.Println("test begin...")

	addEvent()

	addMouse()

	event()
}
