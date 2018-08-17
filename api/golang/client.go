// Copyright (c) 2018 Yuriy Lisovskiy
//
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package golang

// Represents simple client for getting a license.
type Client struct{}

// Downloads the license from https://github.com/YuriyLisovskiy/licenses repository.
func (Client) GetLicense(name string) (License, error) {
	var license License

	// Download the license.
	licenseContent, err := downloadContent(baseUrl + "/licenses/" + name)
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
	license.content = string(licenseContent)
	return license, nil
}

// Downloads the header from https://github.com/YuriyLisovskiy/licenses repository.
func (Client) GetHeader(name string) (string, error) {
	header, err := downloadContent(baseUrl + "/headers/" + name + "-header")
	if err != nil {
		return "", err
	}
	return string(header), nil
}
