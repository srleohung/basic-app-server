package types

import "time"

type ExampleData struct {
	Vendor    string    `json:"vendor"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Model     string    `json:"model"`
	Address   string    `json:"address"`
	Country   string    `json:"country"`
	City      string    `json:"city"`
	District  string    `json:"district"`
	Timestamp time.Time `json:"timestamp"`
}

func (exampleData *ExampleData) Format() ExampleData {
	return *exampleData
}

func (exampleData *ExampleData) Init() ExampleData {
	return ExampleData{
		Vendor:    "example_vendor",
		Code:      "example_code",
		Name:      "example_name",
		Model:     "example_model",
		Address:   "example_address",
		Country:   "example_country",
		City:      "example_city",
		District:  "example_district",
		Timestamp: time.Now(),
	}
}
