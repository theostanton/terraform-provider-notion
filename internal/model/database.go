package model

import "time"

type Database struct {
	Object         string                      `json:"object,omitempty"`
	Id             *string                     `json:"id,omitempty"`
	CreatedTime    *time.Time                  `json:"created_time,omitempty"`
	LastEditedTime *time.Time                  `json:"last_edited_time,omitempty"`
	Title          []RichText                  `json:"title,omitempty"`
	Properties     map[string]DatabaseProperty `json:"properties,omitempty"`
	Parent         *Parent                     `json:"parent,omitempty"`
}

func NewDatabase(title string, parent Parent, properties map[string]DatabaseProperty) Database {
	return Database{
		Title: []RichText{
			NewRichText(title),
		},
		Parent:     &parent,
		Properties: properties,
	}
}

func NewRichTextDatabaseProperty() DatabaseProperty {
	return DatabaseProperty{
		RichText: &struct{}{},
	}
}
func NewTitleDatabaseProperty(title *string) DatabaseProperty {
	return DatabaseProperty{
		Name:  title,
		Title: &struct{}{},
	}
}

type DatabaseProperty struct {
	Name     *string   `json:"name,omitempty"`
	Title    *struct{} `json:"title,omitempty"`
	RichText *struct{} `json:"rich_text,omitempty"`
}
