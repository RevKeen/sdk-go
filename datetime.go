package revkeen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var strictRFC3339Pattern = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d+)?(?:Z|[+-]\d{2}:\d{2})$`)

// ParseRFC3339 parses an API timestamp into time.Time. A timezone is mandatory;
// date-only values and timestamps with surrounding whitespace are rejected.
func ParseRFC3339(value string) (time.Time, error) {
	if !strictRFC3339Pattern.MatchString(value) {
		return time.Time{}, fmt.Errorf("revkeen: timestamp must use strict RFC3339 syntax")
	}
	if value[len(value)-1] != 'Z' {
		offset := value[len(value)-6:]
		hours, hoursErr := strconv.Atoi(offset[1:3])
		minutes, minutesErr := strconv.Atoi(offset[4:6])
		if hoursErr != nil || minutesErr != nil || hours > 23 || minutes > 59 {
			return time.Time{}, fmt.Errorf("revkeen: timestamp has an invalid RFC3339 timezone offset")
		}
	}
	parsed, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		return time.Time{}, fmt.Errorf("revkeen: timestamp must be RFC3339: %w", err)
	}
	return parsed, nil
}

// RFC3339Time is an additive strict JSON wrapper around a real time.Time.
type RFC3339Time struct {
	time.Time
}

func NewRFC3339Time(value time.Time) RFC3339Time {
	return RFC3339Time{Time: value}
}

func (value RFC3339Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.Time.Format(time.RFC3339Nano))
}

func (value *RFC3339Time) UnmarshalJSON(data []byte) error {
	if value == nil {
		return fmt.Errorf("revkeen: cannot unmarshal timestamp into nil receiver")
	}
	data = bytes.TrimSpace(data)
	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		return fmt.Errorf("revkeen: timestamp must be an RFC3339 JSON string")
	}
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("revkeen: timestamp must be an RFC3339 JSON string: %w", err)
	}
	parsed, err := ParseRFC3339(raw)
	if err != nil {
		return err
	}
	value.Time = parsed
	return nil
}
