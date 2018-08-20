// Copyright (c) 2018 Yuriy Lisovskiy
//
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package oslapi

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"encoding/base64"
)

const (
	// Base url for www.gnu.org/licenses web page.
	gnuOrg = "https://www.gnu.org/licenses/"

	// Base url for opensource.org/licenses web page.
	openSourceOrg = "https://opensource.org/licenses/"
)

// Returns license name and link by given license key.
func licenseData(key string) (name string, link string, err error) {
	switch key {
	case "bsd-2-clause":
		name = "BSD 2-Clause License"
		link = openSourceOrg + "BSD-2-Clause"
	case "bsd-3-clause":
		name = "BSD 3-Clause License"
		link = openSourceOrg + "BSD-3-Clause"
	case "agpl-3.0":
		name = "GNU Affero General Public License v3.0"
		link = gnuOrg + "agpl-3.0"
	case "gpl-2.0":
		name = "GNU General Public License v2.0"
		link = gnuOrg + "gpl-2.0"
	case "gpl-3.0":
		name = "GNU General Public License v3.0"
		link = gnuOrg + "gpl-3.0"
	case "lgpl-2.1":
		name = "GNU Lesser General Public License v2.1"
		link = gnuOrg + "lgpl-2.1"
	case "lgpl-3.0":
		name = "GNU Lesser General Public License v3.0"
		link = gnuOrg + "lgpl-3.0"
	case "apache-2.0":
		name = "Apache License Version 2.0"
		link = openSourceOrg + "Apache-2.0"
	case "epl-2.0":
		name = "Eclipse Public License - v2.0"
		link = openSourceOrg + "EPL-2.0"
	case "mit":
		name = "MIT License"
		link = openSourceOrg + "MIT"
	case "mpl-2.0":
		name = "Mozilla Public License Version 2.0"
		link = openSourceOrg + "MPL-2.0"
	case "unlicense":
		name = "Unlicense"
		link = "http://unlicense.org"
	case "aal":
		name = "Attribution Assurance License"
		link = "https://opensource.org/licenses/AAL"
	case "afl-3.0":
		name = "Academic Free License 3.0"
		link = "https://opensource.org/licenses/AFL-3.0"
	case "apsl-2.0":
		name = "Apple Public Source License 2.0"
		link = "https://opensource.apple.com/apsl"
	case "artistic-2.0":
		name = "Artistic License 2.0"
		link = "https://opensource.org/licenses/Artistic-2.0"
	case "bsl-1.0":
		name = "Boost Software License 1.0"
		link = "https://opensource.org/licenses/BSL-1.0"
	case "catosl-1.1":
		name = "Computer Associates Trusted Open Source License 1.1"
		link = "https://opensource.org/licenses/CATOSL-1.1"
	case "cecill-2.1":
		name = "CeCILL License 2.1"
		link = "https://opensource.org/licenses/CECILL-2.1"
	default:
		err = ErrLicenseNotFound
	}
	return
}

func downloadContent(url string) ([]byte, error) {
	var ret []byte

	// Download content from https://github.com/YuriyLisovskiy/licenses
	resp, err := http.Get(url)
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ret, err
	}

	// Decode response.
	var response licenseResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return ret, err
	}

	// Decode license content using base64 encoding.
	decoded, err := base64.StdEncoding.DecodeString(response.Content)
	if err != nil {
		return ret, err
	}
	ret = decoded
	if len(ret) == 0 {
		return ret, ErrContentHotFound
	}
	return decoded, nil
}
