package test

import (
	"fmt"
	"time"
)

func strPtr(v string) *string {
	return &v
}

func boolPtr(v bool) *bool {
	return &v
}

func intPtr(v int) *int {
	return &v
}

func interfacePtr(v interface{}) *interface{} {
	return &v
}

func toPtr[T any](v T) *T {
	return &v
}

func newDate() string {
	t := time.Now().Add(time.Hour)

	return fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
}

func dateFromTime(t time.Time) string {
	return fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
}

func timePtr(v time.Time) *time.Time {
	return &v
}
