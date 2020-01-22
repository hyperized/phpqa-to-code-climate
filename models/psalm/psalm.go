package psalm

import "encoding/json"

func (c *Collection) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &c)
}

type (
	Collection []Psalm
	Psalm      struct {
		Severity     string `json:"severity"`
		LineFrom     int    `json:"line_from"`
		LineTo       int    `json:"line_to"`
		Type         string `json:"type"`
		Message      string `json:"message"`
		FileName     string `json:"file_name"`
		FilePath     string `json:"file_path"`
		Snippet      string `json:"snippet"`
		SelectedText string `json:"selected_text"`
		From         int    `json:"from"`
		To           int    `json:"to"`
		SnippetFrom  int    `json:"snippet_from"`
		SnippetTo    int    `json:"snippet_to"`
		ColumnFrom   int    `json:"column_from"`
		ColumnTo     int    `json:"column_to"`
	}
)
