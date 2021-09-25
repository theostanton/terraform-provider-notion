#output "database_url" {
#  value = notion_database.some_database.url
#}
#
#output "other_notion_database_property_relation_id" {
#  value = notion_database_property_relation.to_some_other.id
#}
#
#output "other_notion_database_property_relation_name" {
#  value = notion_database_property_relation.to_some_other.name
#}
#
#output "other_database_url" {
#  value = notion_database.some_other_database.url
#}

output "analytics_database_id" {
  value = data.notion_database.analytics.id
}

output "analytics_database_url" {
  value = data.notion_database.analytics.url
}


output "tests_page_id" {
  value = data.notion_page.tests_page.id
}

output "tests_page_url" {
  value = data.notion_page.tests_page.url
}


output "some_database_entry_id" {
  value = notion_database_entry.some_child_of_some_workspace_page.id
}

output "some_database_entry_url" {
  value = notion_database_entry.some_child_of_some_workspace_page.url
}

