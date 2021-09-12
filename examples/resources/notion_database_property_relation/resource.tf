resource "notion_database" "one" {
  title              = "Database one"
  parent             = var.parent_page_id
  title_column_title = "Name"
}

resource "notion_database" "two" {
  title              = "Database two"
  parent             = var.parent_page_id
  title_column_title = "Name"
}

resource "notion_database_property_relation" "example" {
  database = notion_database.one.id
  name     = "Relation to database Two"
}