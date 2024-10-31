package models

// Verse provides data of a song.
//
// swagger:model Verse
type Verse struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

// Song provides data of a song.
//
// swagger:model Song
type Song struct {
	Id string
	Song string `json:"song"`
	GroupName string `json:"group"`
	ReleaseDate string `json:"releaseDate"`
	Text string `json:"text"`
	Link string `json:"link"`
}

// ErrorResponse represents a structure for API errors.

// swagger:model ErrorResponse
type ErrorResponse struct {
	Message string `json:"message"`
}

// type SongDetail struct {
// 	ReleaseDate string `json:"releaseDate"`
// 	Text string `json:"text"`
// 	Link string `json:"link"`
// }