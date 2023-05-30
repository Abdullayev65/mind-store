package url_tool

import (
	"mindstore/pkg/config"
	"mindstore/pkg/hash-types"
	"path"
)

func AvatarUrlWithHash(userId hash.Int) *string {
	url := config.GetFilesUrlWith(path.Join("avatar", userId.HashToStr()))
	return &url
}
