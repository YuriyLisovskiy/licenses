// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package golang

import "testing"

var licenseData_TestData = []struct {
	input        string
	expectedName string
	expectedLink string
}{
	{
		input:        "bsd-2-clause",
		expectedName: "BSD 2-Clause License",
		expectedLink: openSourceOrg + "BSD-2-Clause",
	},
	{
		input:        "bsd-3-clause",
		expectedName: "BSD 3-Clause License",
		expectedLink: openSourceOrg + "BSD-3-Clause",
	},
	{
		input:        "agpl-3.0",
		expectedName: "GNU Affero General Public License v3.0",
		expectedLink: gnuOrg + "agpl-3.0",
	},
	{
		input:        "gpl-2.0",
		expectedName: "GNU General Public License v2.0",
		expectedLink: gnuOrg + "gpl-2.0",
	},
	{
		input:        "gpl-3.0",
		expectedName: "GNU General Public License v3.0",
		expectedLink: gnuOrg + "gpl-3.0",
	},
	{
		input:        "lgpl-2.1",
		expectedName: "GNU Lesser General Public License v2.1",
		expectedLink: gnuOrg + "lgpl-2.1",
	},
	{
		input:        "lgpl-3.0",
		expectedName: "GNU Lesser General Public License v3.0",
		expectedLink: gnuOrg + "lgpl-3.0",
	},
	{
		input:        "apache-2.0",
		expectedName: "Apache License Version 2.0",
		expectedLink: openSourceOrg + "Apache-2.0",
	},
	{
		input:        "epl-2.0",
		expectedName: "Eclipse Public License - v2.0",
		expectedLink: openSourceOrg + "EPL-2.0",
	},
	{
		input:        "mit",
		expectedName: "MIT License",
		expectedLink: openSourceOrg + "MIT",
	},
	{
		input:        "mpl-2.0",
		expectedName: "Mozilla Public License Version 2.0",
		expectedLink: openSourceOrg + "MPL-2.0",
	},
	{
		input:        "unlicense",
		expectedName: "Unlicense",
		expectedLink: "http://unlicense.org",
	},
	{
		input:        "aal",
		expectedName: "Attribution Assurance License",
		expectedLink: openSourceOrg + "AAL",
	},
	{
		input:        "afl-3.0",
		expectedName: "Academic Free License 3.0",
		expectedLink: openSourceOrg + "AFL-3.0",
	},
	{
		input:        "apsl-2.0",
		expectedName: "Apple Public Source License 2.0",
		expectedLink: "https://opensource.apple.com/apsl",
	},
	{
		input:        "artistic-2.0",
		expectedName: "Artistic License 2.0",
		expectedLink: openSourceOrg + "Artistic-2.0",
	},
	{
		input:        "bsl-1.0",
		expectedName: "Boost Software License 1.0",
		expectedLink: openSourceOrg + "BSL-1.0",
	},
	{
		input:        "catosl-1.1",
		expectedName: "Computer Associates Trusted Open Source License 1.1",
		expectedLink: openSourceOrg + "CATOSL-1.1",
	},
	{
		input:        "cecill-2.1",
		expectedName: "CeCILL License 2.1",
		expectedLink: openSourceOrg + "CECILL-2.1",
	},
}

func Test_licenseData(test *testing.T) {
	for _, data := range licenseData_TestData {
		actualName, actualLink, _ := licenseData(data.input)
		if actualName != data.expectedName {
			test.Errorf(
				"util_test.Test_licenseData:\n\tactual name -> %s\n is not equal to\n\texpected name -> %s",
				actualName, data.expectedName,
			)
		}
		if actualLink != data.expectedLink {
			test.Errorf(
				"util_test.Test_licenseData:\n\tactual link -> %s\n is not equal to\n\texpected link -> %s",
				actualLink, data.expectedLink,
			)
		}
	}
}

var licenseDataErr_TestData = struct {
	input    string
	expected error
}{
	input: "some-license",
	expected: ErrLicenseNotFound,
}

func Test_licenseDataErr(test *testing.T) {
	data := licenseDataErr_TestData
	_, _, actual := licenseData(data.input)
	if actual == nil {
		test.Errorf(
			"util_test.Test_licenseDataErr:\n\tfunc does not return an error for input: %s",
			data.input,
		)
	}
	if actual.Error() != data.expected.Error() {
		test.Errorf(
			"util_test.Test_licenseDataErr:\n\tfunc return invalid error for input: %s",
			data.input,
		)
	}
}
