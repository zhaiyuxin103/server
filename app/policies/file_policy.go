package policies

import (
	"github.com/spf13/cast"
	"server/app/models/file"
	"server/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyFile(c *gin.Context, fileModel file.File) bool {
	return cast.ToUint64(auth.CurrentUID(c)) == fileModel.UserID
}

// func CanViewFile(c *gin.Context, fileModel file.File) bool {}
// func CanCreateFile(c *gin.Context, fileModel file.File) bool {}
// func CanUpdateFile(c *gin.Context, fileModel file.File) bool {}
// func CanDeleteFile(c *gin.Context, fileModel file.File) bool {}
