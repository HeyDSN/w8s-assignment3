package models

type Data struct {
	Status Status `json:"status"`
}

type Status struct {
	Wind  int `json:"wind"`
	Water int `json:"water"`
}
