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

func NewRichTextDatabaseProperty(name *string) DatabaseProperty {
	return DatabaseProperty{
		Name:     name,
		RichText: &struct{}{},
	}
}

func NewNumberDatabaseProperty(name *string, format *string) DatabaseProperty {
	return DatabaseProperty{
		Name: name,
		Number: &NumberDatabasePropertyInfo{
			Format: format,
		},
	}
}

func NewSelectDatabaseProperty(name *string, options *[]SelectOption) DatabaseProperty {
	return DatabaseProperty{
		Name: name,
		Select: &SelectDatabasePropertyInfo{
			Options: options,
		},
	}
}

func NewTitleDatabaseProperty(name string) DatabaseProperty {
	return DatabaseProperty{
		Name:  &name,
		Title: &struct{}{},
	}
}

type NumberDatabasePropertyInfo struct {
	Format *string `json:"format,omitempty"`
}

type SelectDatabasePropertyInfo struct {
	Options *[]SelectOption `json:"options,omitempty"`
}

type SelectOption struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type DatabaseProperty struct {
	Name     *string                     `json:"name,omitempty"`
	Title    *struct{}                   `json:"title,omitempty"`
	RichText *struct{}                   `json:"rich_text,omitempty"`
	Number   *NumberDatabasePropertyInfo `json:"number,omitempty"`
	Select   *SelectDatabasePropertyInfo `json:"select,omitempty"`
}

type DatabasePropertyValue struct {
	ID    string      `json:"id"`
	Type  string      `json:"type"`
	Title *[]RichText `json:"title,omitempty"`
}
