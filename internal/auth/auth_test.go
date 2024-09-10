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
				"Proper key 1",
				http.Header{"Authorization": []string{"ApiKey key"}},
				"key",
			},
			{
				"Proper key 2",
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

	t.Run("get auth api should fail for cases other than ApiKey", func(t *testing.T) {
		cases := []struct {
			desc   string
			header http.Header
			want   string
		}{
			{
				"auth using basic",
				http.Header{"Authorization": []string{"Basic key"}},
				"key",
			},
			{
				"malformed auth header - no space between ApiKey and value",
				http.Header{"Authorization": []string{"ApiKeyKDLjsldkjasklj3l290"}},
				"KDLjsldkjasklj3l290",
			},
		}

		for _, c := range cases {
			header := c.header
			_, err := GetAPIKey(header)
			if err == nil {
				t.Fatalf("Expected malformed header error to be thrown")
			}
		}
	})

	t.Run("should throw error if no header is present", func(t *testing.T) {
		header := http.Header{}
		_, err := GetAPIKey(header)
		if err == nil {
			t.Fatalf("Expected not found header error to be thrown")
		}
	})
}
