const BASE_URL = 'https://api.github.com/repos/YuriyLisovskiy/licenses/contents';

function downloadContent(url, ret) {
	let xhr = new XMLHttpRequest();
	xhr.onreadystatechange = function()
	{
		if (xhr.readyState === XMLHttpRequest.DONE) {
			if (xhr.status === 200) {
				ret(xhr.responseText)
			} else {
				throw new TypeError('can not download content');
			}
		}
	};
	xhr.open('GET', url, true);
	xhr.send();
}

function parseContent(response) {
	return atob(JSON.parse(response)['content']);
}

function licenseData(l_name, handler) {
	let name = '';
	let link = '';
	let gnuOrg = 'https://www.gnu.org/licenses/';
	let openSourceOrg = 'https://opensource.org/licenses/';
	switch (l_name) {
		case "bsd-2-clause":
			name = "BSD 2-Clause License";
			link = openSourceOrg + "BSD-2-Clause";
			break;
		case "bsd-3-clause":
			name = "BSD 3-Clause License";
			link = openSourceOrg + "BSD-3-Clause";
			break;
		case "agpl-3.0":
			name = "GNU Affero General Public License v3.0";
			link = gnuOrg + "agpl-3.0";
			break;
		case "gpl-2.0":
			name = "GNU General Public License v2.0";
			link = gnuOrg + "gpl-2.0";
			break;
		case "gpl-3.0":
			name = "GNU General Public License v3.0";
			link = gnuOrg + "gpl-3.0";
			break;
		case "lgpl-2.1":
			name = "GNU Lesser General Public License v2.1";
			link = gnuOrg + "lgpl-2.1";
			break;
		case "lgpl-3.0":
			name = "GNU Lesser General Public License v3.0";
			link = gnuOrg + "lgpl-3.0";
			break;
		case "apache-2.0":
			name = "Apache License Version 2.0";
			link = openSourceOrg + "Apache-2.0";
			break;
		case "epl-2.0":
			name = "Eclipse Public License - v2.0";
			link = openSourceOrg + "EPL-2.0";
			break;
		case "mit":
			name = "MIT License";
			link = openSourceOrg + "MIT";
			break;
		case "mpl-2.0":
			name = "Mozilla Public License Version 2.0";
			link = openSourceOrg + "MPL-2.0";
			break;
		case "unlicense":
			name = "Unlicense";
			link = "http://unlicense.org";
			break;
		case "aal":
			name = "Attribution Assurance License";
			link = openSourceOrg + "AAL";
			break;
		case "afl-3.0":
			name = "Academic Free License 3.0";
			link = openSourceOrg + "AFL-3.0";
			break;
		case "apsl-2.0":
			name = "Apple Public Source License 2.0";
			link = "https://opensource.apple.com/apsl";
			break;
		case "artistic-2.0":
			name = "Artistic License 2.0";
			link = openSourceOrg + "Artistic-2.0";
			break;
		case "bsl-1.0":
			name = "Boost Software License 1.0";
			link = openSourceOrg + "BSL-1.0";
			break;
		case "catosl-1.1":
			name = "Computer Associates Trusted Open Source License 1.1";
			link = openSourceOrg + "CATOSL-1.1";
			break;
		case "cecill-2.1":
			name = "CeCILL License 2.1";
			link = openSourceOrg + "CECILL-2.1";
			break;
		default:
			throw new Error('license not found');
	}
	handler(name, link);
}

class License {

    constructor(name, link, content) {
        this._name = name;
        this._link = link;
        this._content = content;
    }

    name() {
        return this._name;
    }

    link() {
        return this._link;
    }

    content() {
        return this._content;
    }
}

class Client {

	constructor() {
		Client.prototype.headerHandler = function() {};
		Client.prototype.licenseHandler = function() {};
		Client.prototype.listHandler = function() {};
	}

	setHeaderHandler(func) {
		Client.prototype.headerHandler = func;
	   return this;
	}

	setLicenseHandler(func) {
		Client.prototype.licenseHandler = func;
		return this;
	}

	setListHandler(func) {
		Client.prototype.listHandler = func;
		return this;
	}

	getLicense(l_name) {
		downloadContent(BASE_URL + '/licenses/' + l_name, function (response) {
			let content = parseContent(response);
			licenseData(l_name, function (name, link) {
				Client.prototype.licenseHandler(new License(name, link, content));
			});
		});
	}

	getHeader(l_name) {
		downloadContent(BASE_URL + '/headers/' + l_name + '-header', function (response) {
			Client.prototype.headerHandler(parseContent(response));
		});
	}

	getList() {
		downloadContent(BASE_URL + '/LICENSE-LIST', function (response) {
			Client.prototype.listHandler(parseContent(response));
		});
	}
}
