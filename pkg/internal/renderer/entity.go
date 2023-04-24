package renderer

type jsonRenderer struct{}

type yamlRenderer struct{}

type csvRenderer struct {
	settings *CsvSettings
}

type CsvSettings struct {
	Price         float32
	TimeFormat    string
	TimeZone      string
	Comma         string
	HeaderEnabled bool
	FilePath      string
}

type pdfRenderer struct {
	settings *PdfSettings
}

type PdfSettings struct {
	Price       float32
	TimeFormat  string
	TimeZone    string
	Comma       string
	PrintHeader bool
	LogoHeader  string
	Settings    GlobalSettings
}

type GlobalSettings struct {
	Firstname string
	Lastname  string
	Street    string
	Postcode  string
	City      string
}
