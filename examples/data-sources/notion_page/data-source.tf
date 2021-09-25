data "notion_page" "example" {
  query = "Example"
}

output "page_url" {
  value = data.notion_page.example.url
}