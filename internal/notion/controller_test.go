package notion

import (
	"context"
	"testing"

	"github.com/jomei/notionapi"
)

func Test_controller_AddNewPageToDataBase(t *testing.T) {
	type fields struct {
		cli notionapi.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller{}
			c.AddNewPageToDataBase(context.Background())
		})
	}
}
