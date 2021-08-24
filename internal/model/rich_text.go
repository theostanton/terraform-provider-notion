package model

type RichText struct {
	Type        string       `json:"type"`
	Text        Text         `json:"text"`
	Annotations *Annotations `json:"annotations,omitempty"`
	PlainText   *string      `json:"plain_text,omitempty"`
	Href        interface{}  `json:"href,omitempty"`
}

func (richText *RichText) Get() string {
	if richText.PlainText != nil {
		return *richText.PlainText
	}
	return richText.Text.Content
}

type Text struct {
	Content string      `json:"content"`
	Link    interface{} `json:"link,omitempty"`
}

type Annotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

func NewAnnotations() Annotations {
	return Annotations{
		Bold:          false,
		Italic:        false,
		Strikethrough: false,
		Underline:     false,
		Code:          false,
		Color:         "default",
	}
}

func NewRichText(text string) RichText {
	return RichText{
		Type: "text",
		Text: Text{
			Content: text,
			Link:    nil,
		},
	}
}
