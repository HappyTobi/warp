package renderer

type jsonRenderer struct{}

type yamlRenderer struct{}

type csvRenderer struct {
	settings *CsvSettings
}

type CsvSettings struct {
	Price         string
	TimeFormat    string
	TimeZone      string
	Comma         string
	HeaderEnabled bool
	FilePath      string
}
