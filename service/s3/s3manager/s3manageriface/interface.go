// Package s3manageriface provides an interface for the s3manager package
package s3manageriface

import (
	"io"

	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/s3"
	"github.com/gosemver/aws_aws-sdk-go_v1.4.3-1-g1f24fa1/service/s3/s3manager"
)

// DownloaderAPI is the interface type for s3manager.Downloader.
type DownloaderAPI interface {
	Download(io.WriterAt, *s3.GetObjectInput, ...func(*s3manager.Downloader)) (int64, error)
}

var _ DownloaderAPI = (*s3manager.Downloader)(nil)

// UploaderAPI is the interface type for s3manager.Uploader.
type UploaderAPI interface {
	Upload(*s3manager.UploadInput, ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error)
}

var _ UploaderAPI = (*s3manager.Uploader)(nil)
