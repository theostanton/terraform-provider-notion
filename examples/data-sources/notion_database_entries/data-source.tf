resource "notion_database" "example" {
  title              = "Some title"
  parent             = var.parent_page_id
  title_column_title = "Name"
}

data "notion_database_entries" "example" {
  database = notion_database.example.id
}

output "entries" {
  value = data.notion_database_entries.example.entries
}
