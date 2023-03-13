package renderer

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

func (cR *csvRenderer) Render(filePath string, render func(writer *csv.Writer, timeZone *time.Location, timeFormat string, price float32)) error {
	//render csv file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()
	csvWriter := csv.NewWriter(file)
	csvWriter.Comma = []rune(cR.settings.Comma)[0]

	if cR.settings.HeaderEnabled {
		header := []string{"Time", "User", "Power meter start", "Power meter end", "Charge (kWh)", "Duration (hh:mm:ss)", "Cost (€)"}
		csvWriter.Write(header)
	}

	price, err := strconv.ParseFloat(cR.settings.Price, 32)
	if err != nil {
		return err
	}

	location, _ := time.LoadLocation(cR.settings.TimeZone)
	timeFormat := cR.settings.TimeFormat
	render(csvWriter, location, timeFormat, float32(price))

	csvWriter.Flush()

	return nil
}
