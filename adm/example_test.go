package textcal_test

import (
	"fmt"
	"time"
	"github.com/mschilli/go-textcal"
	"github.com/fatih/color"
	"testing"
)

func TestBasic(t *testing.T) {
	month := time.Date(2025, time.November, 1, 0, 0, 0, 0, time.UTC)
	cal := textcal.New(month)
	fmt.Println(cal.String())
}

func TestReverse(t *testing.T) {
	month := time.Date(2025, time.November, 1, 0, 0, 0, 0, time.UTC)
	cal := textcal.New(month)

	cal.UseFormatter(11, cal.ReverseFormatter())
	fmt.Println(cal.String())
}

func TestAnnotation(t *testing.T) {
	month := time.Date(2025, time.November, 1, 0, 0, 0, 0, time.UTC)
	cal := textcal.New(month)

	makegreen := cal.ColorFormatter(color.FgGreen, color.Reset)
	makered := cal.ColorFormatter(color.FgRed, color.Reset)

	cal.UseFormatter(11, makegreen)
	cal.Annotate(11, makegreen("$100"))

	cal.UseFormatter(13, makered)
	cal.Annotate(13, makered("$200"))

	fmt.Println(cal.String())
}
