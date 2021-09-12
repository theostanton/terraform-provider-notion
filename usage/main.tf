terraform {
  required_providers {
    notion = {
      source = "theostanton/providers/notion"
    }
  }
}

provider "notion" {
  token = var.notion_token
}

resource "notion_database" "some_database" {
  title = "Some database"
  parent = var.parent_page_id
  title_column_title = "Name"
}

resource "notion_database_property_select" "severity" {
  database = notion_database.some_database.id
  name = "Severity title"
  options = {
    "High":"red"
    "Low":"green"
  }
}

//resource "notion_database_property_number" "count" {
//  database = notion_database.some_database.id
//  name = "Count"
//  format = "percent"
//}
//
//resource "notion_database_property_rich_text" "summary" {
//  database = notion_database.some_database.id
//  name = "Summary"
//}
//
//resource "notion_database_property_date" "date" {
//  database = notion_database.some_database.id
//  name = "Date"
//}
//
//resource "notion_database_property_people" "people" {
//  database = notion_database.some_database.id
//  name = "People"
//}
//

resource "notion_database_property_checkbox" "checkbox" {
  database = notion_database.some_database.id
  name = "Checkbox"
}

resource "notion_database_property_url" "url" {
  database = notion_database.some_database.id
  name = "URL"
}

resource "notion_database_property_email" "email" {
  database = notion_database.some_database.id
  name = "Email"
}

resource "notion_database_property_created_time" "created_time" {
  database = notion_database.some_database.id
  name = "Created time"
}

resource "notion_database_property_created_by" "created_by" {
  database = notion_database.some_database.id
  name = "Created by"
}

resource "notion_database_property_last_edited_time" "last_edited_time" {
  database = notion_database.some_database.id
  name = "Last edited time"
}

resource "notion_database_property_last_edited_by" "last_edited_by" {
  database = notion_database.some_database.id
  name = "Last edited by"
}
//
//data "notion_database_entries" "some_database_entries" {
//  database = notion_database.some_database.id
//}

//output "entries" {
//  value = data.notion_database_entries.some_database_entries.entries
//}