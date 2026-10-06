package revkeen

import (
	"encoding/json"
	"testing"
)

type nullableInt64Wire struct {
	Amount NullableInt64 `json:"amount"`
}

func TestNullableInt64DistinguishesAbsentNullAndZeroOnWire(t *testing.T) {
	tests := []struct {
		name      string
		body      string
		wantSet   bool
		wantNil   bool
		wantValue int64
	}{
		{name: "absent", body: `{}`, wantNil: true},
		{name: "null", body: `{"amount":null}`, wantSet: true, wantNil: true},
		{name: "zero", body: `{"amount":0}`, wantSet: true, wantValue: 0},
		{name: "value", body: `{"amount":42}`, wantSet: true, wantValue: 42},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var wire nullableInt64Wire
			if err := json.Unmarshal([]byte(test.body), &wire); err != nil {
				t.Fatal(err)
			}
			if wire.Amount.IsSet() != test.wantSet {
				t.Fatalf("IsSet = %v; want %v", wire.Amount.IsSet(), test.wantSet)
			}
			value := wire.Amount.Get()
			if (value == nil) != test.wantNil {
				t.Fatalf("nil = %v; want %v", value == nil, test.wantNil)
			}
			if value != nil && *value != test.wantValue {
				t.Fatalf("value = %d; want %d", *value, test.wantValue)
			}
		})
	}
}

func TestNullableInt64MarshalsExplicitNullAndZero(t *testing.T) {
	zero := int64(0)
	for _, test := range []struct {
		name string
		wire NullableInt64
		want string
	}{
		{name: "explicit null", wire: *NewNullableInt64(nil), want: "null"},
		{name: "zero", wire: *NewNullableInt64(&zero), want: "0"},
	} {
		t.Run(test.name, func(t *testing.T) {
			encoded, err := json.Marshal(test.wire)
			if err != nil {
				t.Fatal(err)
			}
			if string(encoded) != test.want {
				t.Fatalf("encoded as %s; want %s", encoded, test.want)
			}
		})
	}
}
