package cmd

import (
	gt "github.com/bas24/googletranslatefree"
)

func translatetxt(text string ,lang string) string { 
	result, _ := gt.Translate(text, "auto", lang)
	return result
}