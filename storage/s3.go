package storage

import (
	"base-site-api/utils"
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// S3Storage provides functionality to store the files
type S3Storage struct {
	config *aws.Config
	bucket string
	acl    string
}

// NewS3 return instance of s3Storage with setup whole config from env
func NewS3() *S3Storage {
	return &S3Storage{
		config: &aws.Config{
			Region:   aws.String(os.Getenv("AWS_DEFAULT_REGION")),
			Endpoint: aws.String(os.Getenv("AWS_ENDPOINT")),
		},
		bucket: os.Getenv("AWS_BUCKET"),
		acl:    "public-read",
	}
}

const (
	// SMALL image path
	SMALL = "resized/small"
	// LARGE image path
	LARGE = "reduced"
)

// UploadFile is return value from Store function from S3Storage
type UploadFile struct {
	URL      string `json:"url"`
	URLSmall string `json:"url-small"`
	IsImage  bool   `json:"-"`
}

// Store file in s3 in path what you want
func (s3 *S3Storage) Store(f *multipart.FileHeader, p string) (*UploadFile, error) {
	ext := s3.getExt(f.Filename)
	filename := fmt.Sprintf("%s.%s", utils.GenerateRandomString(10), ext)
	r, err := f.Open()
	if err != nil {
		return nil, err
	}

	if s3.isImage(ext) {
		var img image.Image

		if ext == "jpg" || ext == "jpeg" {
			img, err = jpeg.Decode(r)
		} else if ext == "png" {
			img, err = png.Decode(r)
		}

		width := 1920
		bounds := img.Bounds()

		if width > bounds.Max.X {
			width = bounds.Max.X
		}

		imgSmall := resize.Resize(300, 0, img, resize.Lanczos3)
		imgLarge := resize.Resize(uint(width), 0, img, resize.Lanczos3)

		largePath, err := s3.uploadImage(imgLarge, ext, path.Join(p, LARGE, filename))
		if err != nil {
			return nil, err
		}

		smallPath, err := s3.uploadImage(imgSmall, ext, path.Join(p, SMALL, filename))
		if err != nil {
			return nil, err
		}

		return &UploadFile{
			URL:      largePath,
			URLSmall: smallPath,
			IsImage:  true,
		}, nil
	}

	url, err := s3.uploadFile(r, path.Join(p, "files", filename))
	if err != nil {
		return nil, err
	}

	return &UploadFile{
		URL:     url,
		IsImage: false,
	}, nil

}

func (s3 *S3Storage) uploadImage(img image.Image, ext string, path string) (string, error) {
	i, err := s3.encodeImage(ext, img)

	if err != nil {
		return "", err
	}

	u, err := s3.uploadToS3(path, &i)
	if err != nil {
		return "", err
	}

	return u.Location, nil
}

func (s3 *S3Storage) uploadFile(file io.Reader, path string) (string, error) {
	w := bytes.Buffer{}
	_, err := w.ReadFrom(file)
	if err != nil {
		return "", err
	}

	u, err := s3.uploadToS3(path, &w)
	if err != nil {
		return "", err
	}

	return u.Location, nil
}

func (s3 *S3Storage) uploadToS3(filePath string, file io.Reader) (*s3manager.UploadOutput, error) {

	sess, err := session.NewSession(s3.config)

	uploader := s3manager.NewUploader(sess)

	res, err := uploader.Upload(&s3manager.UploadInput{
		ACL:    &s3.acl,
		Bucket: aws.String(s3.bucket),
		Key:    aws.String(filePath),
		Body:   file,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s3 *S3Storage) getExt(p string) string {
	s := strings.Split(p, ".")
	ext := s[len(s)-1]
	ext = strings.ToLower(ext)
	if ext == "jpeg" || ext == "jpg" || ext == "png" || ext == "gif" {
		return ext
	}
	return ""
}

func (s3 *S3Storage) encodeImage(ext string, img image.Image) (bytes.Buffer, error) {
	var err error
	w := bytes.Buffer{}

	if ext == "jpg" || ext == "jpeg" {
		err = jpeg.Encode(&w, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	} else if ext == "png" {
		err = png.Encode(&w, img)
	}

	return w, err
}

func (s3 *S3Storage) isImage(ext string) bool {
	for _, b := range []string{"jpg", "jpeg", "png"} {
		if b == ext {
			return true
		}
	}
	return false
}
