package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	wind := myApp.NewWindow("Tasks")
	wind.SetContent()
	wind.ShowAndRun()
}

/*
In the main window, there's a small section containing the a clock as  it's counting.
The a Button for adding new tasks
Each task has a checkmark to the left of its text.
These tasks can be done in such a way that the time counts down or up from/to 24hrs

User Xperience
As a user I want to set a time based task and allow the timer count towards the set period

Backend
Time counter dis


*/
