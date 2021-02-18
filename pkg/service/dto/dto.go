package dto

type Entry struct {
	Id     int    `json:"entry_id,omitempty"`
	Name   string `json:"name"`
	Number string `json:"phone_number"`
	Note   string `json:"note"`
}
