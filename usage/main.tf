terraform {
  required_providers {
    notion = {
      source = "theostanton/notion"
      version = "0.0.2"
    }
  }
}

provider "notion" {

}

resource "notion_database" "lol" {

}