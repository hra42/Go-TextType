# Contribute

## How to get started

This guide will help you contribute to the Go-TextType project. Follow the steps below to get started:

### Get the source code from GitHub
```Bash
git clone https://github.com/HRA42/Go-TextType.git
```
Also, you need to install a C Compiler.
The steps for installing with MSYS2 (recommended) are:
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
- For the compiler to work on other terminals, you will need to set up the windows %PATH% variable to find these tools.  
Go to the “Edit the system environment variables” control panel, tap “Advanced” and add
“C:\msys64\mingw64\bin” to the Path list.

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