# go-textcal

This Go package creates ASCII calendars for terminals, similar to
the `cal` Unix command line utility.

```
   November 2025      
Su Mo Tu We Th Fr Sa  
                   1  
 2  3  4  5  6  7  8  
 9 10 11 12 13 14 15  
16 17 18 19 20 21 22  
23 24 25 26 27 28 29  
30                    
```

## Example
```
package main

import (
	"github.com/mschilli/go-textcal"
	"fmt"
	"time"
)

func main() {
	textcal := textcal.New(time.Now())
	fmt.PrintLn(textcal.String())
}
```

## Author

Mike Schilli, m@perlmeister.com 2025

## License

Released under the [Apache 2.0](LICENSE)
