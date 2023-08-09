// Package file 文件操作辅助函数
package file

import (
	"bytes"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"server/pkg/config"
	"server/pkg/helpers"
	"strings"
)

// Put 将数据存入文件
func Put(data []byte, to string) error {
	err := os.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Exists 判断文件是否存在
func Exists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func SaveUploadImage(c *gin.Context, folder string) (string, error) {

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		return "", err
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	// 保存文件
	fileName := randomNameFromUploadFile(fileHeader)

	switch config.Get("file.default") {
	case "local":
		return uploadToLocal(c, fileHeader, "/"+folder+"/", fileName)
	case "s3":
		return uploadToS3(file, "/"+folder+"/", fileName)
	default:
		return "", errors.New("暂不支持该驱动")
	}
}

func randomNameFromUploadFile(file *multipart.FileHeader) string {
	return helpers.RandomString(16) + filepath.Ext(file.Filename)
}

func uploadToLocal(c *gin.Context, file *multipart.FileHeader, dirName string, fileName string) (string, error) {

	// 确保目录存在，不存在创建
	publicPath := "public"
	err := os.MkdirAll(publicPath+dirName, 0755)
	if err != nil {
		return "", err
	}

	avatarPath := publicPath + dirName + fileName
	if err = c.SaveUploadedFile(file, avatarPath); err != nil {
		return "", err
	}

	if !strings.Contains(fileName, "mp3") {
		var avatar string

		// 裁切图片
		img, err := imaging.Open(avatarPath, imaging.AutoOrientation(true))
		if err != nil {
			return avatar, err
		}

		thumbnails := [...]int{128, 256, 512, 1024}
		for i := 0; i < len(thumbnails); i++ {
			resizeAvatar := imaging.Fit(img, thumbnails[i], thumbnails[i], imaging.Lanczos)
			resizeAvatarPath := publicPath + dirName + strings.ReplaceAll(fileName, ".", "-"+cast.ToString(thumbnails[i])+".")
			err = imaging.Save(resizeAvatar, resizeAvatarPath)
			if err != nil {
				return avatar, err
			}
		}
	}

	return dirName + fileName, nil
}

func uploadToS3(file multipart.File, dirName string, fileName string) (string, error) {
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(config.Get("file.disks.s3.region")),
			Credentials: credentials.NewStaticCredentials(
				config.Get("file.disks.s3.key"),
				config.Get("file.disks.s3.secret"),
				// a token will be created when the session it's used.
				"",
			),
		},
	))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(config.Get("file.disks.s3.bucket")),
		Key:         aws.String(dirName + fileName),
		Body:        file,
		ContentType: aws.String("image"),
	})
	if err != nil {
		return "", err
	}

	if !strings.Contains(fileName, "mp3") {
		var path string

		// Download remote image
		response, err := http.Get(config.Get("file.disks.s3.url") + dirName + fileName)
		if err != nil {
			return path, err
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				return
			}
		}(response.Body)

		// Decode image
		img, format, err := image.Decode(response.Body)
		if err != nil {
			return path, err
		}

		// Resize image
		thumbnails := [...]int{128, 256, 512, 1024}
		for i := 0; i < len(thumbnails); i++ {
			// Resize image
			resized := imaging.Fit(img, thumbnails[i], thumbnails[i], imaging.Lanczos)

			// Compress image
			var buf bytes.Buffer
			switch format {
			case "jpg":
				err = jpeg.Encode(&buf, resized, &jpeg.Options{Quality: 80})
			case "jpeg":
				err = jpeg.Encode(&buf, resized, &jpeg.Options{Quality: 80})
			case "png":
				err = png.Encode(&buf, resized)
			// TODO: pdf
			default:
				return path, err
			}
			if err != nil {
				return path, err
			}

			// 上传到 s3
			_, err = uploader.Upload(&s3manager.UploadInput{
				Bucket:      aws.String(config.Get("file.disks.s3.bucket")),
				Key:         aws.String(dirName + strings.ReplaceAll(fileName, ".", "-"+cast.ToString(thumbnails[i])+".")),
				Body:        bytes.NewReader(buf.Bytes()),
				ContentType: aws.String("image"),
			})

			if err != nil {
				return path, err
			}
		}
	}

	return dirName + fileName, nil
}
