package api

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aisk/cookieparser"
	"github.com/leancloud/lean-cli/lean/utils"
)

// saveCookies saves the cookies to `${HOME}/.leancloud/cookies`
func saveCookies(cookies []*http.Cookie) error {
	os.MkdirAll(filepath.Join(utils.ConfigDir(), "leancloud"), 0700)

	content := []byte(cookieparser.ToString(cookies))
	return ioutil.WriteFile(filepath.Join(utils.ConfigDir(), "leancloud", "cookies"), content, 0600)
}

func getCookies() ([]*http.Cookie, error) {
	content, err := ioutil.ReadFile(filepath.Join(utils.ConfigDir(), "leancloud", "cookies"))
	if err != nil {
		return nil, err
	}
	cookies := cookieparser.Parse(string(content))
	return cookies, nil
}
