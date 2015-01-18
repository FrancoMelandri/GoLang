// In Go an interface is a custom type that specifies a set of one or more method
// signatures. Interfaces are wholly abstract, so it is not possible to instantiate an
// interface. However, it is possible to create a variable whose type is that of an
// interface—and which can then be assigned a value of any concrete type that has
// the methods the interface requires

package main

import (
	"fmt"
)

type Exchanger interface {
	Exchange()
}

type StringPair struct {
	first, second string
}

type IntPair struct {
	first, second int
}

func (s *StringPair) Exchange() {
	s.first, s.second = s.second, s.first
}
func (s *StringPair) String() string {
	return fmt.Sprintf("%q+%q", s.first, s.second)
}

func (i *IntPair) Exchange() {
	i.first, i.second = i.second, i.first
}
func (i *IntPair) String() string {
	return fmt.Sprintf("%q+%q", i.first, i.second)
}

func DoExchange(ex Exchanger) {
	ex.Exchange()
}

func DoExchangeAll(exchangers ...Exchanger) {
	for _, exchanger := range exchangers {
		exchanger.Exchange()
	}
}

func main() {
	sp := StringPair{"first", "second"}
	ip := IntPair{1, 2}
	DoExchange(&sp)
	DoExchange(&ip)
	fmt.Println("SP: ", sp)
	fmt.Println("IP: ", ip)
	DoExchangeAll(&sp, &ip)
	fmt.Println("SP: ", sp)
	fmt.Println("IP: ", ip)
}
