package core

type (
	// ListResult defines a generic result of multiple entries
	// combined with the number of total results and the offset.
	ListResult[T any] struct {
		Entries []T   `json:"entries"`
		Total   int64 `json:"total"`
		Offset  int64 `json:"offset"`
	}
)
