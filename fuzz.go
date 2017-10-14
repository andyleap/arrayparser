package arrayparser

// +build gofuzz

func Fuzz(data []byte) int {
	out, err := Parse(string(data))
	if out == nil && err == nil {
		panic("Out nil and err nil!")
	}
	return 0
}
