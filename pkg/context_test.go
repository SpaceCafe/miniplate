package pkg_test

import (
	"reflect"
	"testing"

	"github.com/spacecafe/miniplate/pkg"
	"github.com/stretchr/testify/require"
)

func TestLoadContexts(t *testing.T) {
	t.Parallel()

	tests := []struct {
		want    map[string]any
		name    string
		input   []string
		wantErr bool
	}{
		{
			name:  "file",
			input: []string{"user1=file://../test/testdata.json"},
			want:  map[string]any{"firstname": "John", "lastname": "Doe"},
		},
		{
			name:  "web",
			input: []string{"user1=https://jsonplaceholder.typicode.com/users/1"},
			want: map[string]any{
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := pkg.LoadContexts(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				t.Log(err)
			} else {
				require.NoError(t, err)
				t.Log(got)
				reflect.DeepEqual(got, tt.want)
			}
		})
	}
}
