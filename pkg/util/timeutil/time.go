package timeutil

import "time"

func Format(t *time.Time, str **string) bool {
	if t == nil {
		return false
	}

	format := t.Format("02.01.2006")
	*str = &format

	return true
}

func Parse(str *string, t **time.Time) bool {
	if str == nil {
		return false
	}

	parse, err := time.Parse("02.01.2006", *str)
	if err != nil {
		return false
	}

	*t = &parse

	return true
}
