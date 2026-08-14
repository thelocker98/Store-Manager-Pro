package db

import (
	"strings"

	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
)

func CreateOCREntry(invoiceID int, filePath string) (int, error) {
	query := `
	INSERT INTO ocr_files (invoice_id, filepath, status )
	VALUES ($1, $2, $3)
	RETURNING file_id;
	`

	var fileID int
	err := DB.QueryRow(query,
		invoiceID,
		filePath,
		models.Pending,
	).Scan(&fileID)

	if err != nil {
		return 0, err
	}

	return fileID, nil
}

func GetOCREntryById(file_id int) (models.OCREntry, error) {
	query := `
	SELECT
		file_id,
		invoice_id,
		filepath,
		status,
		entries_added,
		entries_failed,
		error_message,
		started_at,
		completed_at,
		created_at
	FROM ocr_files
	WHERE file_id = $1
`

	var entry models.OCREntry
	err := DB.QueryRow(query, file_id).Scan(
		&entry.FileID,
		&entry.InvoiceID,
		&entry.FilePath,
		&entry.Status,
		&entry.EntriesAdded,
		&entry.EntriesFailed,
		&entry.ErrorMessage,
		&entry.StartedAt,
		&entry.CompletedAt,
		&entry.CreatedAt,
	)

	if err != nil {
		return entry, err
	}

	return entry, nil
}

func GetOCREntrysByInvoiceID(invoiceID int) ([]models.OCREntry, error) {
	query := `
		SELECT
			file_id,
			invoice_id,
			filepath,
			status,
			entries_added,
			entries_failed,
			error_message,
			started_at,
			completed_at,
			created_at
		FROM ocr_files
		WHERE
			invoice_id = $1
		ORDER BY created_at DESC;
		`

	rows, err := DB.Query(query, invoiceID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entrys []models.OCREntry
	for rows.Next() {
		var entry models.OCREntry
		err := rows.Scan(
			&entry.FileID,
			&entry.InvoiceID,
			&entry.FilePath,
			&entry.Status,
			&entry.EntriesAdded,
			&entry.EntriesFailed,
			&entry.ErrorMessage,
			&entry.StartedAt,
			&entry.CompletedAt,
			&entry.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Get the last part after the last underscore
		parts := strings.Split(entry.FilePath, "_")
		entry.FilePath = parts[len(parts)-1]
		entrys = append(entrys, entry)
	}
	return entrys, nil
}

func DeleteOCREntryById(file_id int, force bool) error {
	var err error
	if force {
		_, err = DB.Exec(`DELETE FROM ocr_files WHERE file_id = $1;`, file_id)
	} else {
		_, err = DB.Exec(`DELETE FROM ocr_files WHERE file_id = $1 and status != 1;`, file_id)
	}
	return err
}

func GetUnfinishedEntrys() ([]models.OCREntry, error) {
	query := `
		SELECT
			file_id,
			invoice_id,
			filepath,
			status,
			entries_added,
			entries_failed,
			error_message,
			started_at,
			completed_at,
			created_at
		FROM ocr_files
		WHERE
			status IN ($1, $2);
		`

	rows, err := DB.Query(query, models.Pending, models.Processing)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entrys []models.OCREntry
	for rows.Next() {
		var entry models.OCREntry
		err := rows.Scan(
			&entry.FileID,
			&entry.InvoiceID,
			&entry.FilePath,
			&entry.Status,
			&entry.EntriesAdded,
			&entry.EntriesFailed,
			&entry.ErrorMessage,
			&entry.StartedAt,
			&entry.CompletedAt,
			&entry.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		entrys = append(entrys, entry)
	}
	return entrys, nil
}

func MarkAsProccessing(fileId int) error {
	query := `
	UPDATE ocr_files
	SET status = $1, started_at = CURRENT_TIMESTAMP
	WHERE file_id = $2;
	`
	_, err := DB.Exec(query, models.Processing, fileId)
	return err
}

func MarkAsFinished(fileId int, status models.OCRStatus, successfulEntrys int, failedEntrys int, errorMessage string) error {
	query := `
	UPDATE ocr_files
	SET status = $1, entries_added = $2, entries_failed = $3, error_message = $4, completed_at = CURRENT_TIMESTAMP
	WHERE file_id = $5;
	`
	_, err := DB.Exec(query, status, successfulEntrys, failedEntrys, errorMessage, fileId)
	return err
}
