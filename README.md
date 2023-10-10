# Go-TextType

![Icon](icon.jpeg)

This program is designed to print text from the clipboard using Keyboard events.
This project contains a program implemented in Go (version 1.21) using the Go SDK (version 1.21.1).
It is designed to print text from the clipboard using Keyboard events.
The application is developed specifically to used with GoLand 2023.2.2 running on a Windows 11 (amd64) system.

## Functionality
The main functionality of this program is as described:
The application is designed to print the text stored in the clipboard.
It achieves this by calling a hotkey combination that triggers a function in the program.
This triggered function fetches the content from the clipboard and prints it on the console.

## Dependencies
The required Go packages for this project to function include:
- embed for embedding files in executable.
- github.com/getlantern/systray for system tray functionality.
- github.com/go-vgo/robotgo for simulating keyboard inputs.
- golang.design/x/clipboard for access to the system clipboard.
- golang.design/x/hotkey for creating and managing hotkeys.

## Running the Program
Download the latest version of the application from the release page.

## Compilation
Download the code from GitHub and run for Windows:
```Bash
go build -ldflags '-extldflags "-static" -w -s -H windowsgui -X main.AppVersion=INSERT_VERSION -X main.BuildID=BUILDID' .
```

## Troubleshooting
The program logs the current version of the application, and the Build ID in its log file TextType.log.
You may refer to this file for troubleshooting and referencing specific versions of the application.