package ocr

import (
	"fmt"
	"os"
	"strings"

	"gitea.locker98.com/locker98/Store-Manager-Pro/db"
	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

var OCRQueue = make(chan models.OCRJob, 100)

func StartOCRWorker() {
	for job := range OCRQueue {
		processOCRFile(job)
	}
}

func processOCRFile(job models.OCRJob) {
	fmt.Println("--Processing OCR Job Started--")

	// Check File was not canceled
	entry, err := db.GetOCREntryById(job.FileID)
	if err != nil || entry.FileID == 0 {
		if job.FilePath != "" {
			os.Remove(job.FilePath)
		}
		return
	}

	// Mark job as started
	err = db.MarkAsProccessing(job.FileID)
	if err != nil {
		fmt.Println("\n\nError: " + err.Error() + "\n")
		return
	}

	fmt.Println("File ID: ", job.FileID)
	fmt.Println("Invoice ID: ", job.InvoiceID)
	fmt.Println("Markup Amount: ", job.MarkupAmount)
	fmt.Println("File Path: ", job.FilePath)

	// Get the last part after the last underscore
	parts := strings.Split(job.FilePath, "/")
	path := strings.Join(parts[:len(parts)-1], "/")
	imageFolderPath := PDFtoImage(job.FilePath, path)
	textFilePath := ImageToText(imageFolderPath)
	entrys, failed, success := TextToDB(textFilePath, job.InvoiceID, job.MarkupAmount)

	for _, entry := range entrys {
		db.AddEntryToInvoice(entry)

	}

	// Mark Job and finished and remove files
	err = db.MarkAsFinished(job.FileID, models.Succeeded, success, failed, "Success")
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	os.Remove(job.FilePath)
	fmt.Println("--Processing OCR Job Finished--")
}
