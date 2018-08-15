// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package golang

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"encoding/base64"
)

// Represents simple client for getting a license.
type Client struct{}

// Downloads license from https://github.com/YuriyLisovskiy/licenses repository.
func (Client) GetLicense(name string) (License, error) {
	var license License

	// Download license from https://github.com/YuriyLisovskiy/licenses/licenses
	resp, err := http.Get(baseUrl + "/" + name)
	if err != nil {
		return license, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return license, err
	}

	// Decode response.
	var response licenseResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return license, err
	}

	// Decode license content using base64 encoding.
	decoded, err := base64.StdEncoding.DecodeString(*response.Content)
	if err != nil {
		return license, err
	}

	// Get license name and source link by its keyword.
	lName, lLink, err := licenseData(name)
	if err != nil {
		return license, err
	}

	// Setup license data.
	license.name = lName
	license.link = lLink
	license.content = string(decoded)
	return license, nil
}
