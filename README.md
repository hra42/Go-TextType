# Go-TextType

> [!Important]
> Download the latest version from the [release page](https://github.com/HRA42/Go-TextType/releases)
> and run the executable.  
> Visit the [docs](https://go-texttype.postrausch.tech/) for more information.

![Icon](icon.jpeg)

This program is designed to print text from the clipboard using Keyboard events.
Password entry into console applications is enabled by this feature.
Version 1.1.1 adds the ability to remove whitespace from the clipboard and
pressing entered after the input is entered, if you select the option from the menu bar.
From version 0.1.2 you can switch the hotkey between: 
- ISO Layout: `left ctrl + left shift + v`
- DE Layout: `links strg + links shift + v`  
or  
- ISO Layout: `left ctrl + left alt + s`
- DE Layout: `links strg + links alt + s`  
or  
- ISO Layout: `left ctrl + Q`
- DE Layout: `links strg + Q`

The HotKey is stored in a file called `hotkey.gob` within the same folder you run the application.

If you need any additional HotKeys or have some issues to report, please contact me or create an issue on
[GitHub](https://github.com/HRA42/Go-TextType/issues).

This project contains a program implemented in Go (version 1.21) using the Go SDK (version 1.21.1).
It is designed to print text from the clipboard using Keyboard events.

## Functionality
The main functionality of this program is as described:
The application is designed to print the text stored in the clipboard.
It achieves this by calling a hotkey combination that triggers a function in the program.
This triggered function fetches the content from the clipboard and prints it on the console.

## Dependencies
The required Go packages for this project to function include:
- embed for embedding files in executable.
- encoding/gob for encoding the hotkey and enter preference to a file.
- github.com/gen2brain/beeep for update notifications.
- github.com/getlantern/systray for system tray functionality.
- github.com/go-vgo/robotgo for simulating keyboard inputs.
- github.com/pkg/browser for opening a browser to the release page on GH.
- github.com/tcnksm/go-latest for getting the latest version from GH.
- golang.design/x/clipboard for access to the system clipboard.
- golang.design/x/hotkey for creating and managing hotkeys.

## Running the Program
Download the latest version of the application from the release page.

## Compilation
Get the source code from GitHub:
```Bash
git clone https://github.com/HRA42/Go-TextType.git
```

Download the dependencies:
```Shell
Powershell.exe -File .\build.ps1 dep
```

After you install the dependencies, run:
```Bash
Powershell.exe -File .\build.ps1 build
```

When you are ready to start a new build, run:
```Bash
Powershell.exe -File .\build.ps1 clean
```

## Troubleshooting
The program logs the current version of the application, and the Build ID in its log file TextType.log.
You may refer to this file for troubleshooting and referencing specific versions of the application.
