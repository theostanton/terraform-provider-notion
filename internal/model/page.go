package model

import (
	"errors"
	"time"
)

type Page struct {
	Object         string                        `json:"object,omitempty"`
	Id             *string                       `json:"id,omitempty"`
	CreatedTime    *time.Time                    `json:"created_time,omitempty"`
	LastEditedTime *time.Time                    `json:"last_edited_time,omitempty"`
	Title          []RichText                    `json:"title,omitempty"`
	Properties     *map[string]PagePropertyValue `json:"properties,omitempty"`
	Parent         *Parent                       `json:"parent,omitempty"`
	Url            *string                       `json:"url,omitempty"`
}

type PagePropertyValue struct {
	Type  string      `json:"type"`
	Id    *string     `json:"id,omitempty"`
	Title *[]RichText `json:"title,omitempty"`
}

type PagePatch struct {
	Properties *map[string]interface{} `json:"properties,omitempty"`
	Parent     *Parent                 `json:"parent,omitempty"`
	Archived   bool                    `json:"archived,omitempty"`
	Url        *string                 `json:"url,omitempty"`
}

func NewPage(title string, parent Parent) PagePatch {
	titleRichText := []RichText{
		NewRichText(title),
	}
	properties := map[string]interface{}{
		"title": &titleRichText,
	}
	return PagePatch{
		Properties: &properties,
		Parent:     &parent,
	}
}

// todo clean solution to properties
func (page *Page) ExtractPageTitle() (string, error) {
	for _, property := range *page.Properties {
		if property.Type == "title" {
			title := *property.Title
			return title[0].Text.Content, nil
		}
	}
	return "", errors.New("getColumnTitleId - couldn't find title property")
}
