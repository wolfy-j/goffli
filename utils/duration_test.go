package utils

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"time"
)

func Test_Duration_Time(t *testing.T) {
	assert.Equal(t, FormatDuration(time.Second), "00:00:01")
	assert.Equal(t, FormatDuration(time.Minute), "00:01:00")
	assert.Equal(t, FormatDuration(time.Hour), "01:00:00")
	assert.Equal(t, FormatDuration(45*time.Hour+61*time.Minute+30*time.Second), "46:01:30")
}

func Test_Duration_Float(t *testing.T) {
	assert.Equal(t, FormatDuration(1.0), "00:00:01")
	assert.Equal(t, FormatDuration(60.0), "00:01:00")
	assert.Equal(t, FormatDuration(3600.0), "01:00:00")
	assert.Equal(t, FormatDuration(3600.0*45+61*60.0+30.0), "46:01:30")
}

func Test_Duration_Int(t *testing.T) {
	assert.Equal(t, FormatDuration(1), "00:00:01")
	assert.Equal(t, FormatDuration(60), "00:01:00")
	assert.Equal(t, FormatDuration(3600), "01:00:00")
	assert.Equal(t, FormatDuration(3600*45+61*60+30), "46:01:30")
}

func Test_Duration_Fallback(t *testing.T) {
	assert.Equal(t, FormatDuration("value"), "value")
}
