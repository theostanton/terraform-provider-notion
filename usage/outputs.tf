output "database_url" {
  value = notion_database.some_database.url
}

output "other_notion_database_property_relation_id" {
  value = notion_database_property_relation.to_some_other.id
}

output "other_notion_database_property_relation_name" {
  value = notion_database_property_relation.to_some_other.name
}

output "other_database_url" {
  value = notion_database.some_other_database.url
}