package update

import (
	"github.com/gen2brain/beeep"
	"github.com/hra42/Go-TextType/internal/errorHandling"
	"github.com/hra42/Go-TextType/internal/logging"
	"github.com/hra42/Go-TextType/internal/trayMenu"
	"github.com/pkg/browser"
	"github.com/tcnksm/go-latest"
	"os"
)

func CheckUpdate(AppVersion string) {
	dumpfile, err := os.CreateTemp("", "icon.*.png")
	errorHandling.CheckError(err)
	defer deleteFile(dumpfile)
	if _, err = dumpfile.Write(trayMenu.ReadIcon()); err != nil {
		errorHandling.CheckError(err)
	}
	if err = dumpfile.Close(); err != nil {
		errorHandling.CheckError(err)
	}
	githubTag := &latest.GithubTag{
		Owner:      "HRA42",
		Repository: "Go-TextType",
	}
	res, err := latest.Check(githubTag, AppVersion)
	errorHandling.CheckError(err)
	if res.Outdated {
		err = beeep.Alert(
			"Update Available!",
			"A new version of Go-TextType is available!",
			dumpfile.Name(),
		)
		errorHandling.CheckError(err)
		logging.Logger.Println("Current Version is:", AppVersion)
		logging.Logger.Println("Update available! Latest Version is:", res.Current)
		err = browser.OpenURL("https://github.com/HRA42/Go-TextType/releases/latest")
		errorHandling.CheckError(err)
	}
	logging.Logger.Println("No update available")
}

func deleteFile(file *os.File) {
	err := os.Remove(file.Name())
	errorHandling.CheckError(err)
}
