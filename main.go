package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"log"
	"nextcloud-passwords/api/ncpasswords/folder"
	"nextcloud-passwords/api/ncpasswords/password"
	"nextcloud-passwords/common/environment"
	"nextcloud-passwords/screen/navigator"
	"time"
)

func main() {
	environment.LoadEnvironmentFile()
	credentials := environment.GetNextCloudCredentials()

	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme())

	navigatorScreen := navigator.Navigator{
		Folders:   nil,
		Passwords: nil,
		Loading:   true,
	}

	foldersLoaded := false
	passwordsLoaded := false

	go func() {
		log.Println("Loading folders")
		folders, err := folder.All(credentials)

		if err != nil {
			log.Fatal(err)
		}

		log.Println("Folders loaded")
		navigatorScreen.Folders = &folders
		foldersLoaded = true
	}()

	go func() {
		log.Println("Loading passwords")
		passwords, err := password.All(credentials)

		if err != nil {
			log.Fatal(err)
		}

		log.Println("Passwords loaded")
		navigatorScreen.Passwords = &passwords
		passwordsLoaded = true
	}()

	go func() {
		for range time.Tick(time.Millisecond * 100) {
			navigatorScreen.Loading = !(passwordsLoaded && foldersLoaded)

			if navigatorScreen.Loading == false {
				log.Println("Set loading status to False")
				return
			}
		}
	}()

	navigator.Initialize(myApp, &navigatorScreen).ShowAndRun()
}
