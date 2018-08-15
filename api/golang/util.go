// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package golang

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
	default:
		err = ErrLicenseNotFound
	}
	return
}
