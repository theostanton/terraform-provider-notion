data "notion_database" "example" {
  query = "Example"
}

output "database_url" {
  value = data.notion_database.example.url
}