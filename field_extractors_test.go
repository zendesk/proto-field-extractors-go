package fieldextractors

import (
	"testing"

	protoV2 "github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	example "github.com/zendesk/field-extractors-go/protobuf"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestValue(t *testing.T) {
	t.Run("uses the given string value", func(t *testing.T) {
		extractor := Value("static-id")

		event := &example.Order{
			Id: wrapperspb.Int64(int64(123)),
		}

		// act
		result, _ := extractor.Extract(protoV2.MessageV2(event))

		// assert
		assert.Equal(t, "static-id", result)
	})

	t.Run("can be used with v alias", func(t *testing.T) {
		extractor := V("static-id")

		event := &example.Order{
			Id: wrapperspb.Int64(int64(123)),
		}

		// act
		result, _ := extractor.Extract(protoV2.MessageV2(event))

		// assert
		assert.Equal(t, "static-id", result)
	})
}

func TestNested(t *testing.T) {
	t.Run("extracts an object", func(t *testing.T) {
		extractor := ObjectExtractor{
			"id":    B("id"),
			"token": B("token"),
		}

		event := &example.Order{
			Id:    wrapperspb.Int64(int64(123)),
			Token: wrapperspb.String("my token"),
		}

		expectedResult := map[string]interface{}{
			"id":    int64(123),
			"token": "my token",
		}

		// act
		result, _ := extractor.Extract(protoV2.MessageV2(event))

		// assert
		assert.Equal(t, expectedResult, result)
	})

	t.Run("does not include keys of nil results", func(t *testing.T) {
		extractor := ObjectExtractor{
			"id":    B("id"),
			"token": B("token"),
		}

		event := &example.Order{
			Id:    wrapperspb.Int64(int64(123)),
			Token: nil,
		}

		expectedResult := map[string]interface{}{
			"id": int64(123),
		}

		// act
		result, _ := extractor.Extract(protoV2.MessageV2(event))

		// assert
		assert.Equal(t, expectedResult, result)
	})

	t.Run("returns nil if the extracted result would be empty", func(t *testing.T) {
		extractor := ObjectExtractor{
			"id": B("id"),
		}

		event := &example.Order{
			Id:    nil,
			Token: nil,
		}

		// act
		result, _ := extractor.Extract(protoV2.MessageV2(event))

		// assert
		assert.Nil(t, result)
	})

	t.Run("extracts an arbitrarily nested object", func(t *testing.T) {
		extractor := ObjectExtractor{
			"id": B("id"),
			"nested": Object{
				"nested_token": B("token"),
				"id":           B("id"),
			},
		}

		event := &example.Order{
			Id:    wrapperspb.Int64(int64(123)),
			Token: wrapperspb.String("my token"),
		}

		expectedResult := map[string]interface{}{
			"id": int64(123),
			"nested": map[string]interface{}{
				"nested_token": "my token",
				"id":           int64(123),
			},
		}

		// act
		result, _ := extractor.Extract(protoV2.MessageV2(event))

		// assert
		assert.Equal(t, expectedResult, result)
	})
}

func TestStructField(t *testing.T) {
	t.Run("extracts a struct into a map", func(t *testing.T) {
		extractor := StructField("properties")

		s, _ := structpb.NewStruct(map[string]interface{}{
			"Phone":     "(415) 555-1212",
			"Id":        "001D000000KneakIAB",
			"Name":      "Blackbeard",
			"AccountId": 1234,
			"Tags":      []interface{}{"foo", "bar"},
		})
		event := &example.Order{
			Properties: s,
		}
		expectedResult := map[string]interface{}{
			"Phone":     "(415) 555-1212",
			"Id":        "001D000000KneakIAB",
			"Name":      "Blackbeard",
			"AccountId": float64(1234),
			"Tags":      []interface{}{"foo", "bar"},
		}

		// act
		result, _ := extractor.Extract(protoV2.MessageV2(event))

		// assert
		assert.Equal(t, expectedResult, result)
	})

	t.Run("returns error when used on something other than a struct field", func(t *testing.T) {
		extractor := StructField("id")

		event := &example.Order{
			Id: wrapperspb.Int64(int64(123)),
		}

		// act
		result, err := extractor.Extract(protoV2.MessageV2(event))

		// assert
		assert.Nil(t, result)
		assert.Equal(t, "given a field that is not a struct", err.Error())
	})

	t.Run("returns error when struct not found", func(t *testing.T) {
		extractor := StructField("yolo")

		event := &example.Order{
			Id: wrapperspb.Int64(int64(123)),
		}

		// act
		result, err := extractor.Extract(protoV2.MessageV2(event))

		// assert
		assert.Nil(t, result)
		assert.Equal(t, "given a field that is not a struct", err.Error())
	})
}

