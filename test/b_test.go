package test

import (
	"context"
	"testing"
	"time"
)

func TestB(t *testing.T) {
	tn := time.Now().AddDate(0, 0, -2)
	t.Log(tn.Format("20060102"))
	year, week := tn.ISOWeek()
	t.Log(year, week)

}

func f(ctx context.Context) {
	context.WithValue(ctx, "foo", -6)
}
