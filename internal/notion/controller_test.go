package notion

import (
	"context"
	"fmt"
	"github.com/jomei/notionapi"
	"testing"
)

func Test_controller_test(t *testing.T) {
	cli := notionapi.NewClient("secret_oVEbCPMpf59cCB4wCfkFqcicR4yae4LMWaCcAoM2cH0")

	d, err := cli.Page.Get(context.Background(), notionapi.PageID("05cde83f1fc04030b02a587e4a144c21"))
	fmt.Println(d, err)
}
