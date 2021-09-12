output "database_url" {
  value = notion_database.some_database.url
}

output "other_database_url" {
  value = notion_database.some_other_database.url
}