param (
    [Parameter(Position=0, Mandatory=$true)]
    [string]$action
)

$Version = "1.1.3"
$BuildID = Get-Date -Format "Hmmss"

function build {
    clean
    Write-Host "Starting build..."
    go mod tidy
    Write-Host "Generating binary icon"
    rsrc -arch amd64 -ico .\internal\trayMenu\icon.ico
    Write-Host "Generating binary version info"
    goversioninfo -64 -product-ver-build $BuildID -ver-build $BuildID
    Write-Host "Building project"
    go build -ldflags "-extldflags '-static' -w -s -H windowsgui -X main.AppVersion=$Version -X main.BuildID=$BUILDID" `
        -o .\bin\TextType.exe .
    Write-Host "Build completed."
}

function clean {
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

function dep {
    Write-Host "Installing dependencies"
    go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo
    go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo
    go get github.com/akavel/rsrc
    go install github.com/akavel/rsrc
    go get .
    go mod tidy
}

switch ($action) {
    'build' {
        build
    }
    'clean' {
        clean
    }
    'dep' {
        dep
    }
    default {
        Write-Host "Invalid action parameter. Use 'build' or 'clean' or 'dep'."
    }
}