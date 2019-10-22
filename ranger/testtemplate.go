package ranger

import "strings"

var testTemplate = strings.TrimLeft(`
// +build rangertest

package {{ .Package }}

import (
	"reflect"
	"testing"
	"math/rand"
	"time"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"

	"github.com/stretchr/testify/assert"
)

func genRandom(n int) []byte {
	data := make([]byte, n)
	n2, err := rand.Read(data)
	if err != nil {
		panic(errors.Wrap(err, "rand.Read"))
	}

	if n != n2 {
		panic(errors.Wrap(err, "short read in rand.Read"))
	}

	return data[:n]
}

{{ range $typekey, $type := .Types -}}
func Test{{ $typekey }}MarshalUnmarshalZeroValue(t *testing.T) {
	obj := &{{ $typekey }}{
		{{ range $foo, $map := $type.Fields -}}
		{{ range $innerkey, $innervalue := $map -}}
		{{ $innerkey }}:  {{ zeroValue $innervalue.TypeName $innervalue.StructureType $innervalue.ValueType $innervalue.Require.Length $innervalue.Interface }},
		{{- end -}}
		{{- end -}}
	}

	obj2 := &{{ $typekey }}{
		{{ range $foo, $map := $type.Fields -}}
		{{ range $innerkey, $innervalue := $map -}}
		{{ $innerkey }}: {{ zeroValue $typekey $innervalue.StructureType $innervalue.ValueType $innervalue.Require.Length $innervalue.Interface }},
		{{- end -}}
		{{- end -}}
	}

	data, err := obj.Marshal()
	assert.Nil(t, err, "marshal failed for {{ $typekey }}")
	assert.Equal(t, len(data), obj.Size(), "{{ $typekey }} size check on zero value")
	assert.Nil(t, obj2.Unmarshal(data), "{{ $typekey }} zero value unmarshal test")
	assert.Equal(t, obj, obj2, "{{ $typekey }} unmarshal equality test")
	obj2 = &{{ $typekey }}{
		{{ range $foo, $map := $type.Fields -}}
		{{ range $innerkey, $innervalue := $map -}}
		{{ $innerkey }}: {{ zeroValue $typekey $innervalue.StructureType $innervalue.ValueType $innervalue.Require.Length $innervalue.Interface }},
		{{- end -}}
		{{- end -}}
	}
	l, err := obj2.UnmarshalFrom(data)
	assert.Nil(t, err, "{{ $typekey }} unmarshalfrom failed")
	assert.Equal(t, obj, obj2, "{{ $typekey }} unmarshalfrom equality test")
	assert.Equal(t, len(data), l, "{{ $typekey }} data length check")
	assert.Equal(t, obj.Size(), l, "{{ $typekey }} data size check")
}

func Test{{ $typekey }}MarshalUnmarshalRandomData(t *testing.T) {
	seed := time.Now().Unix()
	fmt.Printf("Seed is %v\n", seed)
	rand.Seed(seed)

	for i := 0; i < 100; i++ {
		obj := &{{ $typekey }}{}
		{{ range $foo, $map := $type.Fields -}}
		{{ range $innerkey, $innervalue := $map -}}
		obj.{{ $innerkey }} = {{ randomField $innervalue }}
		{{ end }}
		{{ end }}

		obj2 := &{{ $typekey }}{}

		data, err := obj.Marshal()
		assert.Nil(t, err, "marshal failed for {{ $typekey }}")
		assert.Equal(t, len(data), obj.Size(), "{{ $typekey }} size check on random values")
		assert.Nil(t, obj2.Unmarshal(data), "{{ $typekey }} random values unmarshal test")
		assert.Equal(t, obj, obj2, "{{ $typekey }} unmarshal equality test")

		obj2 = &{{ $typekey }}{}

		l, err := obj2.UnmarshalFrom(data)
		assert.Nil(t, err, fmt.Sprintf("{{ $typekey }} unmarshalfrom failed: %q", hex.EncodeToString(data)))
		assert.Equal(t, obj, obj2, "{{ $typekey }} unmarshalfrom equality test")
		assert.Equal(t, len(data), l, "{{ $typekey }} data length check")
		assert.Equal(t, obj.Size(), l, "{{ $typekey }} data size check")

		assert.True(t, reflect.DeepEqual(obj, obj2))
	}
}
{{ end -}}
`, "\n")