func TestRepeatedItem(t *testing.T) {
	t.Run("extracts a valid repeated item into a slice", func(t *testing.T) {
		extractor := RepeatedItem{
			RepeatedFieldPath: "lineitems",
			ItemExtractor: ObjectExtractor{
				"id": B("id"),
			}}

		event := &example.Order{
			Id: wrapperspb.Int64(int64(123)),
			Lineitems: []*example.LineItem{
				{Id: wrapperspb.Int64(int64(1))},
				{Id: wrapperspb.Int64(int64(2))},
			},
		}
		expectedResult := []interface{}{
			map[string]interface{}{
				"id": int64(1),
			},
			map[string]interface{}{
				"id": int64(2),
			},
		}

		// act
		result, _ := extractor.Extract(protoV2.MessageV2(event))

		// assert
		assert.Equal(t, expectedResult, result)
	})

	t.Run("returns nil when used on a non-repeating field", func(t *testing.T) {
		extractor := RepeatedItem{
			RepeatedFieldPath: "id",
			ItemExtractor: ObjectExtractor{
				"id": B("id"),
			}}

		event := &example.Order{
			Id: wrapperspb.Int64(int64(123)),
		}

		// act
		result, _ := extractor.Extract(protoV2.MessageV2(event))

		// assert
		assert.Nil(t, result)
	})

	t.Run("skips nil repeated items", func(t *testing.T) {
		extractor := RepeatedItem{
			RepeatedFieldPath: "lineitems",
			ItemExtractor: ObjectExtractor{
				"id": B("id"),
			}}

		event := &example.Order{
			Id: wrapperspb.Int64(int64(123)),
			Lineitems: []*example.LineItem{
				{Id: wrapperspb.Int64(int64(1))},
				nil,
			},
		}
		expectedResult := []interface{}{
			map[string]interface{}{
				"id": int64(1),
			},
		}

		// act
		result, _ := extractor.Extract(protoV2.MessageV2(event))

		// assert
		assert.Equal(t, expectedResult, result)
	})

	t.Run("can handle non-wrapper repeated items", func(t *testing.T) {
		extractor := RepeatedItem{
			RepeatedFieldPath: "list",
			ItemExtractor:     B("")}

		event := &example.Order{
			List: []string{"feature 1", "feature 2"},
		}

		expectedResult := []interface{}{
			"feature 1", "feature 2",
		}

		// act
		result, _ := extractor.Extract(protoV2.MessageV2(event))

		// assert
		assert.Equal(t, expectedResult, result)
	})
}

func TestBasicFieldExtractor(t *testing.T) {
	t.Run("extracts a value at the top level", func(t *testing.T) {
		extractor := BasicField("id")

		event := &example.Order{
			Id: wrapperspb.Int64(int64(123)),
		}

		// act
		result, _ := extractor.Extract(protoV2.MessageV2(event))

		// assert
		assert.Equal(t, int64(123), result)

	})

	t.Run("extracts a nested field value", func(t *testing.T) {
		extractor := BasicField("cart.id")
		accountID := int64(123)

		event := &example.Order{
			Cart: &example.Cart{
				Id: wrapperspb.Int64(accountID),
			},
		}

		// act
		result, _ := extractor.Extract(protoV2.MessageV2(event))

		// assert
		assert.Equal(t, int64(123), result)
	})

	t.Run("handles non-existent field", func(t *testing.T) {
		extractor := BasicField("cart.made_up_field")
		accountID := int64(123)

		event := &example.Order{
			Cart: &example.Cart{
				Id: wrapperspb.Int64(accountID),
			},
		}

		// act
		result, _ := extractor.Extract(protoV2.MessageV2(event))

		// assert
		assert.Nil(t, result)
	})

	t.Run("extracts without unwrapping the value", func(t *testing.T) {
		extractor := BasicField("cart")
		accountID := int64(123)
		cart := &example.Cart{
			Id: wrapperspb.Int64(accountID),
		}
		event := &example.Order{
			Cart: cart,
		}

		// act
		result, _ := extractor.ExtractPath(protoV2.MessageV2(event))

		// assert
		assert.Equal(t, cart, result)
	})

	t.Run("handles a nil at the top level", func(t *testing.T) {
		extractor := BasicField("cart.id")
		event := &example.Order{
			Cart: nil,
		}

		// act
		result, _ := extractor.ExtractPath(protoV2.MessageV2(event))

		// assert
		assert.Nil(t, result)
	})

	t.Run("handles a nil at some point on the path", func(t *testing.T) {
		extractor := BasicField("cart.id")
		cart := &example.Cart{
			Id: nil,
		}
		event := &example.Order{
			Cart: cart,
		}

		// act
		result, _ := extractor.ExtractPath(protoV2.MessageV2(event))

		// assert
		assert.Nil(t, result)
	})

	t.Run("extracts a non-wrapper type", func(t *testing.T) {
		extractor := BasicField("opaque")
		event := &example.Order{
			Opaque: "my_feature",
		}

		// act
		result, _ := extractor.ExtractPath(protoV2.MessageV2(event))

		// assert
		assert.Equal(t, "my_feature", result)
	})

	t.Run("extracts an enum to the string value", func(t *testing.T) {
		extractor := BasicField("kind")
		event := &example.Order{
			Kind: example.Kind_KINDA,
		}

		// act
		result, _ := extractor.Extract(event.ProtoReflect().Interface())

		// assert
		assert.Equal(t, "kinda", result)
	})
}

