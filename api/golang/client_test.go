// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package golang

import "testing"

var GetLicense_TestData = []struct {
	input    string
	expected License
}{
	{
		input: "mit",
		expected: License{
			name: "MIT License",
			link: "https://opensource.org/licenses/MIT",
			content: `MIT License

{{Copyright (c) <year> <author>}}

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`,
		},
	},
	{
		input: "bsd-3-clause",
		expected: License{
			name: "BSD 3-Clause License",
			link: "https://opensource.org/licenses/BSD-3-Clause",
			content: `BSD 3-Clause License

{{Copyright (c) <year> <author>}}
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of the copyright holder nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`,
		},
	},
}

func Test_GetLicense(test *testing.T) {
	client := Client{}
	for _, data := range GetLicense_TestData {
		actual, _ := client.GetLicense(data.input)
		if actual.Name() != data.expected.Name() {
			test.Errorf(
				"client_test.Test_GetLicense:\n\tactual name -> %s\n is not equal to\n\texpected name -> %s",
				actual.Name(), data.expected.Name(),
			)
		}
		if actual.Link() != data.expected.Link() {
			test.Errorf(
				"client_test.Test_GetLicense:\n\tactual link -> %s\n is not equal to\n\texpected link -> %s",
				actual.Link(), data.expected.Link(),
			)
		}
		if actual.Content() != data.expected.Content() {
			test.Errorf(
				"client_test.Test_GetLicense:\n\tactual content -> %s\n is not equal to\n\texpected content -> %s",
				actual.Content(), data.expected.Content(),
			)
		}
	}
}

var GetHeader_TestData = []struct {
	input    string
	expected string
}{
	{
		input: "gpl-3.0",
		expected: `<program name>
{{Copyright (c) <year> <author>}}

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.`,
	},
}

func Test_GetHeader(test *testing.T) {
	client := Client{}
	for _, data := range GetHeader_TestData {
		actual, _ := client.GetHeader(data.input)
		if actual != data.expected {
			test.Errorf(
				"client_test.Test_GetHeader:\n\tactual name -> %s\n is not equal to\n\texpected name -> %s",
				actual, data.expected,
			)
		}
	}
}

var GetLicenseErr_TestData = struct {
	input string
}{
	input: "some-license",
}

func Test_GetLicenseErr(test *testing.T) {
	client := Client{}
	data := GetLicenseErr_TestData
	_, err := client.GetLicense(data.input)
	if err == nil {
		test.Errorf(
			"client_test.Test_GetLicenseErr:\n\tfunc does not return an error for input: %s",
			data.input,
		)
	}
}

var GetHeaderErr_TestData = struct {
	input string
}{
	input: "some-license",
}

func Test_GetHeaderErr(test *testing.T) {
	client := Client{}
	data := GetHeaderErr_TestData
	_, err := client.GetHeader(data.input)
	if err == nil {
		test.Errorf(
			"client_test.Test_GetHeaderErr:\n\tfunc does not return an error for input: %s",
			data.input,
		)
	}
}

var GetList_TestData = struct {
	expected string
}{
	expected: `BSD
  'bsd-2-clause' - BSD 2-Clause License
  'bsd-3-clause' - BSD 3-Clause License

GNU
  'gpl-2.0' - GNU General Public License v2.0
  'gpl-3.0' - GNU General Public License v3.0
  'agpl-3.0' - GNU Affero General Public License v3.0
  'lgpl-2.1' - GNU Lesser General Public License v2.1
  'lgpl-3.0' - GNU Lesser General Public License v3.0

Other
  'aal' - Attribution Assurance License
  'afl-3.0' - Academic Free License 3.0
  'apache-2.0' - Apache License Version 2.0
  'apsl-2.0' - Apple Public Source License 2.0
  'artistic-2.0' - Artistic License 2.0
  'bsl-1.0' - Boost Software License 1.0
  'catosl-1.1' - Computer Associates Trusted Open Source License 1.1
  'cecill-2.1' - CeCILL License 2.1
  'epl-2.0' - Eclipse Public License - v2.0
  'mit' - MIT License
  'mpl-2.0' - Mozilla Public License Version 2.0
  'unlicense' - Unlicense`,
}

func Test_GetList(test *testing.T) {
	client := Client{}
	data := GetList_TestData
	actual, _ := client.GetList()
	if actual != data.expected {
		test.Errorf(
			"client_test.Test_GetList:\n\tactual list -> %s\n is not equal to\n\texpected list -> %s",
			actual, data.expected,
		)
	}
}
