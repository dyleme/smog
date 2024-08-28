package gen

import (
	"fmt"
	"math/rand/v2"

	"github.com/dyleme/smog/pkg/utils"
	"github.com/go-faker/faker/v4"
)

type Schema struct {
	Type       string            `json:"type"`
	Properties map[string]Schema `json:"properties"`
	Items      *Schema           `json:"items"`
}

func genBySchema(s Schema) any {
	switch s.Type {
	case "string":
		var str string
		err := faker.FakeData(&str)
		utils.NoErr(err)
		return str
	case "int", "number":
		var i int
		err := faker.FakeData(&i)
		utils.NoErr(err)
		return i
	case "bool":
		var b bool
		err := faker.FakeData(&b)
		utils.NoErr(err)
		return b
	case "float":
		var f float64
		err := faker.FakeData(&f)
		utils.NoErr(err)
		return f
	case "email":
		return faker.Email()
	case "object":
		m := make(map[string]any, len(s.Properties))
		for n, p := range s.Properties {
			m[n] = genBySchema(p)
		}
		return m
	case "array":
		length := rand.IntN(5) + 1
		arr := make([]any, 0, length)
		for range length {
			arr = append(arr, genBySchema(*s.Items))
		}
		return arr
	default:
		panic(fmt.Sprintf("unkown type: %q", s.Type))
	}
}
