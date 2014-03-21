package pocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Article struct {
	Item_id        string
	Resolved_title string
	Resolved_url   string
}

type ArticleList struct {
	List map[string]Article
}

func parsePost(uri string, values url.Values) (url.Values, error) {
	resp, err := http.PostForm(uri, values)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Request failed: %d", resp.StatusCode))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Response could not be read")
	}

	values, err = url.ParseQuery(string(body))
	if err != nil {
		return nil, errors.New("Response could not be parsed")
	}

	return values, nil
}

func GetRequestToken(key, uri string) (string, error) {
	values := url.Values{"consumer_key": {key}, "redirect_uri": {uri}}
	values, err := parsePost("https://getpocket.com/v3/oauth/request", values)

	if err != nil {
		return "", err
	}

	code, ok := values["code"]
	if ok {
		redirect_uri := url.QueryEscape(uri + "?code=" + code[0])
		return "https://getpocket.com/auth/authorize?request_token=" + code[0] + "&redirect_uri=" + redirect_uri, nil
	}

	return "", errors.New("Code not found in response")
}

func GetAccessToken(key, code string) (string, string, error) {
	values := url.Values{"consumer_key": {key}, "code": {code}}
	values, err := parsePost("https://getpocket.com/v3/oauth/authorize", values)

	if err != nil {
		return "", "", err
	}

	token, aok := values["access_token"]
	username, bok := values["username"]
	if aok && bok {
		return token[0], username[0], nil
	}

	return "", "", errors.New("Token or username not found in response")
}

func GetArticles(key, token string, options map[string]string) (map[string]Article, error) {
	values := url.Values{"consumer_key": {key}, "access_token": {token}}

	for k, v := range options {
		values.Add(k, v)
	}

	resp, err := http.PostForm("https://getpocket.com/v3/get", values)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Request failed: %d", resp.StatusCode))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Response could not be read")
	}

	fmt.Print(string(body))

	list := ArticleList{}
	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, errors.New("Error parsing body")
	}

	return list.List, nil
}
