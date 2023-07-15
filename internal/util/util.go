package util

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
	"strings"
)

func Getenv(k string, mustGet bool) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading environment variables file")
	}

	v := os.Getenv(k)
	if mustGet && v == "" {
		log.Fatalf(fmt.Sprintf("Fatal Error: %s environment variable not set.\n", k))
	}
	return v
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func GenerateInstagramCaption(caption, author, hashtags, platformCode string) string {
	maxCharacters := 2200
	captionFormat := "%s\n\n(Via: %s/%s)"
	captionFormatWithTags := "%s\n\n(Via: %s/%s)\n\n\n\n[%s]"
	caption = trimWhitespace(caption)
	if len(fmt.Sprintf(captionFormatWithTags, caption, author, platformCode, "")) > maxCharacters {
		firstCharLen := maxCharacters - len(fmt.Sprintf(captionFormatWithTags, caption, author, platformCode, ""))
		caption = caption[:firstCharLen]
		return fmt.Sprintf(captionFormat, caption, author, platformCode)
	}

	newCaption, replacedHashtags := replaceExcessiveHashtags(caption, 30)
	newCaption = replaceExcessiveAtTags(newCaption, 20)
	endHashtags := ""
	if !replacedHashtags {
		inputHashtags := strings.Fields(hashtags)
		captionLength := len(caption)
		hashtagCount := 0
		for _, word := range strings.Fields(caption) {
			if strings.HasPrefix(word, "#") {
				hashtagCount++
			}
		}

		for _, tag := range inputHashtags {
			if hashtagCount >= 30 || captionLength+len(tag) >= maxCharacters {
				break
			}
			appendSpace := " "
			if len(endHashtags) == 0 {
				appendSpace = ""
			}
			endHashtags += appendSpace + tag
			captionLength += len(tag) + 1
			hashtagCount++
		}
	}

	tags := strings.Fields(endHashtags)
	length := len(tags)
	currLen := len(fmt.Sprintf(captionFormatWithTags, newCaption, author, platformCode, endHashtags))

	for currLen > maxCharacters && length > 0 {
		tags = tags[:length-1]
		length = len(tags)
		endHashtags = strings.Join(tags, " ")
		currLen = len(fmt.Sprintf(captionFormatWithTags, newCaption, author, platformCode, endHashtags))
	}

	if endHashtags == "" {
		newCaption = fmt.Sprintf(captionFormat, newCaption, author, platformCode)
	} else {
		newCaption = fmt.Sprintf(captionFormatWithTags, newCaption, author, platformCode, endHashtags)
	}

	return newCaption
}

func replaceExcessiveHashtags(input string, maxHashtags int) (string, bool) {
	replaced := false
	pattern := `#[^\s#]+`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(input, -1)
	if len(matches) > maxHashtags {
		replaced = true
		for i := maxHashtags; i < len(matches); i++ {
			input = strings.Replace(input, matches[i], "", 1)
		}
	}
	return trimWhitespace(input), replaced
}

func replaceExcessiveAtTags(input string, maxAtTags int) string {
	pattern := `@[^\s@]+`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(input, -1)
	if len(matches) > maxAtTags {
		for i := maxAtTags; i < len(matches); i++ {
			input = strings.Replace(input, matches[i], "", 1)
		}
	}
	return trimWhitespace(input)
}

func trimWhitespace(input string) string {
	regex := regexp.MustCompile(`\s+`)
	result := regex.ReplaceAllString(strings.TrimSpace(input), " ")
	return result
}
