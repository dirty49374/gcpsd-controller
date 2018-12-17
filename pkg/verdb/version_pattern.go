package verdb

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type VersionPattern struct {
	Pattern string
	regex   *regexp.Regexp
}

func NewVersionPattern(pattern string) (*VersionPattern, error) {
	regex, err := compilePattern(pattern)
	if err != nil {
		return nil, err
	}
	return &VersionPattern{
		Pattern: pattern,
		regex:   regex,
	}, nil
}

func (dp *VersionPattern) ParseImageVersion(imageID string) Version {
	match := dp.regex.FindAllStringSubmatch(imageID, -1)
	if len(match) == 0 {
		return nil
	}

	version := make(Version, len(match[0])-1)
	for i, str := range match[0] {
		if i != 0 {
			v, err := strconv.Atoi(str)
			if err != nil {
				return nil
			}
			version[i-1] = v
		}
	}
	return version
}

func compilePattern(pattern string) (*regexp.Regexp, error) {
	if !strings.Contains(pattern, "*") {
		return nil, fmt.Errorf("%s does not contains *", pattern)
	}
	if strings.Contains(pattern, "**") {
		return nil, fmt.Errorf("%s is wrong pattern", pattern)
	}

	words := strings.Split(pattern, "*")
	for i, word := range words {
		words[i] = regexp.QuoteMeta(word)
	}
	regex := "^" + strings.Join(words, "(\\d+)") + "$"
	return regexp.Compile(regex)
}
