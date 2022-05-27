package navigator

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"nextcloud-passwords/api/ncpasswords/folder"
)

type FnStringConsumer = func(string)

func createBreadcrumb(folderId string, folders *[]folder.Folder, changeFolderFn FnStringConsumer) *fyne.Container {

	breadcrumb := container.New(layout.NewHBoxLayout())

	pathFolders := findFolderPath(folderId, folders)
	addFolderButtons(&pathFolders, breadcrumb, changeFolderFn)

	breadcrumb.Add(layout.NewSpacer())

	return breadcrumb
}

func findFolderPath(folderId string, folders *[]folder.Folder) []*folder.Folder {
	var breadcrumbFolders []*folder.Folder

	isOnRoot := folderId == rootDir

	if isOnRoot == false {
		// Add current folder
		currentFolder := findFolderById(folderId, folders)
		breadcrumbFolders = append(breadcrumbFolders, currentFolder)

		// Add parent folder chain
		for isOnRoot == false {
			currentFolder = findParentFolder(currentFolder, folders)

			isOnRoot = currentFolder == nil || currentFolder.Id == rootDir

			if isOnRoot == false {
				breadcrumbFolders = append(breadcrumbFolders, currentFolder)
			}
		}
	}

	return breadcrumbFolders
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

func addFolderButtons(pathFolders *[]*folder.Folder, container *fyne.Container, changeFolderFn FnStringConsumer) {

	container.Add(widget.NewButtonWithIcon("", theme.HomeIcon(), func() {
		changeFolderFn(rootDir)
	}))

	for i := len(*pathFolders) - 1; i >= 0; i-- {
		container.Add(createFolderButton((*pathFolders)[i], changeFolderFn))
	}
}

func createFolderButton(folder *folder.Folder, changeFolderFn FnStringConsumer) *widget.Button {
	return widget.NewButton(folder.Label, func() {
		changeFolderFn(folder.Id)
	})
}
