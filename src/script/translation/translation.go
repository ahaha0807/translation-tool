package translation

import (
	"os"
	"net/http"
	"io/ioutil"
	"net/url"
	"regexp"
)

func ToEnglish(str string) string {
	apiToken := getAPIToken()

	rawResult := doTranslate(str, "ja", "en", apiToken)
	result := format(rawResult)

	return result
}

func ToJapanese(str string) string {
	apiToken := getAPIToken()

	rawResult := doTranslate(str, "en", "ja", apiToken)
	result := format(rawResult)

	return result
}

func getAPIToken() string {
	apiKey := string(os.Getenv("MICROSOFT_TRANSLATE_APIKEY"))

	client := &http.Client{}

	req, _ := http.NewRequest("POST", "https://api.cognitive.microsoft.com/sts/v1.0/issueToken", nil)
	req.Header.Set("Ocp-Apim-Subscription-Key", string(apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	apiToken, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	return string(apiToken)
}

func doTranslate(str string, srcLanguage string, destLanguage string, token string) string {
	host := "https://api.microsofttranslator.com"
	path := "/V2/Http.svc/Translate"

	values := url.Values{}
	values.Set("to", destLanguage)
	values.Add("from", srcLanguage)
	values.Add("text", str)

	option := values.Encode()

	client := &http.Client{}

	s := host + path + "?" + option

	req, _ := http.NewRequest("GET", s, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func format(raw string) string {
	re := regexp.MustCompile(">.*<")
	result := string(re.Find([]byte(raw)))

	return result[1:len(result)-1]
}
