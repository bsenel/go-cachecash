package ranger

import "strings"

var fuzzTemplate = strings.TrimLeft(`
// +build rangertest

package {{ .Package }}

import (
	"fmt"
	"reflect"
)

func Fuzz(data []byte) int {
	if len(data) < 2 {
		return 0
	}

	switch int(data[0])%{{ size }}{
	{{ $i := 0 }}
	{{ range $typekey, $type := .Types -}}
	case {{ $i }}:
		obj{{ $typekey }} := &{{ $typekey }}{}
		_, err := obj{{ $typekey }}.UnmarshalFrom(data[1:])
		if err != nil {
			return 0
		}
		data{{ $typekey }}, err := obj{{ $typekey }}.Marshal()
		if err != nil {
			panic(err)
		}

		obj{{ $typekey }}2 := &{{ $typekey }}{}
		_, err = obj{{ $typekey }}2.UnmarshalFrom(data{{ $typekey }})
		if err != nil {
			panic(err)
		}

		if !reflect.DeepEqual(obj{{ $typekey }}, obj{{ $typekey }}2) {
			panic(fmt.Sprintf("obj %T not equal", obj{{ $typekey }}))
		}

		{{- $i = (add $i 1) }}
	{{ end -}}
	}

	return 1
}
`, "\n")
