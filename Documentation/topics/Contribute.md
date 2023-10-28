<!-- set toc level to how -->
<show-structure for="chapter,procedure" depth="2"></show-structure>

# How to contribute to the project?
> The build target for this project is Microsoft Windows 10/11.  
> Instructions are only provided for Windows users.
{style="warning"}

You should start reading the [Contributor Guide](https://github.com/HRA42/Go-TextType/blob/main/CONTRIBUTING.md),
the [Code of Conduct](https://github.com/HRA42/Go-TextType/blob/main/CODE_OF_CONDUCT.md) and
the [Security Policy.](https://github.com/HRA42/Go-TextType/blob/main/SECURITY.md)

## Project Dependencies

This project primarily utilizes the Go programming language.
Presented are the essential Go packages, along with their purposes:

- **[embed](https://golang.org/pkg/embed/)**: Useful for embedding files within the executable.

- **[encoding/gob](https://golang.org/pkg/encoding/gob/)**:
Provides functions for encoding hotkey preferences and other data to a file.

- **[github.com/gen2brain/beeep](https://github.com/gen2brain/beeep)**: Enables update notifications.

- **[github.com/getlantern/systray](https://github.com/getlantern/systray)**: Manages system tray functionality.

- **[github.com/go-vgo/robotgo](https://github.com/go-vgo/robotgo)**: Simulates keyboard inputs.

- **[github.com/pkg/browser](https://github.com/pkg/browser)**:
Facilitates opening of a browser directed to the release page on GitHub.

- **[github.com/tcnksm/go-latest](https://github.com/tcnksm/go-latest)**:
Helps in retrieving the latest version from GitHub.

- **[golang.design/x/clipboard](https://github.com/golang-design/clipboard)**: Grants access to the system clipboard.

- **[golang.design/x/hotkey](https://github.com/golang-design/hotkey)**: Take care of creating and managing hotkeys.

## How to get started

This guide will help you contribute to the Go-TextType project. Follow the steps below to get started:

### System setup
1. Set up the [Go SDK](https://go.dev/doc/install).
2. Install a C-Compiler:
   - The steps for installing with MSYS2 (recommended) are:
     - Download and install the latest version of [MSYS2](https://www.msys2.org/#download).
     - Once installed, do not use the MSYS terminal that opens
     - Open “MSYS2 MinGW64” from the start menu
     - Execute the following commands (if asked for installation options, be sure to choose “all”):
    ```Bash
    pacman -Syu
    pacman -S git mingw-w64-x86_64-toolchain
    ```
    - You will need to add /c/Program\ Files/Go/bin and ~/Go/bin to your $PATH,
    for MSYS2 you can paste the following command into your terminal:
    ```Bash
    echo "export PATH=\$PATH:/c/Program\ Files/Go/bin:~/Go/bin" >> ~/.bashrc`
    ```
    - For the compiler to work on other terminals, you will need to set up the windows `PATH` variable to find these tools.  
    Go to the “Edit the system environment variables” control panel, tap “Advanced” and add
    “C:\msys64\mingw64\bin” to the `Path list`.

### Get the source code from GitHub

```Bash
git clone https://github.com/HRA42/Go-TextType.git
```

### Download the dependencies

```Shell
```
{src="make.ps1" include-lines="36-42"}

Or run:
```Shell
Powershell.exe -File .\build.ps1 dep
```

### To build the application

```Shell
```
{src="make.ps1" include-lines="11-20"}

Or run:
```Bash
Powershell.exe -File .\build.ps1 build
```

### To clean the application folder

```Shell
```
{src="make.ps1" include-lines="24-32"}

Or run:
```Bash
Powershell.exe -File .\build.ps1 clean
```