package folder

import (
	"encoding/json"
	"log"
	"net/http"
	"nextcloud-passwords/api/ncpasswords"
)

func All(credentials ncpasswords.Credentials) ([]Folder, error) {

	response, err := ncpasswords.DoRequest(credentials, http.MethodGet, "folder/list", nil)
	if err != nil {
		return nil, err
	}

	var folders []Folder
	err = json.Unmarshal(response, &folders)

	if err != nil {
		log.Fatalln(err)
	}

	return folders, nil
}
