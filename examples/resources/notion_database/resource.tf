data "notion_page" "example" {
  query = "Example"
}

resource "notion_database" "example" {
  title              = "Some title"
  parent             = data.notion_page.example.id
  title_column_title = "Name"
}