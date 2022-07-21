package notion

import (
	"context"
	"github.com/glebnaz/notion-recurring-tasks/internal/config"
	api "github.com/jomei/notionapi"
	log "github.com/sirupsen/logrus"
)

type controller struct {
	cli *api.Client
}

func (c *controller) AddNewPageToDataBase(ctx context.Context, page config.RecurringTask) error {
	pageRequest := &api.PageCreateRequest{
		Parent: api.Parent{
			Type:       "database_id",
			DatabaseID: api.DatabaseID(page.ParentID),
		},
		Properties: api.Properties{
			//todo delete constant
			"Name":   generateTitle(page.Title),
			"Status": generateStatus(page.Status),
		},
	}

	_, err := c.cli.Page.Create(ctx, pageRequest)
	if err != nil {
		log.Errorf("Error creating page: %s", err)
		return err
	}

	return nil
}

func generateStatus(status config.Status) *api.SelectProperty {
	return &api.SelectProperty{
		Type: "select",
		Select: api.Option{
			Name:  status.Name,
			Color: api.Color(status.Color),
		},
	}
}

func generateTitle(title string) api.TitleProperty {
	return api.TitleProperty{
		Type: "title",
		Title: []api.RichText{
			{
				Type: "text",
				Text: api.Text{
					Content: title,
				},
			},
		},
	}
}
