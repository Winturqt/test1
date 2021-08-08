package main

import "fmt"

var knowledge = k
var k string = "question"
var cry = c
var c string = "wah"

func (p person) speak() {
	fmt.Println(p.name, "says", p.sound)
}

func (d dog) bark() {
	fmt.Println(d.name, "says", d.sound)
}

func (c cat) meow() {
	fmt.Println(c.name, "says", c.sound)
}

type person struct {
	name  string
	age   int
	sound string
}

type dog struct {
	name  string
	age   int
	sound string
}

type cat struct {
	name  string
	age   int
	sound string
}

type chad struct {
	person
	swag bool
}

type teacher struct {
	person
	knowledge string
}

type pug struct {
	dog
	companion bool
}

type lol_player struct {
	dog
	cry          string
	likes_league bool
}

type tabby struct {
	cat
	claws bool
}

type black_cat struct {
	cat
	badluck bool
}

func (pa chad) speak() {
	fmt.Println(pa.name, "says even more", pa.sound)
}

func (da lol_player) bark() {
	fmt.Println(da.name, "says even more", da.sound)
}

func (ta tabby) meow() {
	fmt.Println(ta.name, "says even more", ta.sound)
}

type human interface {
	speak()
}

type animal interface {
	sound()
	bark()
	meow()
}

func foo(h human) {
	h.speak()
}

func foo2(an lol_player) {
	if an.likes_league == true {
		an.sound = "kill me"
	}
	an.bark()
}

func foo3(dg animal) {
	dg.bark()
}

func foo4(ca animal) {
	ca.meow()
}

func main() {
	ch1 := chad{
		person: person{
			name:  "Chadwick",
			age:   35,
			sound: "yo",
		},
		swag: true,
	}
	fmt.Println(ch1)

	tc1 := teacher{
		person: person{
			name:  "Brandon",
			age:   25,
			sound: "what is an interface?",
		},
		knowledge: k,
	}
	fmt.Println(tc1)

	dg1 := pug{
		dog: dog{
			name:  "Rem",
			age:   4,
			sound: "woof",
		},
		companion: true,
	}
	fmt.Println(dg1)

	an1 := lol_player{
		dog: dog{
			name:  "Degenerate",
			age:   40,
			sound: "my teammates are holding me back",
		},
		cry:          cry,
		likes_league: true,
	}
	fmt.Println(an1)

	tb1 := tabby{
		cat: cat{
			name:  "Whiskers",
			age:   3,
			sound: "Meow",
		},
		claws: true,
	}
	fmt.Println(tb1)

	bc1 := black_cat{
		cat: cat{
			name:  "Salem",
			age:   80,
			sound: "hisss",
		},
		badluck: true,
	}
	fmt.Println(bc1)
	foo(ch1)
	foo2(an1)
}
