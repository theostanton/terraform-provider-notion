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

resource "notion_database_property_number" "count" {
  database = notion_database.two.id
  name     = "Count"
}

resource "notion_database_property_relation" "to_database_two" {
  database = notion_database.one.id
  name     = "Relation to database Two"
}

resource "notion_database_property_rollup" "sum_of_count" {
  database          = notion_database.one.id
  name              = "Sum of Count"
  relation_property = notion_database_property_relation.to_database_two.name
  rollup_property   = notion_database_property_number.count.name
  function          = "sum"
}