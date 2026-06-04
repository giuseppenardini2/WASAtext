package structs

type Picture struct {
	Data []byte `json:"data"` // Data
}

func IsPictureNil(picture Picture) bool {
	return len(picture.Data) == 0
}
