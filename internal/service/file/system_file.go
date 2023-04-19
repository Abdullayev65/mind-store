package file

import (
	"errors"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"mindstore/internal/object/model"
	"mindstore/pkg/stream"
	"os"
)

type SystemFile struct{}

func NewSystemFile() *SystemFile {
	return &SystemFile{}
}

func (s *SystemFile) Upload(file *multipart.FileHeader, folder string) (string, error) {
	if file == nil {
		return "", errors.New("file.Upload: file is null")
	}

	filename := uuid.New().String()

	if _, err := os.Stat("./files/" + folder); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll("./files/"+folder, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	dst := "./files/" + folder + "/" + filename

	src, err := file.Open()
	if err != nil {
		return "", err
	}

	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}

	defer out.Close()

	_, err = io.Copy(out, src)

	if err != nil {
		return "err", err
	}

	return dst[1:], nil
}

func (s *SystemFile) Delete(dst string) error {
	return os.Remove("." + dst)
}

func (s *SystemFile) MultipleUpload(files []*multipart.FileHeader, folder string) ([]string, error) {
	var links []string

	for _, file := range files {
		link, err := s.Upload(file, folder)
		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	return links, nil
}

func (s *SystemFile) UploadFile(file *multipart.FileHeader, folder string) (*model.FileData, error) {
	path, err := s.Upload(file, folder)
	if err != nil {
		return nil, err
	}

	fd := new(model.FileData)
	fd.Path = path
	fd.Name = file.Filename
	fd.Size = file.Size

	return fd, nil
}

func (s *SystemFile) MultipleUploadFile(files []*multipart.FileHeader, folder string) (fds []*model.FileData, err error) {
	fds = make([]*model.FileData, 0, len(files))

	defer func() {
		if err == nil {
			return
		}
		stream.ForEach(fds, func(f *model.FileData) {
			s.Delete(f.Path)
		})
	}()

	for _, file := range files {
		fd, er1 := s.UploadFile(file, folder)
		if er1 != nil {
			err = er1
			return nil, err
		}

		fds = append(fds, fd)
	}

	return
}
