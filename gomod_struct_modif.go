package gomod_struct_modif

import "fmt"

type TestStruct struct {
	name   string
	wealth int64
}

func (t *TestStruct) SetName(name string) {
	t.name = name
}
func (t *TestStruct) SetWealth(number int64) {
	t.wealth = number
}

func (t *TestStruct) ShowData() {
	fmt.Println(t.name)
	fmt.Println(t.wealth)
}
