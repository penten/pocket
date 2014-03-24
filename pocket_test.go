package pocket

import (
	"strings"
	"testing"
)

func TestRequestToken(t *testing.T) {
	_, err := GetRequestToken(appkey, "http://google.com")

	if err != nil {
		t.Error("Could not get request token: " + err.Error())
	}
}

func TestGet(t *testing.T) {
	list, err := GetArticles(appkey, apptoken, map[string]string{"count": "2"})

	if err != nil {
		t.Error("Could not get article list: " + err.Error())
	}

	if len(list) != 2 {
		t.Error("Incorrect number of articles in list")
	}

	for _, item := range list {
		if !strings.Contains(item.Url, "http") {
			t.Error("Invalid resolved url", item.Url)
		}
		if item.Favorite > 2 || item.Favorite < 0 {
			t.Error("Invalid favorite value", item.Favorite)
		}
		if item.Status > 3 || item.Status < 0 {
			t.Error("Invalid status", item.Status)
		}
	}

}
