package navigator

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"nextcloud-passwords/api/ncpasswords/folder"
)

func createBreadcrumb(folderId string, folders *[]folder.Folder) *fyne.Container {

	btnHome := widget.NewButtonWithIcon("", theme.HomeIcon(), func() {
		fmt.Println("Tapped")
	})
	breadcrumb := container.New(layout.NewHBoxLayout(), btnHome)

	findAndAddFolders(folderId, folders, breadcrumb)

	breadcrumb.Add(layout.NewSpacer())

	return breadcrumb
}

func findAndAddFolders(folderId string, folders *[]folder.Folder, container *fyne.Container) {
	var buttons []*widget.Button

	isOnRoot := folderId == rootDir

	if isOnRoot == false {
		currentFolder := findFolderById(folderId, folders)
		buttons = append(buttons, createFolderButton(currentFolder))

		for isOnRoot == false {
			currentFolder = findParentFolder(currentFolder, folders)

			isOnRoot = currentFolder == nil || currentFolder.Id == rootDir

			if isOnRoot == false {
				buttons = append(buttons, createFolderButton(currentFolder))
			}
		}

		for i := len(buttons) - 1; i >= 0; i-- {
			container.Add(buttons[i])
		}
	}
}

func findParentFolder(folder *folder.Folder, folders *[]folder.Folder) *folder.Folder {
	return findFolderById(folder.Parent, folders)
}

func findFolderById(folderId string, folders *[]folder.Folder) *folder.Folder {
	for _, currentFolder := range *folders {
		if currentFolder.Id == folderId {
			return &currentFolder
		}
	}

	return nil
}

func createFolderButton(folder *folder.Folder) *widget.Button {
	return widget.NewButton(folder.Label, nil)
}
