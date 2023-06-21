package godate

import (
	"fmt"
	"testing"
	"time"

	r "github.com/stretchr/testify/require"
)

func TestDate(t *testing.T) {
	now := time.Now()
	d := New(now)

	r.Exactly(t, now.Format("2006"), d.Format("YYYY"))
	r.Exactly(t, fmt.Sprint(now.YearDay()), d.Format("YYY"))
	r.Exactly(t, fmt.Sprint(now.Year())[2:], d.Format("YY"))

	r.Exactly(t, now.Month().String(), d.Format("MMMM"))
	r.Exactly(t, now.Month().String()[:3], d.Format("MMM"))
	r.Exactly(t, fmt.Sprintf("%02d", int(now.Month())), d.Format("MM"))
	r.Exactly(t, fmt.Sprintf("%d", int(now.Month())), d.Format("M"))

	r.Exactly(t, now.Weekday().String(), d.Format("DDDD"))
	r.Exactly(t, now.Weekday().String()[:3], d.Format("DDD"))
	r.Exactly(t, fmt.Sprintf("%02d", now.Day()), d.Format("DD"))
	r.Exactly(t, fmt.Sprintf("%d", now.Day()), d.Format("D"))

	r.Exactly(t, fmt.Sprintf("%02d", now.Hour()), d.Format("hh"))
	r.Exactly(t, now.Format(time.Kitchen), d.Format("hha"))
	r.Exactly(t, fmt.Sprint(now.Hour()), d.Format("h"))

	r.Exactly(t, fmt.Sprintf("%02d", now.Minute()), d.Format("mm"))
	r.Exactly(t, fmt.Sprint(now.Minute()), d.Format("m"))

	r.Exactly(t, fmt.Sprintf("%02d", now.Second()), d.Format("ss"))
	r.Exactly(t, fmt.Sprint(now.Second()), d.Format("s"))

	r.Exactly(t, fmt.Sprintf("%d", int32(now.Nanosecond())/int32(time.Millisecond)), d.Format("ms"))
	r.Exactly(t, now.Format("2006-01-02 15:04:05"), d.Format("YYYY-MM-DD hh:mm:ss"))
}

func BenchmarkFormat(b *testing.B) {
	now := time.Now()
	d := New(now)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Format("YYYY")
	}
}
