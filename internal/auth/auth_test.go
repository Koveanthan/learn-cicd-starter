package auth

import (
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	t.Run("checks get auth api function", func(t *testing.T) {
		cases := []struct {
			desc   string
			header http.Header
			want   string
		}{
			{
				"malformed empty header",
				http.Header{"Authorization": []string{"ApiKey key"}},
				"key",
			},
			{
				"malformed empty header",
				http.Header{"Authorization": []string{"ApiKey KDLjsldkjasklj3l290"}},
				"KDLjsldkjasklj3l290",
			},
		}

		for _, c := range cases {
			desc := c.desc
			header := c.header
			want := c.want
			got, err := GetAPIKey(header)
			if err != nil {
				t.Fatalf("%s: want %q got %q with err %v", desc, want, got, err)
			}

			if want != got {
				t.Errorf("Want %s, Got %s", want, got)
			}
		}
	})
}
