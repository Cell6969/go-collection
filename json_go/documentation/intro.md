# Intro 
Golang menyediakan package berupa json yang berfungsi untuk melakuakn proses encoding dan decoding dengan tipe data json. Mengapa json ? karena saat ini data yang sering digunakan untuk melakukan request dan response lebih banyak menggunakan json.

## Encoding
Untuk melakukan encoding data menjadi data json, bisa seperti berikut:
```go
func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("aldo")
	logJson(1)
	logJson(true)
	logJson([]string{"a", "b", "c"})
}
```
jadi dengan marshal bisa langsung menconvert data yang diberikan

## JSON Object
Kasus diatas merupakan bukan kasus json. Hal ini dikarenakan tidak mengikuti kontrak dari json. Untuk penerapan json object menggunakan struct yang ada pada golang.

Contoh implementasi:
```go
type Customer struct {
	FirstName string
	LastName  string
	Age       int
}

func TestJSONObject(t *testing.T) {
	var customer Customer = Customer{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
```
Selain dari struct, bisa juga dari map dikonversikan menjadi json:
```go
func TestJSONObjectFromMap(t *testing.T) {
	var person map[string]interface{} = map[string]interface{}{
		"first_name": "John",
		"last_name":  "Doe",
		"age":        30,
	}

	bytes, _ := json.Marshal(person)
	fmt.Println(string(bytes))
}
```


## Decode json
Selain encode menjadi json, bisa juga decode dari json ke data lain (biasanya string)

Contoh implementasi:
```go
func TestDecodedJSON(t *testing.T) {
	var jsonString string = `{"FirstName":"John","LastName":"Doe","Age":30}`
	var jsonBytes []byte = []byte(jsonString)

	var customer *Customer = &Customer{}
	var err error = json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
}
```