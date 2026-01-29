package pkg_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spacecafe/miniplate/pkg"
	"github.com/stretchr/testify/assert"
)

func TestLoadContexts(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    map[string]any
		wantErr bool
	}{
		{"file", []string{"user1=file://../test/testdata.json"}, map[string]any{"firstname": "John", "lastname": "Doe"}, false},
		{"web", []string{"user1=https://jsonplaceholder.typicode.com/users/1"}, map[string]any{
			"id":       1,
			"name":     "Leanne Graham",
			"username": "Bret",
			"email":    "Sincere@april.biz",
			"address": map[string]any{
				"street":  "Kulas Light",
				"suite":   "Apt. 556",
				"city":    "Gwenborough",
				"zipcode": "92998-3874",
				"geo": map[string]any{
					"lat": "-37.3159",
					"lng": "81.1496",
				},
			},
			"phone":   "1-770-736-8031 x56442",
			"website": "hildegard.org",
			"company": map[string]any{
				"name":        "Romaguera-Crona",
				"catchPhrase": "Multi-layered client-server neural-net",
				"bs":          "harness real-time e-markets",
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pkg.LoadContexts(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				fmt.Printf("%+v\n", got)
				reflect.DeepEqual(got, tt.want)
			}
		})
	}
}
