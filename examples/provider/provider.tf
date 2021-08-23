provider "notion" {
  api_token = "secret_123"
}


resource "notion_database" "some_database" {
  title              = "Some database"
  parent             = var.parent_page_id
  title_column_title = "Name"
}

resource "notion_database_property_select" "severity" {
  database = notion_database.some_database.id
  name     = "Severity"
  options = {
    "High" : "red"
    "Low" : "green"
  }
}

resource "notion_database_property_number" "count" {
  database = notion_database.some_database.id
  name     = "Count"
  format   = "percent"
}

resource "notion_database_property_rich_text" "summary" {
  database = notion_database.some_database.id
  name     = "Summary"
}

data "notion_database_entries" "some_database_entries" {
  database = notion_database.some_database.id
}

output "entries" {
  value = data.notion_database_entries.some_database_entries.entries
}