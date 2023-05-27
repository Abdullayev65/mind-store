package url_tool

import (
	"mindstore/pkg/config"
	"mindstore/pkg/hash-types"
	"path"
)

func AvatarUrlWithHash(userId hash.Int) *string {
	url := path.Join(config.GetFilesBaseUrl(), "avatar", userId.HashToStr())
	return &url
}
