package main

import (
	"GIG/app/models"
	"GIG/app/utility/entityhandlers"
	"GIG/app/utility/importers/etender/decoders"
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var category = "Tenders"

func main() {

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("file path not specified")
		os.Exit(1)
	}
	filePath := args[0]

	csvFile, _ := os.Open(filePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	ignoreHeaders := true

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if ignoreHeaders {
			ignoreHeaders = false
		} else {
			tender := decoders.Decode(line)
			companyEntity := models.Entity{
				Title: tender.Company,
			}.AddCategory("Organization")

			locationEntity := models.Entity{
				Title: tender.Location,
			}.AddCategory("Location")

			entity := decoders.MapToEntity(tender).AddCategory(category)

			entity, addCompanyError := entityhandlers.AddEntityAsAttribute(entity, "Company", companyEntity)
			if addCompanyError!=nil{
				fmt.Println(addCompanyError)
			}
			entity, addLocationError := entityhandlers.AddEntityAsAttribute(entity, "Location", locationEntity)
			if addLocationError!=nil{
				fmt.Println(addLocationError)
			}

			savedEntity, saveErr := entityhandlers.CreateEntity(entity)

			if saveErr != nil {
				fmt.Println(saveErr.Error(), entity)
			}
			fmt.Println(savedEntity.Title)
		}
	}
}
