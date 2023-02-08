package data

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

//Image - struct in database
type Image struct {
	path   string
	IsUser int
	ID     int
}

//AddImage - Adding Image to database and to localfiles
func AddImage(name string, IsUser int, ID int, r *http.Request) error {
	file, fileheader, err := r.FormFile("FileImage")
	if fileheader == nil || file == nil {
		return errors.New("No image")
	}
	if err != nil {
		return err
	}
	defer file.Close()
	if err := AllowedImages(file, fileheader); err != nil {
		return err
	}

	curImage := Image{path: name, IsUser: IsUser, ID: ID}
	Db.Exec("delete from Images where Path = $1", curImage.path)
	_, err = Db.Exec("insert into Images(Path, IsUser, ID) values ($1, $2, $3)", curImage.path, curImage.IsUser, curImage.ID)
	if err != nil {
		return err
	}

	f, err := os.OpenFile("./public/images/"+name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(f, file)
	return nil
}

//AllowedImages - Checking is Image allowed to read and upload
func AllowedImages(file multipart.File, fileheader *multipart.FileHeader) error {
	size := fileheader.Size
	if size > 20<<20 {
		return errors.New("Image is too big")
	}
	buff := make([]byte, 512)
	if _, err := file.Read(buff); err != nil {
		return errors.New("Can't read Image")
	}
	file.Seek(0, 0)
	contentType := http.DetectContentType(buff)
	if len(contentType) < 6 || contentType[:6] != "image/" {
		return errors.New("It's not image")
	}
	return nil
}

//FindImage with this path
func FindImage(path string) error {
	tmpImage := Image{}
	return Db.QueryRow("select * from Images where Path = $1", path).Scan(&tmpImage.path, &tmpImage.IsUser, &tmpImage.ID)
}
