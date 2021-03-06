package model

import (
	"errors"
	"github.com/theostanton/terraform-provider-notion/internal/utils/logger"
	"time"
)

type Database struct {
	Object         string                      `json:"object,omitempty"`
	Id             *string                     `json:"id,omitempty"`
	Url            *string                     `json:"url,omitempty"`
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

func (database *Database) ExtractColumnTitleId() (string, error) {
	for key, property := range database.Properties {
		if *property.Type == "title" {
			return key, nil
		}
	}
	return "", errors.New("getColumnTitleId - couldn't find title property")
}

func (database *Database) ExtractTitle() string {
	return database.Title[0].Text.Content
}

func NewBasicDatabaseProperty(name *string, propertyType string) DatabaseProperty {
	switch propertyType {
	case "rich_text":
		return DatabaseProperty{
			Name:     name,
			RichText: &struct{}{},
		}
	case "people":
		return DatabaseProperty{
			Name:   name,
			People: &struct{}{},
		}
	case "file":
		return DatabaseProperty{
			Name: name,
			File: &struct{}{},
		}
	case "checkbox":
		return DatabaseProperty{
			Name:     name,
			Checkbox: &struct{}{},
		}
	case "url":
		return DatabaseProperty{
			Name: name,
			Url:  &struct{}{},
		}
	case "email":
		return DatabaseProperty{
			Name:  name,
			Email: &struct{}{},
		}
	case "created_time":
		return DatabaseProperty{
			Name:        name,
			CreatedTime: &struct{}{},
		}
	case "created_by":
		return DatabaseProperty{
			Name:      name,
			CreatedBy: &struct{}{},
		}
	case "last_edited_time":
		return DatabaseProperty{
			Name:           name,
			LastEditedTime: &struct{}{},
		}
	case "last_edited_by":
		return DatabaseProperty{
			Name:         name,
			LastEditedBy: &struct{}{},
		}
	case "date":
		return DatabaseProperty{
			Name: name,
			Date: &struct{}{},
		}
	}
	logger.Error("unknown property type=%s", propertyType)
	return DatabaseProperty{}
}

func NewNumberDatabaseProperty(name *string, format *string) DatabaseProperty {
	return DatabaseProperty{
		Name: name,
		Number: &NumberDatabasePropertyInfo{
			Format: format,
		},
	}
}

type RollupFunction string

const (
	CountAll          RollupFunction = "count_all"
	CountValues                      = "count_values"
	CountUniqueValues                = "count_unique_values"
	CountEmpty                       = "count_empty"
	CountNotEmpty                    = "count_not_empty"
	PercentEmpty                     = "percent_empty"
	PercentNotEmpty                  = "percent_not_empty"
	Sum                              = "sum"
	Average                          = "average"
	Median                           = "median"
	Min                              = "min"
	Max                              = "max"
	Range                            = "range"
	ShowOriginal                     = "show_original"
)

var RollupFunctions = [...]RollupFunction{
	CountAll,
	CountValues,
	CountUniqueValues,
	CountEmpty,
	CountNotEmpty,
	PercentEmpty,
	PercentNotEmpty,
	Sum,
	Average,
	Median,
	Min,
	Max,
	Range,
	ShowOriginal,
}

type RollupDatabasePropertyInfo struct {
	RelationPropertyName string         `json:"relation_property_name"`
	RollupPropertyName   string         `json:"rollup_property_name"`
	Function             RollupFunction `json:"function"`
}

func NewRollupDatabaseProperty(name *string, rollupPropertyName string, relationPropertyName string, function RollupFunction) DatabaseProperty {
	return DatabaseProperty{
		Name: name,
		Rollup: &RollupDatabasePropertyInfo{
			RelationPropertyName: relationPropertyName,
			RollupPropertyName:   rollupPropertyName,
			Function:             function,
		},
	}
}

func NewRelationDatabaseProperty(name *string, databaseId string) DatabaseProperty {
	return DatabaseProperty{
		Name: name,
		Relation: &RelationDatabasePropertyInfo{
			DatabaseId: databaseId,
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

func NewMultiDatabaseProperty(name *string, options *[]SelectOption) DatabaseProperty {
	return DatabaseProperty{
		Name: name,
		MultiSelect: &MultiSelectDatabasePropertyInfo{
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

type RelationDatabasePropertyInfo struct {
	DatabaseId string `json:"database_id"`
}

type SelectDatabasePropertyInfo struct {
	Options *[]SelectOption `json:"options,omitempty"`
}

type MultiSelectDatabasePropertyInfo struct {
	Options *[]SelectOption `json:"options,omitempty"`
}

type SelectOption struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type DatabaseProperty struct {
	ID             *string                          `json:"id,omitempty"`
	Type           *string                          `json:"type,omitempty"`
	Name           *string                          `json:"name,omitempty"`
	Title          *struct{}                        `json:"title,omitempty"`
	RichText       *struct{}                        `json:"rich_text,omitempty"`
	Date           *struct{}                        `json:"date,omitempty"`
	People         *struct{}                        `json:"people,omitempty"`
	File           *struct{}                        `json:"file,omitempty"`
	Checkbox       *struct{}                        `json:"checkbox,omitempty"`
	Url            *struct{}                        `json:"url,omitempty"`
	Email          *struct{}                        `json:"email,omitempty"`
	CreatedTime    *struct{}                        `json:"created_time,omitempty"`
	CreatedBy      *struct{}                        `json:"created_by,omitempty"`
	LastEditedTime *struct{}                        `json:"last_edited_time,omitempty"`
	LastEditedBy   *struct{}                        `json:"last_edited_by,omitempty"`
	Number         *NumberDatabasePropertyInfo      `json:"number,omitempty"`
	Relation       *RelationDatabasePropertyInfo    `json:"relation,omitempty"`
	Rollup         *RollupDatabasePropertyInfo      `json:"rollup,omitempty"`
	Select         *SelectDatabasePropertyInfo      `json:"select,omitempty"`
	MultiSelect    *MultiSelectDatabasePropertyInfo `json:"multi_select,omitempty"`
}

type DatabasePropertyValue struct {
	ID    string      `json:"id"`
	Type  string      `json:"type"`
	Title *[]RichText `json:"title,omitempty"`
}
