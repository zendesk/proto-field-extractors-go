package fieldextractors

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type FieldExtractor interface {
	Extract(in proto.Message) (interface{}, error)
}

type StructField string

func (sf StructField) Extract(in proto.Message) (interface{}, error) {
	field, err := BasicField(sf).ExtractPath(in)
	if err != nil {
		return nil, nil
	}

	st, ok := field.(*structpb.Struct)
	if !ok {
		return nil, errors.New("given a field that is not a struct")
	}

	return st.AsMap(), nil
}

type RepeatedItem struct {
	RepeatedFieldPath string
	ItemExtractor     FieldExtractor
}

func (f RepeatedItem) Extract(in proto.Message) (interface{}, error) {
	result := []interface{}{}

	repeatedField, err := BasicField(f.RepeatedFieldPath).ExtractPath(in)
	if err != nil {
		return nil, nil
	}
	list, ok := repeatedField.(protoreflect.List)
	if !ok {
		return nil, nil
	}

	for i := 0; i < list.Len(); i++ {
		item := list.Get(i)
		if !item.IsValid() {
			continue
		}

		message, ok := item.Interface().(protoreflect.Message)
		if !ok {
			result = append(result, item.Interface())
			continue
		}

		if !message.IsValid() {
			continue
		}

		extractResult, err := f.ItemExtractor.Extract(message.Interface())
		if err != nil {
			return nil, err
		}
		result = append(result, extractResult)
	}

	return result, nil
}

type ObjectExtractor map[string]FieldExtractor
type Object = ObjectExtractor

func (f ObjectExtractor) Extract(in proto.Message) (interface{}, error) {
	result := map[string]interface{}{}
	for key, extractor := range f {
		extractResult, err := extractor.Extract(in)
		if err != nil {
			return nil, err
		}
		if extractResult == nil {
			continue
		}
		result[key] = extractResult
	}
	if len(result) == 0 {
		return nil, nil
	}
	return result, nil
}

func OneOf(oneOfPath string, itemExtractor FieldExtractor) OneOfExtractor {
	return OneOfExtractor{
		oneOfPath:     oneOfPath,
		itemExtractor: itemExtractor,
	}
}

type OneOfExtractor struct {
	oneOfPath     string
	itemExtractor FieldExtractor
}

func (f OneOfExtractor) Extract(in proto.Message) (interface{}, error) {
	path := strings.Split(f.oneOfPath, ".")

	oneOfFieldName := path[0]

	message := in.ProtoReflect()
	oneOfDesc := message.Descriptor().Oneofs().ByName(protoreflect.Name(oneOfFieldName))
	if oneOfDesc == nil {
		return nil, nil
	}

	oneOfField := message.WhichOneof(oneOfDesc)
	if oneOfField == nil {
		return nil, nil
	}

	desiredOneOfField := path[1]
	if oneOfField.Name() != protoreflect.Name(desiredOneOfField) {
		return nil, nil
	}

	value := message.Get(oneOfField)
	if !value.IsValid() {
		return nil, nil
	}

	return f.itemExtractor.Extract(value.Message().Interface())
}

type BasicField string
type B = BasicField

func (f BasicField) ExtractPath(in proto.Message) (interface{}, error) {
	return f.getPath(in, string(f), false)
}

func (f BasicField) Extract(in proto.Message) (interface{}, error) {
	return f.getPath(in, string(f), true)
}

func (f BasicField) getPath(in proto.Message, path string, returnBasicType bool) (interface{}, error) {
	var value interface{}
	var currentPath proto.Message
	parts := strings.Split(path, ".")

	currentPath = in
	for _, part := range parts {
		value = f.getProperty(currentPath, part)
		if t, ok := value.(proto.Message); ok {
			currentPath = t
		} else {
			break
		}
	}

	if returnBasicType {
		return extractValue(value), nil
	}

	// clients aren't going to want protoreflect interfaces, so unwrap them
	if v, ok := value.(protoreflect.Value); ok {
		return v.Interface(), nil
	}
	if v, ok := value.(proto.Message); ok {
		return v.ProtoReflect().Interface(), nil
	}
	return value, nil
}

func (f BasicField) getProperty(in proto.Message, name string) interface{} {
	message := in.ProtoReflect()
	field := message.Descriptor().Fields().ByName(protoreflect.Name(name))
	if field == nil || !message.Has(field) {
		return nil
	}

	fieldValue := message.Get(field)
	if field.IsList() {
		return fieldValue.List()
	}
	if field.IsMap() {
		return fieldValue.Map()
	}
	if field.Kind() == protoreflect.MessageKind {
		return fieldValue.Message().Interface()
	}
	if field.Kind() == protoreflect.EnumKind {
		enumString := string(field.Enum().Values().ByNumber(fieldValue.Enum()).Name())
		return strings.ToLower(enumString)
	}

	return fieldValue
}

type Value string
type V = Value

func (v Value) Extract(in proto.Message) (interface{}, error) {
	return string(v), nil
}

type ExtractorWithDefault struct {
	DefaultValue interface{}
	Extractor    FieldExtractor
}

func (f ExtractorWithDefault) Extract(in proto.Message) (interface{}, error) {
	extractResult, err := f.Extractor.Extract(in)
	if err != nil {
		return nil, err
	}
	if extractResult == nil || reflect.ValueOf(extractResult).IsZero() {
		return f.DefaultValue, nil
	}
	return extractResult, nil
}

type Template struct {
	format string
	args   []FieldExtractor
}

func Tmpl(format string, args ...FieldExtractor) Template {
	return Template{format: format, args: args}
}

func (f Template) Extract(in proto.Message) (interface{}, error) {
	resultArgs := []interface{}{}

	for _, r := range f.args {
		extractResult, err := r.Extract(in)
		if err != nil {
			return nil, err
		}

		resultArgs = append(resultArgs, extractResult)
	}

	return fmt.Sprintf(f.format, resultArgs...), nil
}

func extractValue(in interface{}) interface{} {
	if in == nil {
		return nil
	}

	if reflect.ValueOf(in) == reflect.Zero(reflect.TypeOf(in)) {
		return nil
	}
	switch in := in.(type) {
	case nil:
		return nil
	case *wrapperspb.Int32Value:
		return in.GetValue()
	case *wrapperspb.Int64Value:
		return in.GetValue()
	case *wrapperspb.StringValue:
		return in.GetValue()
	case *wrapperspb.BoolValue:
		return in.GetValue()
	default:
		return in
	}
}
