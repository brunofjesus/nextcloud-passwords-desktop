package folder

type Folder struct {
	Id       string `json:"id"`
	Created  int    `json:"created"`
	Updated  int    `json:"updated"`
	Edited   int    `json:"edited"`
	Revision string `json:"revision"`
	Label    string `json:"label"`
	Parent   string `json:"parent"`
	CseKey   string `json:"cseKey"`
	CseType  string `json:"cseType"`
	SseType  string `json:"sseType"`
	Hidden   bool   `json:"hidden"`
	Trashed  bool   `json:"trashed"`
	Favorite bool   `json:"favorite"`
	Client   string `json:"client"`
}

// {
//        "id": "00aac833-17c1-4c0e-9d00-f2b4f9291765",
//        "created": 1642642689,
//        "updated": 1642642689,
//        "edited": 1616240962,
//        "revision": "9cad2f61-11d4-45f5-977f-dcff30f1813c",
//        "label": "Applications",
//        "parent": "00000000-0000-0000-0000-000000000000",
//        "cseKey": "",
//        "cseType": "none",
//        "sseType": "SSEv1r2",
//        "hidden": false,
//        "trashed": false,
//        "favorite": false,
//        "client": "Passwords Session 20.01.22 01:37 - bruno@89.115.179.144"
//    },
