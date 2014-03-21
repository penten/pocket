package pocket

import (
	"strings"
	"testing"
)

func TestRequestToken(t *testing.T) {
	_, err := GetRequestToken(Appkey, "http://google.com")

	if err != nil {
		t.Error("Could not get request token: " + err.Error())
	}
}

func TestGet(t *testing.T) {
	list, err := GetArticles(Appkey, apptoken, map[string]string{"count": "2"})

	if err != nil {
		t.Error("Could not get article list: " + err.Error())
	}

	if len(list) != 2 {
		t.Error("Incorrect number of articles in list")
	}

	for _, item := range list {
		if !strings.Contains(item.Resolved_url, "http") {
			t.Error("Invalid resolved url", item.Resolved_url)
		}
	}

}
