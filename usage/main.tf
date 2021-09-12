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

resource "notion_database" "some_other_database" {
  title = "Some other database"
  parent = var.parent_page_id
  title_column_title = "Name"
}

resource "notion_database_property_relation" "to_some_other" {
  database = notion_database.some_database.id
  name = "Relation to some other"
  database_id = notion_database.some_other_database.id
}


resource "notion_database_property_number" "count" {
  database = notion_database.some_other_database.id
  name = "Count"
}

resource "notion_database_property_rollup" "to_some_other" {
  database = notion_database.some_database.id
  name = "Rollup of some other"
  relation_property = notion_database_property_relation.to_some_other.name
  rollup_property = notion_database_property_number.count.name
  function = "count"
  depends_on = [
    notion_database_property_number.count]
}

/*
resource "notion_database_property_select" "severity" {
  count = 0
  database = notion_database.some_database.id
  name = "Severity title"
  options = {
    "High":"red"
    "Low":"green"
  }
}

resource "notion_database_property_number" "count" {
  count = 0
  database = notion_database.some_database.id
  name = "Count"
  format = "percent"
}

resource "notion_database_property_rich_text" "summary" {
  count = 0
  database = notion_database.some_database.id
  name = "Summary"
}

resource "notion_database_property_date" "date" {
  count = 0
  database = notion_database.some_database.id
  name = "Date"
}

resource "notion_database_property_people" "people" {
  count = 0
  database = notion_database.some_database.id
  name = "People"
}


resource "notion_database_property_checkbox" "checkbox" {
  count = 0
  database = notion_database.some_database.id
  name = "Checkbox"
}

resource "notion_database_property_url" "url" {
  count = 0
  database = notion_database.some_database.id
  name = "URL"
}

resource "notion_database_property_email" "email" {
  count = 0
  database = notion_database.some_database.id
  name = "Email"
}

resource "notion_database_property_created_time" "created_time" {
  count = 0
  database = notion_database.some_database.id
  name = "Created time"
}

resource "notion_database_property_created_by" "created_by" {
  count = 0
  database = notion_database.some_database.id
  name = "Created by"
}

resource "notion_database_property_last_edited_time" "last_edited_time" {
  count = 0
  database = notion_database.some_database.id
  name = "Last edited time"
}

resource "notion_database_property_last_edited_by" "last_edited_by" {
  count = 0
  database = notion_database.some_database.id
  name = "Last edited by"
}

data "notion_database_entries" "some_database_entries" {
  count = 0
  database = notion_database.some_database.id
}
*/