package main

import (
	"fmt"
	"time"
	"flag"
	"github.com/gen2brain/beeep"
)

func main() {
	// Define the flag --times
	timesFlag := flag.Int("times", 1, "defines how many times you want to run the pomocli timer, after 4 pomodoros the break time may increase")

	flag.Parse()

	for i := 1; i <= *timesFlag; i++ {
		// Starts the pomodoro timer and notifies the user 
		fmt.Printf("Pomodoro %v started\n", i)
		beeep.Notify("PomoCLI", "Pomodoro started", "PomoCLI/src/assets/clock.png")
		fmt.Println("25 minutes of focus!")
		timer := time.NewTimer(25 * time.Minute)

		// When the timer for focus ends, the break timer starts
		<-timer.C
		brTime := 5 * time.Minute
		if i > 4 {
			brTime = 15 * time.Minute
		}
		fmt.Printf("Break Time of %v!\n", brTime)
		beeep.Notify("PomoCLI", "Break Time!", "PomoCLI/src/assets/break.jpg")
		breakTimer := time.NewTimer(brTime)

		<-breakTimer.C
		fmt.Println("Break time finished!")
		fmt.Println("============================================")
	} 
}
