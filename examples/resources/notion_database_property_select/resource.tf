resource "notion_database" "example" {
  title              = "Some title"
  parent             = var.parent_page_id
  title_column_title = "Name"
}

resource "notion_database_property_number" "example" {
  database = notion_database.example.id
  name     = "Severity"
  options = {
    "High" : "red"
    "Low" : "green"
  }
}