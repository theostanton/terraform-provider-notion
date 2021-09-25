data "notion_page" "parent" {
  query = "Examples Parent"
}

resource "notion_page" "example" {
  title  = "Some title"
  parent = data.notion_page.parent.id
}