package main

import (
	"./digits"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	showSeparator := true
	for {
		clearScreen()

		currentTime := time.Now()
		hour := [...]int{currentTime.Hour() / 10, currentTime.Hour() % 10}
		minute := [...]int{currentTime.Minute() / 10, currentTime.Minute() % 10}
		second := [...]int{currentTime.Second() / 10, currentTime.Second() % 10}

		for i := 0; i < 5; i++ {
			separator := " "
			if showSeparator && (i == 1 || i == 3) {
				separator = digits.Separator
			}

			fmt.Printf("  %s  %s  %s  %s  %s  %s  %s  %s\n",
				getNumberLine(i, hour[0]), getNumberLine(i, hour[1]),
				separator,
				getNumberLine(i, minute[0]), getNumberLine(i, minute[1]),
				separator,
				getNumberLine(i, second[0]), getNumberLine(i, second[1]),
			)
		}
		showSeparator = !showSeparator
		time.Sleep(time.Second)

	}
}

func getNumberLine(l int, n int) string {
	return digits.Numbers[n][l]
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
