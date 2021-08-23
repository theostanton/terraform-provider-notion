package model

import "time"

type Page struct {
	Object         string                            `json:"object,omitempty"`
	Id             *string                           `json:"id,omitempty"`
	CreatedTime    *time.Time                        `json:"created_time,omitempty"`
	LastEditedTime *time.Time                        `json:"last_edited_time,omitempty"`
	Title          []RichText                        `json:"title,omitempty"`
	Properties     *map[string]DatabasePropertyValue `json:"properties,omitempty"`
	Parent         *Parent                           `json:"parent,omitempty"`
}
