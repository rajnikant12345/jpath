package jpath

import (
	"encoding/json"
	"testing"
)

func TestGetJsonAtPath(t *testing.T) {
	j, _ := CompileNewJsonPath([]string{"[*].car.color"})

	in := `[
 { "car": {
    "color": "blue",
    "price": "$20,000"
  }
  },
 { "car": {
    "color": "white",
    "price": "$120,000"
  }
  }
]`

	var m interface{}

	json.Unmarshal([]byte(in), &m)

	out, err := j.GetJsonAtPathValue(m)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}

func TestJsonPath_MapJsonAtPathValue(t *testing.T) {
	j, _ := CompileNewJsonPath([]string{"[1].car.color"})

	in := `[
 { "car": {
    "color": "blue",
    "price": "$20,000"
  }
  },
 { "car": {
    "color": "white",
    "price": "$120,000"
  }
  }
]`
	var m interface{}

	json.Unmarshal([]byte(in), &m)

	out, err := j.MapJsonAtPathValue(m, func(in interface{}) interface{} {
		return "hello"
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)

}
