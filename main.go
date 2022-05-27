package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/joho/godotenv"
	"log"
	"nextcloud-passwords/api/ncpasswords"
	"nextcloud-passwords/api/ncpasswords/folder"
	"nextcloud-passwords/api/ncpasswords/password"
	"nextcloud-passwords/screen/navigator"
	"os"
)

func main() {
	loadEnvFile()

	credentials := ncpasswords.Credentials{
		Username: os.Getenv("nc_username"),
		Password: os.Getenv("nc_passwd"),
		Url:      os.Getenv("nc_url"),
	}

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

	navigator.Initialize(myApp, navigator.Navigator{Folders: &folders, Passwords: &passwords}).ShowAndRun()

}

func loadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}
