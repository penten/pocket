Interface to the pocket API

To run tests, a config.go file should be added containing the following values:

- Appkey string
- apptoken string

Authentication can be accomplished as in this example:

```Go
package main

import (
	"fmt"
	"github.com/penten/pocket"
	"net/http"
)

func pocketAuthHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")

	if code == "" {
		url, err := pocket.GetRequestToken(pocket.Appkey, "http://localhost:8080/update-pocket")

		if err != nil {
			fmt.Fprintf(w, "Error getting auth token: "+err.Error())
		}
		http.Redirect(w, r, url, http.StatusFound)
	} else {
		token, username, _ := pocket.GetAccessToken(pocket.Appkey, code)

		fmt.Fprintf(w, "%s has been authorised with the token %s", username, token)
	}
}

func main() {
	http.HandleFunc("/update-pocket", pocketAuthHandler)
	http.ListenAndServe(":8080", nil)
}
```