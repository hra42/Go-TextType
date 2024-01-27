param (
    [Parameter(Position=0, Mandatory=$true)]
    [string]$action
)

$Version = "1.1.6"
$BuildID = Get-Date -Format "Hmmss"

function Clean {
    Write-Host "Cleaning project"
    go clean
    Write-Host "Cleaning binary folder"
    if (Test-Path .\bin) {
        Remove-Item -Force -Recurse .\bin
    } else {
        Write-Host "No binary folder to clean."
    }
    Write-Host "Cleaning resources folder completed."
}

function Dep {
    Write-Host "Installing dependencies"
    go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo
    go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo
    go get github.com/akavel/rsrc
    go install github.com/akavel/rsrc
    go get .
    go mod tidy
}

function Build {
    & Clean
    Write-Host "Starting build..."
    go mod tidy
    Write-Host "Generating binary icon"
    rsrc -arch amd64 -ico .\internal\trayMenu\icon.ico
    Write-Host "Generating binary version info"
    goversioninfo -64 -product-ver-build $BuildID -ver-build $BuildID
    Write-Host "Building project"
    go build -ldflags `
        "-extldflags '-static' -w -s -H windowsgui -X main.AppVersion=$Version -X main.BuildID=$BUILDID" `
        -o .\bin\TextType.exe .
    Write-Host "Build completed."
    Write-Host "Compressing the binary"
    upx --brute .\bin\TextType.exe
}

switch ($action) {
    'build' {
        Build
    }
    'clean' {
        Clean
    }
    'dep' {
        Dep
    }
    default {
        Write-Host "Invalid action parameter. Use 'build' or 'clean' or 'dep'."
    }
}