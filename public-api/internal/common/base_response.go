package common

type ErrResponseDto struct {
	Results bool   `json:"result"`
	Errors  string `json:"errors"`
}
