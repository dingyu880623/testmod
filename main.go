package main

import (
	"fmt"
	"github.com/rs/xid"
)

func main() {
	fmt.Println("dingyu")
	id := xid.New()
	fmt.Printf("%s \n", id)
}
