// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package golang

// Represents content response from https://github.com/YuriyLisovskiy/licenses repository.
type licenseResponse struct {
	// Encoded license content.
	Content string `json:"content,omitempty"`

	// Content encoding
	Encoding string `json:"encoding,omitempty"`
}

// Represents downloaded license.
type License struct {
	// License full name.
	name string

	// License source link.
	link string

	// License content.
	content string
}

// Returns license full name.
func (self License) Name() string {
	return self.name
}

// Returns license source link.
func (self License) Link() string {
	return self.link
}

// Returns license content.
func (self License) Content() string {
	return self.content
}
