package hatebu

import (
	"fmt"
	"io/ioutil"
)

type AddRequest struct {
	URL          string
	Comment      string
	PostTwitter  bool
	PostFacebook bool
	PostEvernote bool
}

// Add bookmark
func Add(req AddRequest) int {

	form := map[string][]string{
		"url":     {req.URL},
		"comment": {req.Comment},
	}

	if req.PostTwitter {
		form["post_twitter"] = []string{"true"}
	}

	if req.PostFacebook {
		form["post_facebook"] = []string{"true"}
	}

	if req.PostEvernote {
		form["post_evernote"] = []string{"true"}
	}

	c := &client{}
	response, err := c.post("http://api.b.hatena.ne.jp/1/my/bookmark", form)

	if err != nil {
		fmt.Printf("http request error: %v", err)
		return -1
	}

	if _, err := ioutil.ReadAll(response.Body); err != nil {
		fmt.Printf("read body error: %v", err)
		return -1
	}

	return 0
}
