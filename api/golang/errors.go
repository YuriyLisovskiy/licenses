// Copyright (c) 2018 Yuriy Lisovskiy
//
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package golang

import "errors"

var (
	ErrHeaderNotFound  = errors.New("header does not exist")
	ErrLicenseNotFound = errors.New("license does not exist")
	ErrContentHotFound = errors.New("content does not exist")
)
