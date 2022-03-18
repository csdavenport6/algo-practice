package helpers

import (
	"fmt"
	"math/rand"
	"os"
)

func GenerateInput() {
	fp := "input.txt"
	f, err := os.Create(fp)
	Check(err)

	defer f.Close()

	for i := 0; i < 100000; i++ {
		num := rand.Intn(1000)
		_, err := f.WriteString(fmt.Sprintf("%d\n", num))
		Check(err)
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
