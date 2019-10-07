package ranger

import "strings"

var goTemplate = strings.TrimLeft(`
{{ define "unmarshalType" }}
{{- if (native .ValueType) }}
{{- if (isBytes .ValueType) -}}
	{{ declare .Item .TypeName "unmarshal" "iL" "uint64" }}
	{{ declare .Item .TypeName "unmarshal" "ni" "int" }}
	iL, ni = binary.Uvarint(data[n:])
	if ni <= 0 {
		return 0, errors.Wrap(ranger.ErrShortRead, "Obtaining length of {{ .TypeName }}.{{ .FieldName }}")
	}
	n += ni
	{{ declare .Item .TypeName "unmarshal" "byt" "[]byte"	}}
	{{- if (and .Require.Length (native .ValueType)) }}
	if iL != {{ .Require.Length }} {
		return 0, errors.Wrap(ranger.ErrLengthMismatch, "{{ .TypeName }}.{{ .FieldName }}")
	}
	{{- else if (and .Require.MaxLength (native .ValueType)) }}
	if iL > {{ .Require.MaxLength }} {
		{{ if .Item }}
		return 0, errors.Wrapf(ranger.ErrTooMany, "{{ .TypeName }}.{{ .FieldName }} index %d", i)
		{{ else }}
		return 0, errors.Wrap(ranger.ErrTooMany, "{{ .TypeName }}.{{ .FieldName }}")
		{{ end }}
	}
	{{ end }}

	if iL > uint64(len(data[n:])) {
		{{ if .Item }}
		return 0, errors.Wrapf(ranger.ErrShortRead, "{{ .TypeName }}.{{ .FieldName }} index %d", i)
		{{ else }}
		return 0, errors.Wrap(ranger.ErrShortRead, "{{ .TypeName }}.{{ .FieldName }}")
		{{ end }}
	}

	if iL > {{ .MaxByteRange }} {
		return 0, errors.Wrap(ranger.ErrTooMany, "Out of range in {{ .TypeName }}.{{ .FieldName }}")
	}

	byt = make([]byte, iL)
	n += copy(byt, data[n:uint64(n)+iL])
	obj.{{ .FieldName }}{{ if .Item }}[i]{{ end }} = {{ if (eq .ValueType "string") }}string({{ end }}byt{{ if (eq .ValueType "string") }}){{ end }}
{{- else -}}
	{{- if (not .Require.Static) }}
	{{ declare .Item .TypeName "unmarshal" "iL" "uint64" }}
	{{- end }}
	{{ declare .Item .TypeName "unmarshal" "ni" "int" }}

	{{- if .Require.Static }}
	if len(data[n:]) < {{ typeLength .ValueType .FieldName .Require.Static }} {
		return 0, errors.Wrap(ranger.ErrShortRead, "{{ .TypeName }}.{{ .FieldName }}")
	}
	{{- end }}
	{{ unmarshaler .ValueType (printf "obj.%s" .FieldName) .Item .Require.Static }}

	{{- if (and (not .Require.Static) (not (eq .StructureType "array"))) }}
	if iL & {{ truncated .ValueType }} != iL {
		return 0, errors.Wrap(ranger.ErrTooLarge, "{{ .TypeName }}.{{ .FieldName }}")
	}
	{{- end }}
	{{- if (not .Require.Static) }}
	n += ni
	{{- end }}
{{- end -}}
{{- else -}}
{{- if .Interface }}
	if len(data[n:]) < {{ typeLength .Interface.Input .FieldName true }} {
		return 0, errors.Wrap(ranger.ErrShortRead, "{{ .TypeName }}.{{ .FieldName }} ({{ .ValueType }})")
	}
	var intf {{ .Interface.Input }}
	{{ unmarshaler .Interface.Input "intf" .Item true }}

	var v {{ .ValueType }}

	switch intf {
	{{ range $foo, $map := .Interface.Cases -}}
	{{ range $key, $value := $map -}}
	case {{ $key }}:
		v = {{ printf "&%s{}" $value }}
	{{ end -}}
	{{ end -}}
	default:
		return 0, errors.Wrap(ranger.ErrBadInterface, "{{ .TypeName }}.{{ .FieldName }} ({{ .ValueType }})")
	}

	obj.{{ .FieldName }} = v
	{{ declare .Item .TypeName "unmarshal" "iL" "uint64" }}
	{{ declare .Item .TypeName "unmarshal" "ni" "int" }}
	{{ declare .Item .TypeName "unmarshal" "err" "error" }}
	iL, ni = binary.Uvarint(data[n:])
	if ni <= 0 {
		return 0, errors.Wrap(ranger.ErrShortRead, "Obtaining length of {{ .TypeName }}.{{ .FieldName }}")
	}
	n += ni
	if iL > uint64(len(data[n:])) {
		return 0, errors.Wrap(ranger.ErrShortRead, "{{ .TypeName }}.{{ .FieldName }}")
	}
	if iL > {{ .MaxByteRange }} {
		return 0, errors.Wrap(ranger.ErrTooMany, "Out of range in {{ .TypeName }}.{{ .FieldName }}")
	}
	ni, err = obj.{{ .FieldName }}.UnmarshalFrom(data[n:uint64(n)+iL])
	if err != nil {
		return 0, errors.Wrap(err, "{{ .TypeName }}.{{ .FieldName }} ({{ .ValueType }})")
	}

	n += ni
{{- else if (or .InlineStruct (eq .StructureType "scalar")) }}
		{{ declare .Item .TypeName "unmarshal" "iL" "uint64" }}
		{{ declare .Item .TypeName "unmarshal" "ni" "int" }}
		{{ declare .Item .TypeName "unmarshal" "err" "error" }}
		iL, ni = binary.Uvarint(data[n:])
		{{- if (and .Require.Length (native .ValueType))}}
		if iL != {{ .Require.Length }} {
			return 0, errors.Wrap(ranger.ErrLengthMismatch, "{{ .TypeName }}.{{ .FieldName }}")
		}
		{{- else if (and .Require.MaxLength (native .ValueType)) }}
		if iL > {{ .Require.MaxLength }} {
			return 0, errors.Wrap(ranger.ErrTooMany, "{{ .TypeName }}.{{ .FieldName }}")
		}
		{{ end }}
		n += ni
		if iL > uint64(len(data[n:])) {
			return 0, errors.Wrap(ranger.ErrShortRead, "{{ .TypeName }}.{{ .FieldName }}")
		}
		if iL > {{ .MaxByteRange }} {
			return 0, errors.Wrap(ranger.ErrTooMany, "Out of range in {{ .TypeName }}.{{ .FieldName }}")
		}

		ni, err = obj.{{ .FieldName }}.UnmarshalFrom(data[n:uint64(n)+iL])
		if err != nil {
			return 0, errors.Wrap(err, "Unmarshaling into {{ .TypeName }}.{{ .FieldName }}")
		}

		n += ni
{{- else -}}
		{{ declare .Item .TypeName "unmarshal" "iL" "uint64" }}
		{{ declare .Item .TypeName "unmarshal" "ni" "int" }}
		{{ declare .Item .TypeName "unmarshal" "err" "error" }}
		iL, ni = binary.Uvarint(data[n:])
		if ni <= 0 {
			{{ if .Item }}
			return 0, errors.Wrapf(ranger.ErrShortRead, "Obtaining length of {{ .TypeName }}.{{ .FieldName }} index %d", i)
			{{ else }}
			return 0, errors.Wrap(ranger.ErrShortRead, "Obtaining length of {{ .TypeName }}.{{ .FieldName }}")
			{{ end }}
		}
		{{- if (and .Require.Length (native .ValueType)) }}
		if iL != {{ .Require.Length }} {
			{{ if .Item }}
			return 0, errors.Wrapf(ranger.ErrLengthMismatch, "{{ .TypeName }}.{{ .FieldName }} index %d", i)
			{{ else }}
			return 0, errors.Wrap(ranger.ErrLengthMismatch, "{{ .TypeName }}.{{ .FieldName }}")
			{{ end }}
		}
		{{- else if (and .Require.MaxLength (native .ValueType)) }}
		if iL > {{ .Require.MaxLength }} {
			{{ if .Item }}
			return 0, errors.Wrapf(ranger.ErrTooMany, "{{ .TypeName }}.{{ .FieldName }} index %d", i)
			{{ else }}
			return 0, errors.Wrap(ranger.ErrTooMany, "{{ .TypeName }}.{{ .FieldName }}")
			{{ end }}
		}
		{{ end }}
		n += ni
		if iL > uint64(len(data[n:])) {
			{{ if .Item }}
			return 0, errors.Wrapf(ranger.ErrShortRead, "{{ .TypeName }}.{{ .FieldName }} index %d", i)
			{{ else }}
			return 0, errors.Wrap(ranger.ErrShortRead, "{{ .TypeName }}.{{ .FieldName }}")
			{{ end }}
		}
		if iL > {{ .MaxByteRange }} {
			return 0, errors.Wrap(ranger.ErrTooMany, "Out of range in {{ .TypeName }}.{{ .FieldName }}")
		}

		ni, err = obj.{{ .FieldName }}{{ if .Item }}[i]{{ end }}.UnmarshalFrom(data[n:uint64(n)+iL])
		if err != nil {
			{{ if .Item }}
			return 0, errors.Wrapf(ranger.ErrShortRead, "Unmarshaling into {{ .TypeName }}.{{ .FieldName }} index %d", i)
			{{ else }}
			return 0, errors.Wrap(ranger.ErrShortRead, "Unmarshaling into {{ .TypeName }}.{{ .FieldName }}")
			{{ end }}
		}
		n += ni
{{- end -}}
{{- end -}}
{{- end }}

{{- define "marshalType" }}
{{- if (native .ValueType) }}
{{- if .Require.Static }}
	{{- if .Item }}
	if len(data[n:]) < {{ typeLength .ValueType "item" .Require.Static }} {
		return 0, errors.Wrap(ranger.ErrShortWrite, "{{ .TypeName }}.{{ .FieldName }}")
	}
	{{ marshaler .ValueType "item" .Require.Static }}
	n += {{ typeLength .ValueType "item" .Require.Static }}
	{{- else }}
	if len(data[n:]) < {{ typeLength .ValueType (printf "obj.%s" .FieldName) .Require.Static }} {
		return 0, errors.Wrap(ranger.ErrShortWrite, "{{ .TypeName }}.{{ .FieldName }}")
	}
	{{ marshaler .ValueType (printf "obj.%s" .FieldName) .Require.Static }}
	n += {{ typeLength .ValueType (printf "obj.%s" .FieldName) .Require.Static }}
	{{- end }}
{{- else }}
{{- if (isBytes .ValueType) -}}
	if len(data[n:]) < ranger.UvarintSize(uint64(len({{ if .Item }}item{{ else }}obj.{{ .FieldName }}{{ end }}))) + len({{ if .Item }}item{{ else }}obj.{{ .FieldName }}{{ end }}) {
		return 0, errors.Wrap(ranger.ErrShortWrite, "{{ .TypeName }}.{{ .FieldName }}")
	}
	n += binary.PutUvarint(data[n:], uint64(len({{ if .Item }}item{{ else }}obj.{{ .FieldName }}{{ end }})))
{{- end }}
{{- if .Item }}
n += {{ marshaler .ValueType "item" .Require.Static }}
{{- else }}
n += {{ marshaler .ValueType (printf "obj.%s" .FieldName) .Require.Static }}
{{- end }}
{{- end -}}
{{- else -}}
{{- if .Interface }}
	{{ declare .Item .TypeName "marshal" "ni" "int" }}
	{{ declare .Item .TypeName "marshal" "err" "error" }}
	{{- if .Item }}
		{{ marshaler .Interface.Input "item" true }}
		if len(data[n:]) < {{ typeLength .ValueType "item" .Require.Static }} {
			return 0, errors.Wrap(ranger.ErrShortWrite, "{{ .TypeName }}.{{ .FieldName }}")
		}
		n += {{ typeLength .Interface.Input "item" true }}
		ni, err = {{ if .Item }}item{{ else }}obj.{{ .FieldName }}{{ end }}.MarshalTo(data[n:n+{{ if .Item }}item{{ else }}obj.{{ .FieldName }}{{ end }}.Size()])
		if err != nil {
			return 0, errors.Wrap(err, "{{ .TypeName }}.{{ .FieldName }}")
		}
		n += ni
	{{- else }}
	if obj.{{ .FieldName }} != nil {
		{{ marshaler .Interface.Input (printf "obj.%s.%s()" .FieldName .Interface.Output) true }}
		n += {{ typeLength .Interface.Input (printf "obj.%s.%s()" .FieldName .Interface.Output) true }}
		if len(data[n:]) < obj.{{ .FieldName }}.Size() {
			return 0, errors.Wrap(ranger.ErrShortWrite, "{{ .TypeName }}.{{ .FieldName }}")
		}
		n += binary.PutUvarint(data[n:], uint64(obj.{{ .FieldName }}.Size()))
		ni, err = {{ if .Item }}item{{ else }}obj.{{ .FieldName }}{{ end }}.MarshalTo(data[n:n+{{ if .Item }}item{{ else }}obj.{{ .FieldName }}{{ end }}.Size()])
		if err != nil {
			return 0, errors.Wrap(err, "{{ .TypeName }}.{{ .FieldName }}")
		}
		n += ni
	} else {
		return 0, errors.Wrap(ranger.ErrShortWrite, "{{ .TypeName }}.{{ .FieldName }} interface is nil, cannot continue")
	}
	{{- end }}
{{- else }}
	{{ declare .Item .TypeName "marshal" "ni" "int" }}
	{{ declare .Item .TypeName "marshal" "err" "error" }}
	if {{ if .Item }}item{{ else }}obj.{{ .FieldName }}{{ end }} != nil {
		if len(data[n:]) < ranger.UvarintSize(uint64({{ if .Item }}item{{ else }}obj.{{ .FieldName }}{{ end }}.Size()))+{{ if .Item }}item{{ else }}obj.{{ .FieldName }}{{ end }}.Size() {
			return 0, errors.Wrap(ranger.ErrShortWrite, "{{ .TypeName }}.{{ .FieldName }}")
		}
		n += binary.PutUvarint(data[n:], uint64({{ if .Item }}item{{ else }}obj.{{ .FieldName }}{{ end }}.Size()))
		ni, err = {{ if .Item }}item{{ else }}obj.{{ .FieldName }}{{ end }}.MarshalTo(data[n:n+{{ if .Item }}item{{ else }}obj.{{ .FieldName }}{{ end }}.Size()])
		if err != nil {
			return 0, errors.Wrap(err, "{{ .TypeName }}.{{ .FieldName }}")
		}
		n += ni
	} else {
		n += ranger.UvarintSize(0)
	}
{{- end -}}
{{- end -}}
{{- end -}}
{{- $outer := . }}
package {{ .Package }}

import (
	"encoding/binary"

	"github.com/pkg/errors"
	"github.com/cachecashproject/go-cachecash/ranger"
)

{{ range $typekey, $type := .Types -}}
type {{ $typekey }} struct {
{{- range $foo, $map := $type.Fields -}}
{{- range $key, $value := $map }}
	{{ if $value.InlineStruct }}*{{ end }}{{ $key }} {{ if (not $value.InlineStruct) }}{{ if (eq $value.StructureType "array") }}[]{{ end }}{{ end }}{{ if (and (and (not $value.InlineStruct) (not $value.Interface)) (not (native $value.ValueType))) }}*{{ end }}{{ if (not $value.InlineStruct) }}{{ $value.ValueType }}{{ end }}
{{- end }}
{{- end }}
}

func (obj *{{ $typekey }}) Marshal() ([]byte, error) {
	data := make([]byte, obj.Size())
	n, err := obj.MarshalTo(data)
	if err != nil {
		return nil, errors.Wrap(err, {{ printf "%q" $typekey }})
	}

	if n != len(data) {
		return nil, errors.Wrap(ranger.ErrMarshalLength, {{ printf "%q" $typekey }})
	}

	return data, nil
}

func (obj *{{ $typekey }}) MarshalTo(data []byte) (int, error) {
	var n int

{{ range $foo, $map := $type.Fields -}}
{{- range $key, $value := $map }}
{{- if (and $value.Require.Length (not (or .InlineStruct (or (native $value.ValueType) $value.Interface)))) }}
	if len(obj.{{ $key }}) != {{ $value.Require.Length }} {
		return 0, errors.Wrapf(ranger.ErrLengthValidation, "{{ $typekey }}.{{ $key }} - %d/{{ $value.Require.Length }}", len(obj.{{ $key }}))
	}
{{- end }}
{{- if (and $value.Require.MaxLength (not (or .InlineStruct (or (native $value.ValueType) $value.Interface)))) }}
	if len(obj.{{ $key }}) > {{ $value.Require.MaxLength }} {
		return 0, errors.Wrap(ranger.ErrTooMany, "{{ $typekey }}.{{ $key }}")
	}
{{- end }}
{{ if (eq $value.StructureType "scalar") }}
{{ template "marshalType" $value }}
{{ end }}
{{- if (eq $value.StructureType "array") }}
	{{- if $value.Match.LengthOfField }}
	if len(obj.{{ $key }}) != len(obj.{{ $value.Match.LengthOfField }}) {
		return 0, errors.Wrap(ranger.ErrLengthMismatch, "{{ $typekey }}: {{ $key }} and {{ $value.Match.LengthOfField }}")
	}
	{{- end }}

	{{- if (and $value.Require.Length (not (or (native $value.ValueType) $value.Interface))) }}
	if len(obj.{{ $key }}) != {{ $value.Require.Length }} {
		return 0, errors.Wrap(ranger.ErrLengthMismatch, "{{ $typekey }}.{{ $key }}")
	}
	{{- else if (and $value.Require.MaxLength (not (or (native $value.ValueType) $value.Interface))) }}
	if len(obj.{{ $key }}) > {{ $value.Require.MaxLength }} {
		return 0, errors.Wrap(ranger.ErrTooMany, "{{ $typekey }}.{{ $key }}")
	}
	{{ end }}
	n += binary.PutUvarint(data[n:], uint64(len(obj.{{ $key }})))
	for _, item := range obj.{{ $key }} {
	{{- template "marshalType" (itemValue $value)}}
	}
{{- end }}
{{- end -}}
{{- end }}
	return n, nil
}

func (obj *{{ $typekey }}) Size() int {
	var n int
{{ range $foo, $map := $type.Fields -}}
{{- range $key, $value := $map }}
{{- if $value.Interface }}
	if obj.{{ $key }} == nil {
		n += ranger.UvarintSize(0)
	} else {
		n += {{ typeLength $value.Interface.Input "" true }}
	}
{{- end -}}
{{- if (eq $value.StructureType "scalar") }}
{{ if (not (native $value.ValueType)) }}
	if obj.{{ $key }} == nil {
		n += ranger.UvarintSize(0)
	} else {
		n += {{ typeLength $value.ValueType (printf "obj.%s" $key) $value.Require.Static }}
	}
{{ else }}
	n += {{ typeLength $value.ValueType (printf "obj.%s" $key) $value.Require.Static }}
{{ end }}

{{- end -}}
{{- if (eq $value.StructureType "array") }}
	if obj.{{ $key }} == nil {
		n += ranger.UvarintSize(0)
	} else {
		n += ranger.UvarintSize(uint64(len(obj.{{ $key }})))
		for _, item := range obj.{{ $key }} {
			{{- if (native $value.ValueType) }}
			n += {{ typeLength $value.ValueType "item" $value.Require.Static }}
			{{- else }}
			n += ranger.UvarintSize(uint64(item.Size()))
			n += item.Size()
			{{- end }}
		}
	}
{{- end -}}
{{- end -}}
{{- end }}
	return n
}

func (obj *{{ $typekey }}) Unmarshal(data []byte) error {
	_, err := obj.UnmarshalFrom(data)
	return err
}

func (obj *{{ $typekey }}) UnmarshalFrom(data []byte) (int, error) {
{{- if (eq (len $type.Fields) 0) }}
	return 0, nil
{{ else -}}
	if len(data) == 0 {
		return 0, errors.Wrap(ranger.ErrShortRead, "{{ $typekey }}")
	}
	var n int
{{ range $foo, $map := $type.Fields -}}
{{- range $key, $value := $map }}
{{ if (eq $value.StructureType "scalar") }}
{{ if $value.InlineStruct -}}
	{{ declare false .TypeName "unmarshal" "iL" "uint64" }}
	{{ declare false .TypeName "unmarshal" "ni" "int" }}
	{{ declare false .TypeName "unmarshal" "err" "error" }}
	iL, ni = binary.Uvarint(data[n:])
	if ni <= 0 {
		return 0, errors.Wrap(ranger.ErrShortRead, "Obtaining length of {{ .TypeName }}.{{ .FieldName }}")
	}
	n += ni

	if iL > uint64(len(data[n:])) {
		return 0, errors.Wrap(ranger.ErrShortRead, "{{ $typekey }}.{{ $key }}")
	}

	obj.{{ $key }} = &{{ $value.ValueType }}{}

	ni, err = obj.{{ $key }}.UnmarshalFrom(data[n:uint64(n)+iL])
	if err != nil {
		return 0, errors.Wrap(err, "{{ $typekey }}.{{ $key }}")
	}
	n += ni
	if iL != uint64(obj.{{ $key }}.Size()) {
		return 0, errors.Wrap(ranger.ErrUnmarshalLength, "{{ $typekey }}.{{ $key }}")
	}
{{- else -}}
	{{- if (and $value.Require.MaxLength (or (and (not .InlineStruct) (and (not .Interface) (not (native $value.ValueType)))) (isBytes $value.ValueType)))}}
	if len(obj.{{ $key }}) > {{ $value.Require.MaxLength }} {
		return 0, errors.Wrap(ranger.ErrTooMany, "{{ $typekey }}.{{ $key }}")
	}
	{{- end }}
	{{ template "unmarshalType" $value }}
{{ end }}
{{ end }}
{{- if (eq $value.StructureType "array") }}
	{{ declare false .TypeName "unmarshal" "iLen" "uint64" }}
	{{ declare false .TypeName "unmarshal" "ni" "int" }}
	iLen, ni = binary.Uvarint(data[n:])
	if ni <= 0 {
		return 0, errors.Wrap(ranger.ErrShortRead, "Obtaining length of {{ $typekey }}.{{ $key }}")
	}
	n += ni

	{{- if $value.Require.Length }}
	if iLen != {{ $value.Require.Length }} {
		return 0, errors.Wrap(ranger.ErrLengthMismatch, "{{ $typekey }}.{{ $key }}")
	}
	{{- else if $value.Require.MaxLength }}
	if iLen > {{ $value.Require.MaxLength }} {
		return 0, errors.Wrap(ranger.ErrTooMany, "{{ $typekey }}.{{ $key }}")
	}
	{{ end }}

	obj.{{ $key }} = make([]{{ if (not (native $value.ValueType)) }}*{{ end }}{{ $value.ValueType }}, iLen)
	for i := uint64(0); i < iLen; i++ {
		{{- if (not (native $value.ValueType)) }}
		obj.{{ $key }}[i] = &{{ $value.ValueType }}{}
		{{- else if (eq $value.ValueType "[]byte") }}
		obj.{{ $key }}[i] = []byte{}
		{{- end }}
		{{ template "unmarshalType" (itemValue $value) -}}
	}
{{- end }}
{{- if (and $value.Require.Length (or (and (not .InlineStruct) (and (not .Interface) (not (native $value.ValueType)))) (isBytes $value.ValueType)))}}
	if len(obj.{{ $key }}) != {{ $value.Require.Length }} {
		return 0, errors.Wrapf(ranger.ErrLengthValidation, "{{ $typekey }}.{{ $key }} - %d/{{ $value.Require.Length }}", len(obj.{{ $key }}))
	}
{{- end -}}
{{- end }}
{{- end }}
{{- range $foo, $map := $type.Fields -}}
{{- range $key, $value := $map }}
{{ if (and (eq $value.StructureType "array") $value.Match.LengthOfField) }}
	if len(obj.{{ $key }}) != len(obj.{{ $value.Match.LengthOfField }}) {
		return 0, errors.Wrap(ranger.ErrLengthMismatch, "{{ $typekey }}: {{ $key }} and {{ $value.Match.LengthOfField }}")
	}
{{- end }}
{{- end -}}
{{- end -}}

	if len(data[n:]) != 0 {
		return 0, errors.Wrap(ranger.ErrUnmarshalLength, "{{ $typekey }}")
	}
	return n, nil
{{ end }}
}

{{ end }}
`, "\n")
