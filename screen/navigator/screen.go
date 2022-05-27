package navigator

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"nextcloud-passwords/api/ncpasswords/folder"
	"nextcloud-passwords/api/ncpasswords/password"
)

type Navigator struct {
	Folders   *[]folder.Folder
	Passwords *[]password.Password
}

const rootDir = "00000000-0000-0000-0000-000000000000"

func Initialize(app fyne.App, navigator *Navigator) fyne.Window {
	window := app.NewWindow("Nextcloud Passwords")

	window.Resize(fyne.Size{
		Width:  640,
		Height: 480,
	})

	SwitchFolder(&window, rootDir, navigator)

	return window
}

func SwitchFolder(window *fyne.Window, folderId string, navigator *Navigator) {
	list := createList(window, folderId, navigator)
	breadcrumb := createBreadcrumb(
		folderId,
		navigator.Folders,
		func(btnFolderId string) {
			SwitchFolder(window, btnFolderId, navigator)
		},
	)

	vbox := container.New(layout.NewBorderLayout(breadcrumb, nil, nil, nil), breadcrumb, list)

	(*window).SetContent(vbox)
}

func createList(window *fyne.Window, folderId string, navigator *Navigator) *widget.List {
	folders, passwords := filter(folderId, navigator.Folders, navigator.Passwords)

	list := widget.NewList(
		func() int {
			return len(*passwords) + len(*folders)
		},
		func() fyne.CanvasObject {
			return widget.NewButtonWithIcon("", nil, nil)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			button := o.(*widget.Button)
			button.Alignment = widget.ButtonAlignLeading
			button.Importance = widget.LowImportance

			endFolders := len(*folders)

			if i < endFolders {
				currentFolder := (*folders)[i]

				button.SetText(currentFolder.Label)
				button.SetIcon(theme.FolderIcon())
				button.OnTapped = func() {
					SwitchFolder(window, currentFolder.Id, navigator)
				}
			} else {
				passwordIdx := i - endFolders

				currentPassword := (*passwords)[passwordIdx]

				button.SetText(currentPassword.Label)
				button.SetIcon(theme.FileIcon())
			}
		})

	return list
}

func filter(folderId string, folders *[]folder.Folder, passwords *[]password.Password) (*[]folder.Folder, *[]password.Password) {

	var filteredFolders []folder.Folder

	for _, currentFolder := range *folders {
		if currentFolder.Parent == folderId {
			filteredFolders = append(filteredFolders, currentFolder)
		}
	}

	var filteredPasswords []password.Password

	for _, currentPassword := range *passwords {
		if currentPassword.Folder == folderId {
			filteredPasswords = append(filteredPasswords, currentPassword)
		}
	}

	return &filteredFolders, &filteredPasswords
}
