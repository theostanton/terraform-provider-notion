package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/theostanton/terraform-provider-notion/internal/model"
	"github.com/theostanton/terraform-provider-notion/internal/utils"
	"github.com/theostanton/terraform-provider-notion/internal/utils/logger"
	"net/http"
)

func (client *Client) GetPage(ctx context.Context, pageId string) (page model.Page, err error) {
	path := fmt.Sprintf("pages/%s", pageId)
	req, err := client.generateGet(ctx, path)
	if err != nil {
		return
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return model.Page{}, fmt.Errorf("failed to find page: %w", parseErrorResponse(res))
	}

	var response *model.Page

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return model.Page{}, fmt.Errorf("failed to parse HTTP response: %w", err)
	}

	normalizedId := utils.NormalizeId(*response.Id)
	response.Id = &normalizedId

	return *response, nil
}

func (client *Client) ArchivePage(ctx context.Context, pageId string) (err error) {
	path := fmt.Sprintf("pages/%s", pageId)

	body := struct {
		archived bool
	}{
		archived: true,
	}

	req, err := client.generatePatch(ctx, path, body)
	if err != nil {
		return
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete page: %w", parseErrorResponse(res))
	}

	return

}

func (client *Client) CreatePage(ctx context.Context, page model.PagePatch) (storedPage model.Page, err error) {

	logger.InfoObject("CreatePage:body", page)

	req, err := client.generatePost(ctx, "pages", page)
	if err != nil {
		return
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to create page: %w", parseErrorResponse(res))
		return
	}

	err = json.NewDecoder(res.Body).Decode(&storedPage)
	if err != nil {
		err = fmt.Errorf("failed to parse HTTP response: %w", err)
		return
	}

	pageId := utils.NormalizeId(*storedPage.Id)
	storedPage.Id = &pageId

	return
}

func (client *Client) UpdatePageTitle(ctx context.Context, pageId string, title string) (err error) {

	path := fmt.Sprintf("pages/%s", pageId)

	body := model.Page{
		Title: []model.RichText{
			model.NewRichText(title),
		},
	}

	req, err := client.generatePatch(ctx, path, body)
	if err != nil {
		return
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update page title: %w", parseErrorResponse(res))
	}

	return nil
}

func (client *Client) UpdatePageParent(ctx context.Context, pageId string, parent model.Parent) (err error) {
	path := fmt.Sprintf("pages/%s", pageId)

	body := model.Page{
		Parent: &parent,
	}

	req, err := client.generatePatch(ctx, path, body)
	if err != nil {
		return
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update page parent: %w", parseErrorResponse(res))
	}

	return nil
}
