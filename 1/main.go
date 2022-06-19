package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Height int
	Weight int
	Action
}

type Action struct {
	Running
	Walking
	Swimming
}

type Running struct {
	Speed
}

type Walking struct {
	Speed
}

type Swimming struct {
	Speed
}

type Speed float64

func main() {
	Alice := Human{
		Name:   "Alice",
		Age:    30,
		Height: 170,
		Weight: 60,
		Action: Action{
			Running: Running{
				Speed: 12.5,
			},
			Walking: Walking{
				Speed: 5.5,
			},
			Swimming: Swimming{
				Speed: 7.2,
			},
		},
	}

	BobsWalking := Walking{Speed: 6.7}
	BobsRunning := Running{Speed: 13.5}
	Bob := Human{
		Name:   "Bob",
		Age:    22,
		Height: 176,
		Weight: 78,
		Action: Action{
			Walking: BobsWalking,
			Running: BobsRunning,
		},
	}

	fmt.Printf("%#v\n%#v", Alice, Bob)

}
