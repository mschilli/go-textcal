package textcal

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Calendar struct {
	month         time.Time
	AnnotationMap map[int]Annotation
}

type Annotation struct {
	dayFormatter func(string) string
	text         string
}

func New(month time.Time) *Calendar {
	c := Calendar{
		month:         month,
		AnnotationMap: map[int]Annotation{},
	}

	return &c
}

func (c *Calendar) String() string {
	out := ""

	year, month, _ := c.month.Date()

	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	startWeekday := firstDay.Weekday()

	daysInMonth := daysInMonth(year, month)

	out += fmt.Sprintf("    %s %d\n", month, year)
	out += fmt.Sprintf("Su Mo Tu We Th Fr Sa\n")

	// padding
	for i := 0; i < int(startWeekday); i++ {
		out += fmt.Sprintf("   ")
	}

	curDate := firstDay
	annos := []string{}

	for day := 1; day <= daysInMonth; day++ {
		dayStr := fmt.Sprintf("%2d", day)

		if f := c.AnnotationMap[day].dayFormatter; f != nil {
			dayStr = f(dayStr)
		}

		out += dayStr + " "

		if text := c.AnnotationMap[day].text; text != "" {
			annos = append(annos, text)
		}

		if day == daysInMonth {
			// last row, pad till Saturday (for annotations)
			wd := int(curDate.Weekday())
			target := int(time.Saturday)
			for i := 0; i < (target-wd+7)%7; i++ {
				out += fmt.Sprintf("   ")
			}
		}

		if (int(startWeekday)+day)%7 == 0 || day == daysInMonth {
			if len(annos) != 0 {
				out += strings.Join(annos, ", ")
				annos = []string{}
			}
			out += "\n"
		}

		curDate = curDate.AddDate(0, 0, 1)
	}

	return out
}

func daysInMonth(year int, month time.Month) int {
	nextMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
	return nextMonth.Day()
}

func (c *Calendar) ColorFormatter(
	fg color.Attribute, bg color.Attribute) func(string) string {
	return func(str string) string {
		fu := color.New(bg, fg).SprintFunc()
		return fu(str)
	}
}
