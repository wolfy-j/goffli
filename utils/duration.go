package utils

import (
	"fmt"
	"time"
)

// FormatDuration renders human friendly duration string.
func FormatDuration(d interface{}) string {
	var t time.Duration
	switch d.(type) {
	case float64:
		t = time.Second * time.Duration(d.(float64))
	case int64:
		t = time.Second * time.Duration(d.(int64))
	case int:
		t = time.Second * time.Duration(d.(int))
	case time.Duration:
		t = d.(time.Duration)
	default:
		return fmt.Sprintf("%v", d)
	}

	h := int(t.Hours())
	i := int(t.Minutes()) % 60
	s := t.Seconds() - float64(h*3600+i*60)

	return fmt.Sprintf("%02d:%02d:%02.0f", h, i, s)
}
