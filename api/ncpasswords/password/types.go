package password

type Password struct {
	//TODO: CustomFields
	//TODO: Share

	Id           string `json:"id"`
	Created      int    `json:"created"`
	Updated      int    `json:"updated"`
	Edited       int    `json:"edited"`
	Share        string `json:"share"`
	Shared       bool   `json:"shared"`
	Revision     string `json:"revision"`
	Label        string `json:"label"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Notes        string `json:"notes"`
	CustomFields string `json:"customFields"`
	Url          string `json:"url"`
	Status       int    `json:"status"`
	StatusCode   string `json:"statusCode"`
	Hash         string `json:"hash"`
	Folder       string `json:"folder"`
	CseKey       string `json:"cseKey"`
	CseType      string `json:"cseType"`
	SseType      string `json:"sseType"`
	Hidden       bool   `json:"hidden"`
	Trashed      bool   `json:"trashed"`
	Favorite     bool   `json:"favorite"`
	Editable     bool   `json:"editable"`
	Client       string `json:"client"`
}

//    {
//        "id": "cb56416b-77e2-4edc-b23e-e0d89bb2888c",
//        "created": 1642642689,
//        "updated": 1642642689,
//        "edited": 1614799304,
//        "share": null,
//        "shared": false,
//        "revision": "4971d299-055f-4cdf-b7e1-c0ef6c0288c0",
//        "label": "TrueNAS",
//        "username": "root",
//        "password": "damn#",
//        "notes": "",
//        "customFields": "[]",
//        "url": "http://192.168.1.251",
//        "status": 0,
//        "statusCode": "GOOD",
//        "hash": "d0da89db9cd5e84581f6ae149669a57e024bea4b",
//        "folder": "73db48e7-71d1-429e-b29b-21758daccd0f",
//        "cseKey": "",
//        "cseType": "none",
//        "sseType": "SSEv1r2",
//        "hidden": false,
//        "trashed": false,
//        "favorite": false,
//        "editable": true,
//        "client": "Passwords Session 20.01.22 01:37 - bruno@89.115.179.144"
//    },
