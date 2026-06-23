package structs

import (
	"encoding/base64"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

type Picture struct {
	Data     []byte `json:"data"`     // Data
	MimeType string `json:"mimeType"` // MIME type
}

type Base64Picture struct {
	Data string `json:"data"` // Data
}

func getPicture(file multipart.File) (Picture, error) {
	// NOTA: Il file viene passato già aperto. Fare in modo che il file venga passato chiuso.
	var picture Picture
	var err error

	picture.Data, err = io.ReadAll(file)
	if err != nil {
		defer file.Close()
		return Picture{}, err
	}

	picture.MimeType = http.DetectContentType(picture.Data)
	if picture.MimeType != "image/jpeg" && picture.MimeType != "image/png" {
		defer file.Close()
		return Picture{}, errors.New("file type is not correct")
	}

	defer file.Close()
	return picture, nil
}

func convertBinaryPictureToBase64(picture Picture) (Base64Picture, error) {
	var base64Picture Base64Picture
	base64Picture.Data = base64.StdEncoding.EncodeToString(picture.Data)
	return base64Picture, nil
}

func convert64BasePictureToBinary(base64Picture Base64Picture) (Picture, error) {
	var picture Picture
	var err error

	parts := strings.Split(base64Picture.Data, ",")
	if len(parts) != 2 {
		return Picture{}, errors.New("base 64 encoding is wrong")
	}
	picture.MimeType = strings.TrimSuffix(strings.TrimPrefix(parts[0], "data:"), ";base64")
	picture.Data, err = base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return Picture{}, err
	}
	return picture, nil
}
