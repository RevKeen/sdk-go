package revkeen

import (
	"encoding/json"
	"testing"
	"time"
)

func TestParseRFC3339ReturnsTime(t *testing.T) {
	parsed, err := ParseRFC3339("2026-08-21T12:34:56.123456789+01:00")
	if err != nil {
		t.Fatal(err)
	}
	var _ time.Time = parsed
	if parsed.Nanosecond() != 123456789 || parsed.Format(time.RFC3339Nano) != "2026-08-21T12:34:56.123456789+01:00" {
		t.Fatalf("unexpected parsed timestamp: %s", parsed.Format(time.RFC3339Nano))
	}
}

func TestRFC3339TimeJSONRoundTrip(t *testing.T) {
	var value RFC3339Time
	if err := json.Unmarshal([]byte(`"2026-08-21T12:34:56Z"`), &value); err != nil {
		t.Fatal(err)
	}
	if value.Time.Location() != time.UTC {
		t.Fatalf("timestamp location = %v; want UTC", value.Time.Location())
	}
	encoded, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	if string(encoded) != `"2026-08-21T12:34:56Z"` {
		t.Fatalf("timestamp encoded as %s", encoded)
	}
}

func TestRFC3339TimeRejectsInvalidWireValues(t *testing.T) {
	for _, input := range []string{
		`null`,
		`123`,
		`"2026-08-21"`,
		`"2026-08-21T12:34:56"`,
		`" 2026-08-21T12:34:56Z"`,
		`"2026-02-30T12:34:56Z"`,
		`"2026-08-21t12:34:56z"`,
		`"2026-08-21T12:34:56,1Z"`,
		`"2026-08-21T12:34:56+24:00"`,
		`"2026-08-21T12:34:56+00:60"`,
	} {
		t.Run(input, func(t *testing.T) {
			var value RFC3339Time
			if err := json.Unmarshal([]byte(input), &value); err == nil {
				t.Fatalf("accepted invalid timestamp %s", input)
			}
		})
	}
}
