// example that prints out a timestamp including some information
// about the position of the Moon in the zodiac
// example output:
// 11:11PM Luna: New Moon 0.65%, Scorpio 11.957°
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dyindude/moonphase"
)

func main() {
	t := time.Now()
	m := moonphase.New(t)
	fmt.Printf("%s Luna: %s %.2f%%, %s %.3f°\n", t.Format(time.Kitchen), m.PhaseName(), m.Illumination()*100, strings.Title(m.ZodiacSignTropical()), m.DegreesInSignTropical())
}
