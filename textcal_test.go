package textcal

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
	"testing"
	"time"
)

var debug = os.Getenv("DEBUG") != ""

func TestBasic(t *testing.T) {
	month := time.Date(2025, time.November, 1, 0, 0, 0, 0, time.UTC)
	cal := New(month)
	str := cal.String()

	exp := `    November 2025
Su Mo Tu We Th Fr Sa
                   1 
 2  3  4  5  6  7  8 
 9 10 11 12 13 14 15 
16 17 18 19 20 21 22 
23 24 25 26 27 28 29 
30                   
`

	t.Log(fmt.Sprintf("out: [%s] exp: [%s]", str, exp))

	if str != exp {
		t.Fail()
	}
}

func TestColor(t *testing.T) {
	month := time.Date(2025, time.November, 1, 0, 0, 0, 0, time.UTC)
	cal := New(month)

	cal.AnnotationMap = map[int]Annotation{
		10: {dayFormatter: cal.ColorFormatter(color.FgRed, color.BgBlack)},
	}
	str := cal.String()

	if debug {
		fmt.Println(str)
	}

	count := strings.Count(str, "\n")
	t.Log(fmt.Sprintf("out: [%s]", str))
	if count != 8 {
		t.Fail()
	}
}

func TestAnno(t *testing.T) {
	month := time.Date(2025, time.November, 1, 0, 0, 0, 0, time.UTC)
	cal := New(month)

	greenit := cal.ColorFormatter(color.FgGreen, color.Reset)
	redit := cal.ColorFormatter(color.FgRed, color.Reset)

	cal.AnnotationMap = map[int]Annotation{
		1: {dayFormatter: greenit,
			text: greenit("$1000"),
		},
		4: {dayFormatter: greenit,
			text: greenit("$100"),
		},
		5: {dayFormatter: redit,
			text: redit("$200"),
		},
		30: {dayFormatter: redit,
			text: redit("$500"),
		},
	}
	str := cal.String()

	if debug {
		fmt.Println(str)
	}

	count := strings.Count(str, "\n")
	t.Log(fmt.Sprintf("out: [%s]", str))
	if count != 8 {
		t.Fail()
	}
}