func TestOneOfExractor(t *testing.T) {
	t.Run("extracts a basic oneof value", func(t *testing.T) {
		extractor := OneOf("event.orderupdated", B("cart.id"))

		event := &example.Order{
			Event: &example.Order_Orderupdated{
				Orderupdated: &example.OrderUpdated{
					Cart: &example.Cart{
						Id: wrapperspb.Int64(int64(123)),
					}}}}

		result, _ := extractor.Extract(protoV2.MessageV2(event))

		assert.Equal(t, int64(123), result)
	})

	t.Run("extracts using an arbitrary extractor", func(t *testing.T) {
		extractor := OneOf("event.orderupdated", ObjectExtractor{
			"id": B("cart.id"),
		})

		event := &example.Order{
			Event: &example.Order_Orderupdated{
				Orderupdated: &example.OrderUpdated{
					Cart: &example.Cart{
						Id: wrapperspb.Int64(int64(123)),
					}}}}
		expectedResult := map[string]interface{}{
			"id": int64(123),
		}

		result, _ := extractor.Extract(protoV2.MessageV2(event))

		assert.Equal(t, expectedResult, result)
	})

	t.Run("returns nil when given path is not a oneof", func(t *testing.T) {
		extractor := OneOf("cart", ObjectExtractor{
			"id":    B("order.id"),
			"token": B("order.token"),
		})

		event := &example.Order{
			Cart: &example.Cart{
				Id: wrapperspb.Int64(121),
			},
		}

		result, _ := extractor.Extract(protoV2.MessageV2(event))

		assert.Nil(t, result)
	})

	t.Run("returns nil when the oneof value is nil", func(t *testing.T) {
		extractor := OneOf("event.orderupdated", ObjectExtractor{
			"id":    B("cart.id"),
			"token": B("cart.token"),
		})

		event := &example.Order{
			Event: nil,
		}

		result, _ := extractor.Extract(protoV2.MessageV2(event))

		assert.Nil(t, result)
	})

	t.Run("returns nil when the given oneof path doesn't match the set value", func(t *testing.T) {
		extractor := OneOf("event.ordercreated", ObjectExtractor{
			"id":    B("order.id"),
			"token": B("order.token"),
		})

		event := &example.Order{
			Event: &example.Order_Orderupdated{
				Orderupdated: &example.OrderUpdated{
					Cart: &example.Cart{
						Id: wrapperspb.Int64(int64(123)),
					}}}}

		result, _ := extractor.Extract(protoV2.MessageV2(event))

		assert.Nil(t, result)
	})
}

