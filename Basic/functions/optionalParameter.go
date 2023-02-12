package main

import "fmt"




type Option func(p Persosn) Persosn

type Persosn struct {
	name  string
	phone string
	email string

	// optional fields
	twitterHandle string
}

func NewPerson(name, phone, email string, options ...Option) Persosn {
	return Persosn{}
}

func AddTwitterHandle(twitterHandle string) Option {
	return func(p Persosn) Persosn {
		p.twitterHandle = twitterHandle
		return p
	}
}

func main() {
	person_1 := NewPerson("Surya", "7068528089", "suryaprakashgond@gmail.com", AddTwitterHandle("surya@1"))
	fmt.Println(person_1)
	person_2 := NewPerson("Sai", "7068528088", "sai@gmail.com", AddTwitterHandle("sai@1"))
	fmt.Println(person_2)
	person_3 := NewPerson("Sai", "7068528088", "sai@gmail.com")
	fmt.Println(person_3)
}
