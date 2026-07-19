package ocr

import (
	"bufio"
	"image/png"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"gitea.locker98.com/locker98/Store-Manager-Pro/models"
	"github.com/gen2brain/go-fitz"
	gosseract "github.com/otiai10/gosseract/v2"
)

var jobFolder = "job"

func PDFtoImage(pdfFilePath string, filePathImage string) string {

	err := os.Mkdir(filePathImage+"/"+jobFolder, 0755)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	doc, err := fitz.New(pdfFilePath)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer doc.Close()

	for i := 0; i < doc.NumPage(); i++ {

		img, err := doc.Image(i)
		if err != nil {
			log.Fatal(err)
		}

		fileName := filePathImage + "/" + jobFolder + "/page_" + strconv.Itoa(i) + ".png"
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		png.Encode(file, img)
	}

	return filePathImage + "/" + jobFolder
}

func ImageToText(imageFolderPath string) string {
	// 1. Create the Tesseract client
	client := gosseract.NewClient()
	defer client.Close()

	// 2. Open (or Create) the output text file
	// We use filepath.Join to make sure it works on Windows/Linux/Mac
	outputPath := filepath.Join(imageFolderPath, "..", "text.txt")

	file, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 3. Read the directory contents
	entries, err := os.ReadDir(imageFolderPath)
	if err != nil {
		log.Fatal(err)
	}

	// 4. Loop through files in the folder
	for _, entry := range entries {
		// Skip directories, we only want images
		if entry.IsDir() || filepath.Ext(entry.Name()) == ".txt" {
			continue
		}

		// Get the full path to the image
		imagePath := filepath.Join(imageFolderPath, entry.Name())

		// Run OCR
		client.SetImage(imagePath)
		text, err := client.Text()
		if err != nil {
			log.Printf("Error reading %s: %v", entry.Name(), err)
			continue // Skip this file if OCR fails
		}

		// 5. Write text to the file
		// We add a newline after each page so they don't run together
		if _, err := file.WriteString(text + "\n"); err != nil {
			log.Fatal(err)
		}
	}

	// Remove Images
	err = os.RemoveAll(imageFolderPath)
	if err != nil {
		return ""
	}

	// Return the path to the created file
	return outputPath
}

func CheckIfFail(row string) bool {
	triggers := []string{
		"SIZE",
		"UPC",
		"PRODUCT",
		"WAREHOUSE DAMAGE",
		"TOTAL",
		"DISCCOST",
		"Batch",
		"Department",
		"AMOUNT",
		"DUE",
		"Discount",
	}

	if strings.TrimSpace(row) == "" {
		return false
	}

	for _, trigger := range triggers {
		if strings.Contains(row, trigger) {
			return false
		}
	}
	return true
}

func TextToDB(textFilePath string, invoice_id int, markup int) ([]models.InvoiceEntry, int, int) {
	var success, failed int
	var entrys []models.InvoiceEntry

	file, err := os.Open(textFilePath)
	if err != nil {
		return []models.InvoiceEntry{}, 0, 0
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	// Regex
	specialChar := regexp.MustCompile(`[^A-Za-z0-9,.$'()&]+`)

	// Loop through the file line by line
	for scanner.Scan() {
		var entry models.InvoiceEntry

		row := scanner.Text()
		row = specialChar.ReplaceAllString(row, " ")
		row = strings.ReplaceAll(row, "§", "$")
		row = strings.ReplaceAll(row, ",", ".")
		row = regexp.MustCompile(`\s+`).ReplaceAllString(row, " ")
		row = strings.TrimSpace(row)

		// Find UPC
		upcRegex := regexp.MustCompile(`\d{8,12}`)
		upc := upcRegex.FindString(row)
		if upc == "" {
			if CheckIfFail(row) {
				failed++
			}
			continue
		}

		// Find all prices
		priceRegex := regexp.MustCompile(`\d+\.\d{2}`)
		spaceFreeRow := strings.ReplaceAll(row, " ", "")
		prices := priceRegex.FindAllString(spaceFreeRow, -1)

		if len(prices) < 3 {
			if CheckIfFail(row) {
				failed++
			}
			continue
		}

		cost, err1 := strconv.ParseFloat(prices[len(prices)-3], 64)
		totalCost, err2 := strconv.ParseFloat(prices[len(prices)-2], 64)
		discCost, err3 := strconv.ParseFloat(prices[len(prices)-1], 64)

		if err1 != nil || err2 != nil || err3 != nil {
			if CheckIfFail(row) {
				failed++
			}
			continue
		}

		// Remove UPC
		detail := strings.TrimPrefix(row, upc)

		// Remove last 3 prices from detail
		for i := len(prices) - 3; i < len(prices); i++ {
			detail = strings.Replace(detail, prices[i], "", 1)
		}

		detail = strings.ReplaceAll(detail, "$", "")
		detail = regexp.MustCompile(`\s+`).ReplaceAllString(detail, " ")
		detail = strings.TrimSpace(detail)

		entry.RawText = row
		entry.UPC = upc
		entry.Detail = detail
		entry.ItemCost = cost
		entry.TotalCost = totalCost
		entry.DiscountCost = discCost
		entry.InvoiceID = invoice_id

		// Check that Costs are not zero
		if cost <= 0 || discCost <= 0 {
			failed++
			continue
		}
		entry.QTY = int(totalCost/cost + 0.5)

		// Do not include rows with zero quantity
		if entry.QTY <= 0 {
			failed++
			continue
		}

		entry.TrueCost = math.Round((discCost/float64(entry.QTY))*(1+float64(markup)/100.0)*100) / 100

		// Add entry to list
		entrys = append(entrys, entry)
		success++
	}

	// Remove Text File
	os.RemoveAll(textFilePath)

	// Return the path to the created file
	return entrys, failed, success
}
