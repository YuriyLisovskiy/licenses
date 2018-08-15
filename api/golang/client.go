package golang

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"encoding/base64"

	"github.com/YuriyLisovskiy/licenses/api"
)

type Client struct {}

func (Client) GetLicense(name string) (License, error) {
	var license License
	resp, err := http.Get(api.BaseUrl + "/" + name)
	if err != nil {
		return license, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return license, err
	}
	var response LicenseResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return license, err
	}
	decoded, err := base64.StdEncoding.DecodeString(*response.Content)
	if err != nil {
		return license, err
	}
	license.Content = string(decoded)
	return license, nil
}
