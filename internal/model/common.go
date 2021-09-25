package model

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Parent struct {
	Type       string `json:"type,omitempty"`
	Workspace  bool   `json:"workspace,omitempty"`
	PageId     string `json:"page_id,omitempty"`
	DatabaseId string `json:"database_id,omitempty"`
}

func NewParentFromPageId(pageId string) Parent {
	return Parent{
		//Type:   "page_id",
		PageId: pageId,
	}
}

func NewParentFromDatabaseId(databaseId string) Parent {
	return Parent{
		//Type:       "database_id",
		DatabaseId: databaseId,
	}
}

func NewWorkspacePageParent() Parent {
	return Parent{
		//Type:      "workspace",
		Workspace: true,
	}
}

func NewParentFromData(data *schema.ResourceData) Parent{

	parentPageId, parentPageOk := data.GetOk("parent_page_id")
	databaseId, databaseOk := data.GetOk("database")
	var parent Parent

	switch {
	case parentPageOk:
		parent = NewParentFromPageId(parentPageId.(string))
	case databaseOk:
		parent = NewParentFromDatabaseId(databaseId.(string))
	default:
		parent = NewWorkspacePageParent()
	}

	return parent
}