package main

import (
	"basic/database"
	_ "basic/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
