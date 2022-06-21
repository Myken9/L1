package main

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

func main() {
	hammer := &hammer{4}
	hammer.nail()
	phone := phone{7}
	adapter := phoneAdapter{phone}
	adapter.nail()
}

type tool interface {
	nail()
}

type hammer struct {
	punchesForHammeringNail int
}

func (h *hammer) nail() {
	fmt.Printf("We need make %d punches to hammer a nail\n", h.punchesForHammeringNail)
}

type phone struct {
	punchesForBreak int
}

func (p *phone) throw() {
	fmt.Printf("It will work %d times\n", p.punchesForBreak)
}

type phoneAdapter struct {
	phone
}

func (p phoneAdapter) nail() {
	p.throw()
}
