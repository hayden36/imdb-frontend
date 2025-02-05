package utils

import (
	"strings"
)

func RemoveRefs(link string) string {
	link = strings.Split(link, "?ref")[0]
	return link
}

func ChangeImageParams(link string) string {
	var imagePath string = strings.Split(link, "@.")[0]
	// https://stackoverflow.com/a/73501833
	return imagePath + "@." + "_V1_QL100_UY256_" + ".jpg"
}
