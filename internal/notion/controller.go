package notion

import (
	"context"
	"fmt"

	api "github.com/jomei/notionapi"
)

type controller struct {
	cli api.Client
	//secret_oVEbCPMpf59cCB4wCfkFqcicR4yae4LMWaCcAoM2cH0
}

func (c *controller) AddNewPageToDataBase(ctx context.Context) {
	cli := api.NewClient("secret_oVEbCPMpf59cCB4wCfkFqcicR4yae4LMWaCcAoM2cH0")

	p, err := cli.Page.Get(ctx, "ba85f0c44f3b4ce4b3579dc6f39fa11e")
	fmt.Println(err)

	fmt.Printf("Icon: %v\n  ", *p.Icon)

	page, err := cli.Page.Create(ctx, &api.PageCreateRequest{
		Parent: api.Parent{
			Type:       "database_id",
			DatabaseID: "295540f6c3eb41a5bfa657a169a60610",
		},
		Properties: api.Properties{
			"Name": api.TitleProperty{
				Type: "title",
				Title: []api.RichText{
					{
						Type: "text",
						Text: api.Text{
							Content: "name form golang",
						},
					},
				},
			},
			//"Title": api.RichTextProperty{
			//	RichText: []api.RichText{
			//		{
			//			Type: "text",
			//			Text: api.Text{
			//				Content: "title form golang",
			//			},
			//		},
			//	},
			//},
		},
	})

	fmt.Println(page, err)
}
