provider "notion" {
  token = "secret_123"
}

data "notion_page" "example_parent" {
  query = "Example Parent"
}

data "notion_database" "example_database" {
  query = "Example Database"
}

resource "notion_database" "one" {
  title              = "Database one"
  parent             = data.notion_page.example_parent.id
  title_column_title = "Name"
}

resource "notion_database_property_number" "count" {
  database = notion_database.one.id
  name     = "Count"
}

resource "notion_database_property_relation" "to_example_database" {
  database         = notion_database.one.id
  name             = "Relation to database Two"
  related_database = data.notion_database.example_database.id
}

resource "notion_database_property_rollup" "sum_of_count" {
  database          = notion_database.one.id
  name              = "Sum of Count"
  relation_property = notion_database_property_relation.to_example_database.name
  rollup_property   = notion_database_property_number.count.name
  function          = "sum"
}

resource "notion_database_property_select" "severity" {
  database = notion_database.one.id
  name     = "Severity"
  options = {
    "High" : "red"
    "Low" : "green"
  }
}

resource "notion_database_property_number" "count" {
  database = notion_database.one.id
  name     = "Count"
  format   = "percent"
}


resource "notion_database_property_rich_text" "summary" {
  database = notion_database.one.id
  name     = "Summary"
}

resource "notion_database_property_date" "date" {
  database = notion_database.one.id
  name     = "Date"
}

resource "notion_database_property_people" "people" {
  database = notion_database.one.id
  name     = "People"
}

resource "notion_database_property_file" "file" {
  database = notion_database.one.id
  name     = "File"
}

resource "notion_database_property_checkbox" "checkbox" {
  database = notion_database.one.id
  name     = "Checkbox"
}

resource "notion_database_property_url" "url" {
  database = notion_database.one.id
  name     = "URL"
}

resource "notion_database_property_email" "email" {
  database = notion_database.one.id
  name     = "Email"
}

resource "notion_database_property_created_time" "created_time" {
  database = notion_database.one.id
  name     = "Created time"
}

resource "notion_database_property_created_by" "created_by" {
  database = notion_database.one.id
  name     = "Created by"
}

resource "notion_database_property_last_edited_time" "last_edited_time" {
  database = notion_database.one.id
  name     = "Last edited time"
}

resource "notion_database_property_last_edited_by" "last_edited_by" {
  database = notion_database.one.id
  name     = "Last edited by"
}

data "notion_database_entries" "some_database_entries" {
  database = notion_database.one.id
}

output "entries" {
  value = data.notion_database_entries.some_database_entries.entries
}

output "database_url" {
  value = notion_database.one.url
}