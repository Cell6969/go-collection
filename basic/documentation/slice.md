# Slice

Tipe data slice memiliki 3 data yaitu pointer, length dan capacity
Contoh implementasi membuat slice dari array
```go
names := [...]string{"alfred", "kasano", "daji", "limbo"}

slice1 := names[:4]
fmt.Println(slice1)

slice2 := names[2:]
fmt.Println(slice2)

var sliceInit []string = names[2:]
fmt.Println(sliceInit)

slice3 := names[:]
fmt.Println(slice3)
```
## Function pada Slice
Slice juga memiliki function seperti array. Contoh implementasinya
```go
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
```
Pada code diatas jika dieksekusi maka akan memiliki output
```shell
[senin selasa rabu kamis jumat senin baru selasa baru]
daySlice1 [senin baru selasa baru]
daySlice2 [senin baru selasa baru libur baru]
days [senin selasa rabu kamis jumat senin baru selasa baru]
```
### Make
Selain append terdapat juga make yang berfungsi untuk membuat slice
contoh implementasinya:
```go
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
```

### Copy Slice
```go
fromSlice := days[:]
toSlice := make([]string, len(fromSlice), cap(fromSlice))

copy(toSlice, fromSlice)
fmt.Println(fromSlice)
fmt.Println(toSlice)
```

### Array vs Slice
```go
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
```