package data

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/KarbozovDamir/forum/internal/models"
)

//AddImage - Adding Image to database and to localfiles
func AddImage(name string, IsUser int, ID int, r *http.Request) error {
	file, fileheader, err := r.FormFile("FileImage")
	if fileheader == nil || file == nil {
		return errors.New("no image")
	}
	if err != nil {
		return err
	}
	defer file.Close()
	if err := AllowedImages(file, fileheader); err != nil {
		return err
	}

	curImage := models.Image{}
	Db.Exec("delete from Images where Path = $1", curImage)
	_, err = Db.Exec("insert into Images(Path, IsUser, ID) values ($1, $2, $3)", curImage, curImage.IsUser, curImage.ID)
	if err != nil {
		return err
	}

	f, err := os.OpenFile("/static/images/"+name, os.O_WRONLY|os.O_CREATE, 0666)
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
	if size > 20<<20 { //req size of image
		return errors.New("image is too big")
	}
	buff := make([]byte, 512)
	if _, err := file.Read(buff); err != nil {
		return errors.New("can't read Image")
	}
	file.Seek(0, 0)
	contentType := http.DetectContentType(buff)
	if len(contentType) < 6 || contentType[:6] != "image/" {
		return errors.New("it's not image")
	}
	return nil
}

//FindImage with this path
func FindImage(path string) error {
	tmpImage := models.Image{}
	return Db.QueryRow("select * from Images where Path = $1", path).Scan(&tmpImage, &tmpImage.IsUser, &tmpImage.ID)
}
