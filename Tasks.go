package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func taskPage() fyne.CanvasObject{
	taskNO:=0
	taskList := widget.NewList(
		func() int{
			return taskNO
		},
		func() fyne.CanvasObject{},
		func(lii widget.ListItemID, co fyne.CanvasObject) {}
	)
	splitView := container.NewHSplit()
}

func main() {
	myApp := app.New()
	wind := myApp.NewWindow("Tasks")
	wind.SetContent(taskPage())
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
