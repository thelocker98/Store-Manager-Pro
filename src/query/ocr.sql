-- name: CreateOCREntry :one
INSERT INTO ocr_files (invoice_id, filepath, status)
	VALUES ($1, $2, $3)
	RETURNING file_id;

-- name: GetOCREntryById :one
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
WHERE file_id = $1;

-- name: GetOCREntrysByInvoiceID :many
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
WHERE invoice_id = $1
ORDER BY created_at DESC;

-- name: GetUnfinishedEntrys :many
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
WHERE status IN ($1, $2);

-- name: DeleteOCREntryById :execrows
DELETE FROM ocr_files
WHERE file_id = $1 AND (sqlc.arg(force)::boolean = true OR status != 1);

-- name: MarkAsProccessing :execrows
UPDATE ocr_files
	SET status = $1, started_at = CURRENT_TIMESTAMP
	WHERE file_id = $2;

-- name: MarkAsFinished :execrows
UPDATE ocr_files
	SET status = $1, entries_added = $2, entries_failed = $3, error_message = $4, completed_at = CURRENT_TIMESTAMP
	WHERE file_id = $5;
