package main

import "fmt"

func main() {
	// Membuat slice dari array
	names := [...]string{"alfred", "kasano", "daji", "limbo"}

	slice1 := names[:4]
	fmt.Println(slice1)

	slice2 := names[2:]
	fmt.Println(slice2)

	var sliceInit []string = names[2:]
	fmt.Println(sliceInit)

	slice3 := names[:]
	fmt.Println(slice3)

	//	Function Slice
	// create slice from array
	var days = [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	var daysSlice1 []string = days[5:]
	fmt.Println("daySlice", daysSlice1)

	daysSlice1[0] = "senin baru"
	daysSlice1[1] = "selasa baru"
	fmt.Println("daySlice", daysSlice1)
	fmt.Println(days)

	daysSlice2 := append(daysSlice1, "libur baru")
	fmt.Println("daySlice1", daysSlice1)
	fmt.Println("daySlice2", daysSlice2)
	fmt.Println("days", days)

	// make
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "var"
	newSlice[1] = "var"
	fmt.Println(newSlice)
	fmt.Println("Length: ", len(newSlice), "Capacity: ", cap(newSlice))

	newSlice2 := append(newSlice, "var") // harus append karena maksimum len
	fmt.Println(newSlice2)
	fmt.Println("Length: ", len(newSlice2), "Capacity: ", cap(newSlice2))

	newSlice2[0] = "val"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// CopySlice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// Perbedaan array dan slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
