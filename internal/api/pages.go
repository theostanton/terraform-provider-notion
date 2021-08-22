package api

import (
	"context"
	"fmt"
	"net/http"
)

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
