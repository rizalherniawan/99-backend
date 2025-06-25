package common

type ErrorBaseResponseDto struct {
	Results bool   `json:"result"`
	Errors  string `json:"errors"`
}
