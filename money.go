package revkeen

import (
	"bytes"
	"fmt"
	"strconv"
)

// MoneyMinor is an exact signed amount in a currency's minor unit. Currency is
// deliberately carried by the surrounding API model; MoneyMinor never performs
// currency conversion or rounding.
type MoneyMinor int64

// NewMoneyMinor constructs an exact minor-unit amount from an integer.
func NewMoneyMinor(value int64) MoneyMinor {
	return MoneyMinor(value)
}

// ParseMoneyMinor accepts only a base-10 integer. Decimal points, exponents,
// whitespace and values outside int64 are rejected rather than rounded.
func ParseMoneyMinor(value string) (MoneyMinor, error) {
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("revkeen: money minor must be a signed 64-bit integer: %w", err)
	}
	if strconv.FormatInt(parsed, 10) != value && !(parsed == 0 && value == "-0") {
		return 0, fmt.Errorf("revkeen: money minor must use canonical base-10 integer syntax")
	}
	return MoneyMinor(parsed), nil
}

// Int64 returns the exact minor-unit value.
func (amount MoneyMinor) Int64() int64 {
	return int64(amount)
}

func (amount MoneyMinor) MarshalJSON() ([]byte, error) {
	return strconv.AppendInt(nil, int64(amount), 10), nil
}

func (amount *MoneyMinor) UnmarshalJSON(data []byte) error {
	if amount == nil {
		return fmt.Errorf("revkeen: cannot unmarshal money minor into nil receiver")
	}
	data = bytes.TrimSpace(data)
	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		return fmt.Errorf("revkeen: money minor must be a JSON integer")
	}
	parsed, err := ParseMoneyMinor(string(data))
	if err != nil {
		return fmt.Errorf("revkeen: money minor must be a JSON integer: %w", err)
	}
	*amount = parsed
	return nil
}
