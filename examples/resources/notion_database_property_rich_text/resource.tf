resource "notion_database" "example" {
  title              = "Some title"
  parent             = var.parent_page_id
  title_column_title = "Name"
}

resource "notion_database_property_rich_text" "example" {
  database = notion_database.example.id
  name     = "Summary"
}