package qparty

// A link to the media accompaniment for a challenge
// (or for the commentary of an episode).
#Media: {
  mime: #MimeType
  url: string
}

// The allowed mime types for media assets (plus text/plain;encoding=UTF-8).
#MimeType: "image/jpeg" | "audio/mpeg" | "video/mp4"
