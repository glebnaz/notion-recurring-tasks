package notion

import (
	"context"
	"github.com/glebnaz/notion-recurring-tasks/internal/config"
	api "github.com/jomei/notionapi"
)

type Controller interface {
	AddNewPageToDataBase(ctx context.Context, page config.RecurringTask) error
}

func NewController(token string) Controller {
	cli := api.NewClient(api.Token(token))
	return &controller{
		cli: cli,
	}
}

type Page struct {
	Title    string `json:"title"`
	ParentId string `json:"parent_id"`
}
