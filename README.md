# Go-TextType - DEVELOPMENT IS PAUSED

> [!Important]
> Work on this project is currently paused, as I don't have the need anymore. Teamviewer can now paste passwords via keyboard emulation. I might be getting back to this in the future.

[![Go Report Card](https://goreportcard.com/badge/github.com/hra42/Go-TextType)](https://goreportcard.com/report/github.com/hra42/Go-TextType)

> [!Important]
> Download the latest version from the [release page](https://github.com/HRA42/Go-TextType/releases)
> and run the executable.  
> Visit the [docs](https://go-texttype.postrausch.tech/) for more information.

![Icon](icon.png)

This program is designed to print text from the clipboard using Keyboard events.
Password entry into console applications is enabled by this feature.

## Available HotKeys
- ISO Layout: `left ctrl + left shift + v`
- DE Layout: `links strg + links shift + v`  
or  
- ISO Layout: `left ctrl + left alt + s`
- DE Layout: `links strg + links alt + s`  
or  
- ISO Layout: `left ctrl + Q`
- DE Layout: `links strg + Q`

## Functionality
The main functionality of this program is as described:
The application is designed to print the text stored in the clipboard.
It achieves this by calling a hotkey combination that triggers a function in the program.
This triggered function fetches the content from the clipboard and prints it on the console.

## Running the Program
Download the latest version of the application from the release page.

## Change Log
Check out the [Change Log](https://go-texttype.postrausch.tech/changelog.html)

## Contributing
If you want to contribute or just compile the app yourself,
you can follow the [guide.](https://go-texttype.postrausch.tech/contribute.html)
Before you start to contribute, please read the [contribution file.](./CONTRIBUTING.md)

## Troubleshooting
The program logs the current version of the application, and the Build ID in its log file TextType.log.
You may refer to this file for troubleshooting and referencing specific versions of the application.

## Feedback
If you have any suggestions or questions,
please feel free to open an issue on [GitHub](https://github.com/HRA42/Go-TextType/issues).