func TestExractorWithDefault(t *testing.T) {
	t.Run("returns a provided default if the extracted value is nil", func(t *testing.T) {
		extractor := B("token")
		defaultExtractor := ExtractorWithDefault{
			DefaultValue: "default",
			Extractor:    extractor,
		}

		event := &example.Order{}

		result, _ := defaultExtractor.Extract(protoV2.MessageV2(event))

		assert.Equal(t, string("default"), result)
	})

	t.Run("returns a provided default if the extracted value is 0", func(t *testing.T) {
		extractor := B("id")
		defaultExtractor := ExtractorWithDefault{
			DefaultValue: 31,
			Extractor:    extractor,
		}

		event := &example.Order{}

		result, _ := defaultExtractor.Extract(protoV2.MessageV2(event))

		assert.Equal(t, int(31), result)
	})

	t.Run("returns a provided default if the extracted value is an emptry array", func(t *testing.T) {
		extractor := RepeatedItem{RepeatedFieldPath: "lineitems",
			ItemExtractor: ObjectExtractor{
				"id":       B("id"),
				"quantity": B("quantity"),
			},
		}

		defaultExtractor := ExtractorWithDefault{
			DefaultValue: []map[string]interface{}{
				{
					"id":       -1,
					"quantity": 2,
				},
			},
			Extractor: extractor,
		}
		event := &example.Order{
			Lineitems: []*example.LineItem{},
		}

		result, _ := defaultExtractor.Extract(protoV2.MessageV2(event))

		assert.Equal(t, []map[string]interface{}{
			{
				"id":       -1,
				"quantity": 2,
			},
		}, result)
	})

	t.Run("returns a provided default if the extracted value is blank", func(t *testing.T) {
		extractor := B("opaque")
		defaultExtractor := ExtractorWithDefault{
			DefaultValue: "default",
			Extractor:    extractor,
		}
		event := &example.Order{
			Opaque: "",
		}

		result, _ := defaultExtractor.Extract(protoV2.MessageV2(event))

		assert.Equal(t, string("default"), result)
	})

	t.Run("returns the extracted value when it is not nil", func(t *testing.T) {
		extractor := B("token")
		defaultExtractor := ExtractorWithDefault{
			DefaultValue: "default",
			Extractor:    extractor,
		}
		event := &example.Order{
			Token: wrapperspb.String("pending"),
		}

		result, _ := defaultExtractor.Extract(protoV2.MessageV2(event))

		assert.Equal(t, string("pending"), result)
	})
}

func TestTemplate(t *testing.T) {
	t.Run("returns a templated string with single argument extractor", func(t *testing.T) {
		extractor := B("id")
		templateExtractor := Template{
			format: "Order updated - %d",
			args:   []FieldExtractor{extractor},
		}

		event := &example.Order{
			Id: wrapperspb.Int64(12414),
		}

		result, _ := templateExtractor.Extract(protoV2.MessageV2(event))

		assert.Equal(t, string("Order updated - 12414"), result)
	})

	t.Run("returns a templated string with multiple argument extractors", func(t *testing.T) {
		orderNoExtractor := B("id")
		orderTokenExtractor := B("token")
		templateExtractor := Template{
			format: "Order updated - %d : %s",
			args:   []FieldExtractor{orderNoExtractor, orderTokenExtractor},
		}

		event := &example.Order{
			Id:    wrapperspb.Int64(12414),
			Token: wrapperspb.String("my token"),
		}

		result, _ := templateExtractor.Extract(protoV2.MessageV2(event))

		assert.Equal(t, string("Order updated - 12414 : my token"), result)
	})

	t.Run("returns a best effort templated string when an argument is missing", func(t *testing.T) {
		templateExtractor := Template{
			format: "Order updated - %d",
			args:   []FieldExtractor{},
		}

		event := &example.Order{
			Id: wrapperspb.Int64(11211),
		}

		result, _ := templateExtractor.Extract(protoV2.MessageV2(event))

		assert.Equal(t, string("Order updated - %!d(MISSING)"), result)
	})

	t.Run("returns a best effort templated string when an extracted value is `nil`", func(t *testing.T) {
		orderNoExtractor := B("id")
		orderTokenExtractor := B("token")
		templateExtractor := Template{
			format: "Order updated - %d : %s",
			args:   []FieldExtractor{orderNoExtractor, orderTokenExtractor},
		}

		event := &example.Order{
			Id:    wrapperspb.Int64(12414),
			Token: nil,
		}

		result, _ := templateExtractor.Extract(protoV2.MessageV2(event))

		assert.Equal(t, string("Order updated - 12414 : %!s(<nil>)"), result)
	})

	t.Run("returns a best effort templated string when an attribute specified for extraction is absent", func(t *testing.T) {
		orderNoExtractor := B("id")
		orderTokenExtractor := B("token")
		templateExtractor := Template{
			format: "Order updated - %d : %s",
			args:   []FieldExtractor{orderNoExtractor, orderTokenExtractor},
		}

		event := &example.Order{
			Id: wrapperspb.Int64(12414),
		}

		result, _ := templateExtractor.Extract(protoV2.MessageV2(event))

		assert.Equal(t, string("Order updated - 12414 : %!s(<nil>)"), result)
	})
}
