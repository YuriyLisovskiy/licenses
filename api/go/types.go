// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package oslapi

// Represents content response from https://github.com/YuriyLisovskiy/licenses repository.
type licenseResponse struct {
	// Encoded license content.
	Content string `json:"content,omitempty"`

	// Content encoding
	Encoding string `json:"encoding,omitempty"`
}

// Represents downloaded license.
type License struct {
	// License title.
	title string

	// License source link.
	link string

	// License content.
	content string
}

// Returns license title.
func (self License) Title() string {
	return self.title
}

// Returns license source link.
func (self License) Link() string {
	return self.link
}

// Returns license content.
func (self License) Content() string {
	return self.content
}
