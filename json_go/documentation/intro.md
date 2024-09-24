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

## JSON Array
Terkadang tipe data di json merupakan array baik array primitive atau array json. 

Contoh implementasi untuk json array primitive
```go
type Customer struct {
    FirstName string
    LastName  string
    Age       int
    Hobbies   []string
}

func TestJSONArray(t *testing.T) {
    var customer *Customer = &Customer{
        FirstName: "Aldo",
        LastName:  "Jonathan",
        Hobbies:   []string{"Gaming", "Coding"},
}

bytes, _ := json.Marshal(customer)
fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
    var jsonString string = `{"FirstName":"Aldo","LastName":"Jonathan","Age":0,"Hobbies":["Gaming","Coding"]}`
    var jsonBytes []byte = []byte(jsonString)

    var customer *Customer = &Customer{}
    var err error = json.Unmarshal(jsonBytes, customer)

    if err != nil {
        panic(err)
    }

    fmt.Println(customer)
}
```

Kemudian untuk case dimana datanya lebih kompleks:
```go
func TestJSONArrayComplexEncode(t *testing.T) {
	var customer *Customer = &Customer{
		FirstName: "Aldo",
		LastName:  "Jonathan",
		Addressess: []Address{
			{
				Street:     "bekasi",
				Country:    "indonesia",
				PostalCode: "123",
			},
			{
				Street:     "bekasi1",
				Country:    "indonesia1",
				PostalCode: "1231",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	var jsonString string = `{"FirstName":"Aldo","LastName":"Jonathan","Age":0,"Hobbies":null,"Addressess":[{"Street":"bekasi","Country":"indonesia","PostalCode":"123"},{"Street":"bekasi1","Country":"indonesia1","PostalCode":"1231"}]}`
	var jsonBytes []byte = []byte(jsonString)

	var customer *Customer = &Customer{}
	var err error = json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Addressess[0].Street)
	fmt.Println(customer.Addressess[0].Country)
	fmt.Println(customer.Addressess[0].PostalCode)
}
```

Bisa juga langsung mendecode array json secara langsung:
```go
func TestDecodeOnlyJSONArray(t *testing.T) {
	var jsonString string = `[{"Street": "Bekasi", "Country" : "Indonesia", "PostalCode": "123"},{"Street": "Bekasi", "Country" : "Indonesia", "PostalCode": "123"}]`
	var jsonBytes []byte = []byte(jsonString)

	var addresses *[]Address = &[]Address{}

	var err error = json.Unmarshal(jsonBytes, addresses)

	if err != nil {
	}

	fmt.Println(addresses)
}
```

## JSON Tag
By default ketika mengkonversi data struct ke json, key nya akan sama dikarenakan case sensitive. Namun ada beberapa kondisi dimana data json yang diterima menggunakan huruf kecil semua. Untuk itu di package json terdapat tag untuk aliasing.

Contoh:
```go
type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURl string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	var product *Product = &Product{
		Id:       "P0001",
		Name:     "Handphone",
		ImageURl: "http://image.com",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	var jsonString string = `{"id":"P0001","name":"Handphone","image_url":"http://image.com"}`
	var jsonBytes []byte = []byte(jsonString)

	var product *Product = &Product{}

	json.Unmarshal(jsonBytes, product)

	fmt.Println(product)
}
```

## Map
Penggunaan struct untuk json terkadang cukup menyulitkan ketika data json nya dynamic. Jikalau data json nya dynamic maka bisa menggunakan map. 

Sebagai contoh:
```go
func TestMapDecode(t *testing.T) {
	var jsonString string = `{"id":"P0001","name":"Handphone","image_url":"http://image.com"}`
	var jsonBytes []byte = []byte(jsonString)

	var result map[string]interface{}

	json.Unmarshal(jsonBytes, &result)
	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["image_url"])
}
```

## Streaming Decoder
Jikalau biasanya data json yang diterima akan disimpan dalam variable, namun hal ini tak perlu dilakukan karena pada package json sudah tersedia fitur stream decoder.

Contoh, buat file json:
```json
{
  "FirstName": "Aldo",
  "LastName": "kevin"
}
```

Kemudian pada code golang:
```go
func TestDecoderStream(t *testing.T) {
	reader, _ := os.Open("Customer.json")
	var decoder *json.Decoder = json.NewDecoder(reader)

	var customer *Customer = &Customer{}
	decoder.Decode(customer)
	fmt.Println(customer)
}
```

## Streaming Encoder
Kebalikannya , bisa juga stream encode tanpa harus pindah ke variable. 

Contoh implementasi:
```go

```