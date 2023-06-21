package godate

import (
	"fmt"
	"strings"
	"time"
)

type Date struct {
	time time.Time
}

type format struct {
	layout string
	value  string
}

func New(t time.Time) Date {
	return Date{time: t}
}

func (d *Date) Format(layout string) (t string) {
	t = layout

	// Day layouts need to be evaluated at the end to prevent replacing parts of the weekday.
	formats := []format{
		{layout: "YYYY", value: fmt.Sprintf("%d", d.time.Year())},
		{layout: "YYY", value: fmt.Sprintf("%d", d.time.YearDay())},
		{layout: "YY", value: fmt.Sprintf("%d", d.time.Year())[2:]},
		{layout: "MMMM", value: d.time.Month().String()},
		{layout: "MMM", value: d.time.Month().String()[:3]},
		{layout: "MM", value: fmt.Sprintf("%02d", int(d.time.Month()))},
		{layout: "M", value: fmt.Sprintf("%d", int(d.time.Month()))},
		{layout: "hha", value: d.time.Format(time.Kitchen)},
		{layout: "hh", value: fmt.Sprintf("%02d", d.time.Hour())},
		{layout: "h", value: fmt.Sprintf("%d", d.time.Hour())},
		{layout: "ms", value: fmt.Sprintf("%d", int32(d.time.Nanosecond())/int32(time.Millisecond))},
		{layout: "mm", value: fmt.Sprintf("%02d", d.time.Minute())},
		{layout: "m", value: fmt.Sprintf("%d", d.time.Minute())},
		{layout: "ss", value: fmt.Sprintf("%02d", d.time.Second())},
		{layout: "s", value: fmt.Sprintf("%d", d.time.Second())},
		{layout: "DDDD", value: d.time.Weekday().String()},
		{layout: "DDD", value: d.time.Weekday().String()[:3]},
		{layout: "DD", value: fmt.Sprintf("%02d", d.time.Day())},
		{layout: "D", value: fmt.Sprintf("%d", d.time.Day())},
	}
	for _, f := range formats {
		t = strings.ReplaceAll(t, f.layout, f.value)
	}
	return
}
