package main

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

func main() {
	hammer := &hammer{4}
	hammer.Nail()
	phone := phone{7}
	adapter := phoneAdapter{phone}
	adapter.Nail()
}

type tool interface {
	Nail()
}

type hammer struct {
	punchesForHammeringNail int
}

func (h *hammer) Nail() {
	fmt.Printf("We need make %d punches to hammer a nail\n", h.punchesForHammeringNail)
}

type phone struct {
	punchesForBreak int
}

func (p *phone) Throw() {
	fmt.Printf("It will work %d times\n", p.punchesForBreak)
}

type phoneAdapter struct {
	phone
}

func (p phoneAdapter) Nail() {
	p.Throw()
}
