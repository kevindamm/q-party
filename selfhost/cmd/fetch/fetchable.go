// Copyright (c) 2025 Kevin Damm
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
// github:kevindamm/q-party/cmd/fetch/jarchive/fetchable.go

package fetch

import "io"

type Fetchable interface {
	// Descriptive string used only in debugging / log output.
	String() string

	// The remote URL where this resource can be fetched.
	URL() string
	// A local file path (relative to datapath) where the HTML source is mirrored.
	FilepathHTML() string
	// A local file path (relative to datapath) where the JSON source is stored.
	FilepathJSON() string

	// Converts the resource from its HTML representation
	// into the properties of the current Fetchable instance.
	ParseHTML([]byte) error

	// Writes the JSON-formatted resource to the provided writer, and closes it.
	WriteJSON(io.WriteCloser) error

	// Loads the resource from the provided reader, and closes the reader.
	// Assumes the reader is providing JSON-formatted bytes.
	LoadJSON(io.ReadCloser) error
}
