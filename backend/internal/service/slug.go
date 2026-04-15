package service

import (
	"crypto/rand"
	"math/big"
	"regexp"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"
const nanoidLen = 8

func slugify( title string ) string {
	s := strings.ToLower(title)
	s = strings.TrimSpace(s)

	s = strings.ReplaceAll(s , " " , "-")
	s = strings.ReplaceAll(s , "_" , "-")

	reg := regexp.MustCompile(`[^a-z0-9-]`)
	s = reg.ReplaceAllString(s , "")

	reg = regexp.MustCompile("-+")
	s = reg.ReplaceAllString(s , "-")

	s = strings.Trim(s , "-")

	return s
}

func nanoid() ( string , error ) {
	result := make([]byte , nanoidLen)

	for i := range result {
		n , err := rand.Int(rand.Reader , big.NewInt(int64(len(alphabet))))
		if err != nil {
			return "" , err
		}
		result[i] = alphabet[n.Int64()]
	}

	return string(result) , nil
}

func generateSlug( title string ) (string , error) {
	base := slugify(title)
	id , err := nanoid()
	if err != nil {
		return  "" , err
	}

	if base == "" {
		return id , nil
	}

	return base + "-" + id , nil
}