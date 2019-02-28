package BTagChecker

import (
	"golang.org/x/net/html"
	"io"
	"log"
	"strings"
)

func HasValidClosingTags(content string) bool {
	depth := 0
	madeItToEOF := false

	tokenizer := html.NewTokenizer(strings.NewReader(content))
	for {
		token := tokenizer.Next()

		// handle errors
		if token == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				madeItToEOF = true
				break;
			}

			log.Fatalf("Error tokenizing: %s", tokenizer.Err())
		}

		if token == html.StartTagToken {
			depth += 1
		}

		if token == html.EndTagToken {
			if depth <= 0 {
				break;
			}

			depth -= 1
		}
	}

	if depth == 0 && madeItToEOF {
		return true
	}

	return false
}