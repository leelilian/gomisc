package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	
	"golang.org/x/oauth2"
)

const htmlIndex = `<html><body>
<a href="/GoogleLogin">Log in with Google</a>
</body></html>
`

var endpotin = oauth2.Endpoint{
	AuthURL:  "https://github.com/login/oauth/authorize",
	TokenURL: "https://github.com/login/oauth/access_token",
}

var googleOauthConfig = &oauth2.Config{
	ClientID:     "d6ae439cc8b45e1a0789",
	ClientSecret: "3ed12c59d1bc3c7c5a821dd1a28ad88716ba61e7",
	RedirectURL:  "http://localhost:8000/callback",
	Scopes:       []string{},
	Endpoint:     endpotin,
}

const oauthStateString = "random"

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/GoogleLogin", handleGoogleLogin)
	http.HandleFunc("/callback", handleGoogleCallback)
	fmt.Println(http.ListenAndServe(":8000", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlIndex)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString, )
	fmt.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Println(state)
	
	code := r.FormValue("code")
	fmt.Println(code)
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	fmt.Println(token)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	
	/*
	response, err := http.Get("https://api.github.com/user?access_token=" + token.AccessToken)
	*/
	
	client := http.Client{}
	request, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	request.Header.Add("Authorization", "token "+token.AccessToken);
	response, _ := client.Do(request)
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	fmt.Fprintf(w, "Content: %s\n", contents)
}
