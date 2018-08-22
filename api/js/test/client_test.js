let assert = require('assert');
let oslapi = require('../src/oslapi');

describe('Client', function () {
    it('Should return license', function () {
        let client = new oslapi.Client();
        client.setLicenseHandler(function (license) {
            assert.equal(`BSD 2-Clause License

{{Copyright (c) <year> <author> All rights reserved.}}

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

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
`, license.content())
        }).getLicense('bsd-2-clause');
    });
    it('Should return header', function () {
        let client = new oslapi.Client();
        client.setHeaderHandler(function (header) {
            assert.equal(`{{<program name> by <author>}}

To the extent possible under law, the person who associated CC0 with
<program name> has waived all copyright and related or neighboring rights
to <program name>.

You should have received a copy of the CC0 legalcode along with this
work.  If not, see <http://creativecommons.org/publicdomain/zero/1.0/>.`, header);
        }).getHeader('cc0-header');
    });
    it('Should return list', function () {
        let client = new oslapi.Client();
        client.setListHandler(function (list) {
            assert.equal(`BSD:
 'bsd-2-clause'  - BSD 2-Clause License
 'bsd-3-clause'  - BSD 3-Clause License

Creative Commons:
  'cc0'          - CC0 1.0 Universal
  'cc-by'        - Attribution 3.0 Unported
  'cc-by-nc'     - Attribution-NonCommercial 3.0 Unported
  'cc-by-nc-nd'  - Attribution-NonCommercial-NoDerivs 3.0 Unported
  'cc-by-nc-sa'  - Attribution-NonCommercial-ShareAlike 3.0 Unported
  'cc-by-nd'     - Attribution-NoDerivs 3.0 Unported
  'cc-by-sa'     - Attribution-ShareAlike 3.0 Unported

GNU:
  'gpl-2.0'      - GNU General Public License v2.0
  'gpl-3.0'      - GNU General Public License v3.0
  'agpl-3.0'     - GNU Affero General Public License v3.0
  'lgpl-2.1'     - GNU Lesser General Public License v2.1
  'lgpl-3.0'     - GNU Lesser General Public License v3.0

Other:
  'aal'          - Attribution Assurance License
  'afl-3.0'      - Academic Free License 3.0
  'apache-2.0'   - Apache License Version 2.0
  'apsl-2.0'     - Apple Public Source License 2.0
  'artistic-2.0' - Artistic License 2.0
  'bsl-1.0'      - Boost Software License 1.0
  'catosl-1.1'   - Computer Associates Trusted Open Source License 1.1
  'cddl-1.0'     - COMMON DEVELOPMENT AND DISTRIBUTION LICENSE Version 1.0
  'cecill-2.1'   - CeCILL License 2.1
  'epl-2.0'      - Eclipse Public License - v2.0
  'mit'          - MIT License
  'mpl-2.0'      - Mozilla Public License Version 2.0
  'unlicense'    - Unlicense
  'wtfpl'        - DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE Version 2
  'x11'          - X11
  'zlib'         - ZLIB`, list);
        }).getList();
    });
});
