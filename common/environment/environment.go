package environment

import (
	"github.com/joho/godotenv"
	"log"
	"nextcloud-passwords/api/ncpasswords"
	"os"
)

func LoadEnvironmentFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func GetNextCloudCredentials() ncpasswords.Credentials {
	return ncpasswords.Credentials{
		Username: os.Getenv("nc_username"),
		Password: os.Getenv("nc_passwd"),
		Url:      os.Getenv("nc_url"),
	}
}
