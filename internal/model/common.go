package model

type Parent struct {
	Type      string `json:"type"`
	Workspace bool  `json:"workspace,omitempty"`
	PageId    string `json:"page_id"`
}

func NewParentFromPageId(pageId string) Parent {
	return Parent{
		Type:   "page_id",
		PageId: pageId,
	}
}
