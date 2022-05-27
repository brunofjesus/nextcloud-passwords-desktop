package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"log"
	"nextcloud-passwords/api/ncpasswords/folder"
	"nextcloud-passwords/api/ncpasswords/password"
	"nextcloud-passwords/common/environment"
	"nextcloud-passwords/screen/navigator"
)

func main() {
	environment.LoadEnvironmentFile()
	credentials := environment.GetNextCloudCredentials()

	folders, err := folder.All(credentials)

	if err != nil {
		log.Fatal(err)
	}

	passwords, err := password.All(credentials)

	if err != nil {
		log.Fatal(err)
	}

	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme())

	navigatorScreen := navigator.Navigator{
		Folders:   &folders,
		Passwords: &passwords,
	}

	navigator.Initialize(myApp, &navigatorScreen).ShowAndRun()
}
