package util

import "time"

func CheckZeroArrayString(value []string, set func(value []string)) {
	if value != nil {
		set(value)
	}
}

func CheckUpdateArrayString(value []string, target *[]string) {
	CheckZeroArrayString(value, func(value []string) {
		*target = value
	})
}

func CheckUpdateString(value string, target *string) {
	CheckZeroString(value, func(value string) {
		*target = value
	})
}

func CheckZeroString(value string, set func(value string)) {
	if value != "" {
		set(value)
	}
}

func CheckUpdateInt64(value int64, target *int64) {
	CheckZeroInt64(value, func(value int64) {
		*target = value
	})
}

func CheckZeroInt64(value int64, set func(value int64)) {
	if value != 0 {
		set(value)
	}
}

func CheckUpdateInt(value int, target *int) {
	CheckZeroInt(value, func(value int) {
		*target = value
	})
}
func CheckZeroInt(value int, set func(value int)) {
	if value != 0 {
		set(value)
	}
}

func CheckUpdateTime(value time.Time, target *time.Time) {
	CheckZeroTime(value, func(value time.Time) {
		*target = value
	})
}
func CheckZeroTime(value time.Time, set func(value time.Time)) {
	if !value.IsZero() {
		set(value)
	}
}

func CheckUpdateFloat32(value float32, target *float32) {
	CheckZeroFloat32(value, func(value float32) {
		*target = value
	})
}

func CheckZeroFloat32(value float32, set func(value float32)) {
	if value != 0 {
		set(value)
	}
}

func CheckUpdateFloat64(value float64, target *float64) {
	CheckZeroFloat64(value, func(value float64) {
		*target = value
	})
}

func CheckZeroFloat64(value float64, set func(value float64)) {
	if value != 0 {
		set(value)
	}
}
