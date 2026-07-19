package utils

import (
	"os"
	"path/filepath"
)

var DBFolderPath string
var InvoiceFolderPath string

func SetupFolders() (string, string) {
	// Create App Directory
	appData, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	DBFolderPath = filepath.Join(appData, "StoreManagerPro")
	err = os.MkdirAll(DBFolderPath, 0755)
	if err != nil {
		panic(err)
	}

	// Setup Invoice Upload Folder
	InvoiceFolderPath = filepath.Join(DBFolderPath, "InvoiceFolder")
	err = os.RemoveAll(InvoiceFolderPath)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(InvoiceFolderPath, 0755)
	if err != nil {
		panic(err)
	}

	return DBFolderPath, InvoiceFolderPath
}
