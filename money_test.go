package revkeen

import (
	"encoding/json"
	"math"
	"strconv"
	"testing"
)

func TestMoneyMinorExactJSONRoundTrip(t *testing.T) {
	for _, value := range []int64{math.MinInt64, -1, 0, 1, math.MaxInt64} {
		amount := NewMoneyMinor(value)
		encoded, err := json.Marshal(amount)
		if err != nil {
			t.Fatal(err)
		}
		if string(encoded) != strconv.FormatInt(value, 10) {
			t.Fatalf("encoded %d as %s", value, encoded)
		}
		var decoded MoneyMinor
		if err := json.Unmarshal(encoded, &decoded); err != nil {
			t.Fatal(err)
		}
		if decoded.Int64() != value {
			t.Fatalf("decoded %d as %d", value, decoded.Int64())
		}
	}
}

func TestMoneyMinorRejectsNonIntegerOrUnsafeInput(t *testing.T) {
	invalidJSON := []string{`null`, `1.0`, `1.5`, `1e3`, `"100"`, `true`, `9223372036854775808`, `-9223372036854775809`, `01`, `1 true`}
	for _, input := range invalidJSON {
		t.Run(input, func(t *testing.T) {
			amount := MoneyMinor(77)
			if err := json.Unmarshal([]byte(input), &amount); err == nil {
				t.Fatalf("accepted unsafe money JSON %q", input)
			}
			if amount != 77 {
				t.Fatalf("failed decode mutated receiver to %d", amount)
			}
		})
	}

	for _, input := range []string{"", " 1", "+1", "01", "1.0", "1e3", "9223372036854775808"} {
		if _, err := ParseMoneyMinor(input); err == nil {
			t.Fatalf("ParseMoneyMinor accepted %q", input)
		}
	}
}

func TestParseMoneyMinorAcceptsCanonicalIntegers(t *testing.T) {
	for _, input := range []string{"-9223372036854775808", "-1", "-0", "0", "1", "9223372036854775807"} {
		if _, err := ParseMoneyMinor(input); err != nil {
			t.Fatalf("ParseMoneyMinor(%q): %v", input, err)
		}
	}
}
