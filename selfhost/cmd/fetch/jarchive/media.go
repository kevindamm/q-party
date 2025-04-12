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
// github:kevindamm/q-party/selfhost/cmd/fetch/jarchive/media.go

package main

// This enumeration over available media types is modeled after its equivalent
// MIME type such as image/jpeg, image/png, audio/mpeg, etc.  The default (its
// zero value) is an empty string which implicitly represents an unknown media.
type MimeType string

// Challenges may have zero or more media clues (image, audio, video).  Each is
// represented by its own MediaClue instance.  MediaURL is relative a base URL.
type MediaClue struct {
	MimeType MimeType `json:"mime"`
	MediaURL string   `json:"url"`
}

const (
	MediaImageJPG MimeType = "image/jpeg"
	MediaImagePNG MimeType = "image/png"
	MediaImageSVG MimeType = "image/svg+xml"
	MediaAudioMP3 MimeType = "audio/mpeg"
	MediaVideoMP4 MimeType = "video/mp4"
	MediaVideoMOV MimeType = "video/quicktime"
)
