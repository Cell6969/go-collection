# Sub Test

Golang mendukung untuk pembuatan function unit test didalam function unit test
Sebagai contoh:
```go
func TestSubTest(t *testing.T) {
	t.Run("jonathan", func(t *testing.T) {
		result := HelloWord("Jonathan")
		assert.Equal(t, "Hello Jonathan", result)
	})
	t.Run("aldo", func(t *testing.T) {
		result := HelloWord("Aldo")
		assert.Equal(t, "Hello Aldo", result)
	})
}
```
