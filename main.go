package main

import (
	"fmt"
	cbor "github.com/brianolson/cbor_go"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
type DemoStruct struct {
	FieldNeedsJsonVsCborName int
	Map map[string]int
}
func main() {
	d := &DemoStruct{
		FieldNeedsJsonVsCborName: 10,
		Map: map[string]int{},
	}
	d.Map["name"] = 2342

	bytes, e := cbor.Dumps(d)
	check(e)

	var ds DemoStruct
	err := cbor.Loads(bytes, &ds)
	check(err)

	fmt.Println(ds)

	fmt.Println(string(bytes))
}

