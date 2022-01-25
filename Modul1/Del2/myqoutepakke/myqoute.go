package myquote

import (
	"fmt"
	"rsc.io/quote"
)

func ShowGlass() {
	fmt.Println(quote.Glass())
}

func ShowGo() {
	fmt.Println(quote.Go())
}

func ShowOpt() {
	fmt.Println(quote.Opt())
}

func ShowHello() {
	fmt.Println(quote.Hello())
}