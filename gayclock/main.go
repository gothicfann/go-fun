package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

func main() {
	var blink bool
	var counter int

	screen.Clear()

	go func() {
		for {
			screen.MoveTopLeft()

			t := time.Now()
			h := t.Hour()
			m := t.Minute()
			s := t.Second()

			c := [8]placeholder{
				0: digits[h/10],
				1: digits[h%10],
				3: digits[m/10],
				4: digits[m%10],
				6: digits[s/10],
				7: digits[s%10],
			}

			if counter%2 == 0 {
				blink = !blink
			}

			if blink {
				c[2], c[5] = separator, separator
			} else {
				c[2], c[5] = void, void
			}

			var output string

			for i := 0; i < len(placeholder{}); i++ {
				for j, e := range c {
					rand.Seed(time.Now().UnixNano())
					randColor := rand.Intn(len(colors))
					if j == len(c)-1 {
						output += colors[randColor].Sprintf("%s\n", e[i])
					} else {
						output += colors[randColor].Sprintf("%s ", e[i])
					}
				}
			}

			fmt.Println(output)

			counter++
			time.Sleep(time.Second)
		}
	}()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()
	<-done
	fmt.Println()
	screen.Clear()
	screen.MoveTopLeft()
	fmt.Println("Exited")
}

type placeholder [5]string

var (
	zero = placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one = placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two = placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three = placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four = placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five = placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six = placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven = placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight = placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine = placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	separator = placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	void = placeholder{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	}

	colors = [6]*color.Color{
		color.New(color.FgGreen),
		color.New(color.FgBlue),
		color.New(color.FgCyan),
		color.New(color.FgRed),
		color.New(color.FgMagenta),
		color.New(color.FgYellow),
	}

	digits = [10]placeholder{
		0: zero,
		1: one,
		2: two,
		3: three,
		4: four,
		5: five,
		6: six,
		7: seven,
		8: eight,
		9: nine,
	}
)
