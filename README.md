# godate
## Installation
Use `go get -u github.com/xiroxasx/godate` to get started.  

## Formats
Define a format that suites your needs like in the example down below.  
Currenty supported:  

| Format   | Example   | Description                                |
|----------|-----------|--------------------------------------------|
| YYYY     | 2023      | The 4-digit year                           |
| YYY      | 172       | The day of the year                        |
| YY       | 23        | The 2-digit format of `YYYY`               |
| MMMM     | June      | The english name of the month              |
| MMM      | Jun       | The 3-letter abbriviation of `MMMM`        |
| MM       | 06        | The 2-digit format of the month            |
| M        | 6         | The 1 or 2-digit format of the month       |
| DDDD     | Wednesday | The english name of the day                |
| DDD      | Wed       | The 3-letter abbriviation of `DDDD`        |
| DD       | 21        | The 2-digit format of the day              |
| D        | 21        | The 1 or 2-digit format of the day         |
| hha      | 7:00PM    | The time in the 12 hour format             |
| hh       | 19        | The 2-digit format of the hour             |
| h        | 19        | The 1 or 2-digit format of the hour        |
| mm       | 00        | The 2-digit format of the minute           |
| m        | 0         | The 1 or 2-digit format of the minute      |
| ss       | 00        | The 2-digit format of the second           |
| s        | 0         | The 1 or 2-digit format of the second      |
| ms       | 0         | The 1 or 3-digit format of the millisecond |

## Usage

```go
package main

import (
	"time"

	"github.com/xiroxasx/godate"
)

func main() {
    d := godate.New(time.Now())
	fmt.Println(d.Format("YYYY-MM-DD@hh:mm:ss")) // Prints something like "2023-06-21@19:00:00"
}
```