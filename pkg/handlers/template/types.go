package template

type CityValues struct {
	City              string
	Latitude          string
	Longitude         string
	ConditionClimatic ConditionClimatic
	Pressure          string
	Humidity          string
	Cloudiness        string
	Temp              string
	TempMax           string
	TempMin           string
}

type ConditionClimatic struct {
	Description string
}
