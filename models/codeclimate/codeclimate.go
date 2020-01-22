package codeclimate

import "encoding/json"

func (c *CodeClimate) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

type (
	Collection  []CodeClimate
	CodeClimate struct {
		Description string `json:"description"`
		Fingerprint string `json:"fingerprint"`
		Location    struct {
			Path  string `json:"path"`
			Lines struct {
				Begin int `json:"begin"`
				End   int `json:"end"`
			} `json:"lines"`
		} `json:"location"`
	}
)