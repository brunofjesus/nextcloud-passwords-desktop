package password

import (
	"encoding/json"
	"log"
	"net/http"
	"nextcloud-passwords/api/ncpasswords"
)

func All(credentials ncpasswords.Credentials) ([]Password, error) {

	response, err := ncpasswords.DoRequest(credentials, http.MethodGet, "password/list", nil)

	if err != nil {
		return nil, err
	}

	var passwords []Password
	err = json.Unmarshal(response, &passwords)

	if err != nil {
		log.Fatalln(err)
	}

	return passwords, nil
}
