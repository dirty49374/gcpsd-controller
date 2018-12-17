package verdb

import (
	"reflect"
	"testing"
)

func mustCompile(t *testing.T, pattern string) {
	_, err := compilePattern(pattern)
	if err != nil {
		t.Errorf("'%s' must compile, but not (%s)", pattern, err)
	}
}

func mustNotCompile(t *testing.T, pattern string) {
	regex, err := compilePattern(pattern)
	if err == nil {
		t.Errorf("'%s' must not compile, but %s", pattern, regex)
	}
}
func TestCompile(t *testing.T) {
	mustNotCompile(t, "")
	mustNotCompile(t, "s")
	mustNotCompile(t, "\\")
	mustNotCompile(t, "**")
	mustCompile(t, "*")
	mustCompile(t, ".*")
	mustCompile(t, "*.")
	mustCompile(t, "*.*")
}

func mustEqual(t *testing.T, pattern, imageId string, version Version) {
	vp, err := NewVersionPattern(pattern)
	if err != nil {
		t.Errorf("must compile pattern %s", pattern)
	}

	matchedVersion := vp.ParseImageVersion(imageId)
	if !reflect.DeepEqual(version, matchedVersion) {
		t.Errorf("version should be %v, not %v", version, matchedVersion)
	}
}

func TestParseVersion(t *testing.T) {
	mustEqual(t, "*", "1", Version{1})
	mustEqual(t, "*", "11", Version{11})
	mustEqual(t, "*", "01", Version{1})
	mustEqual(t, "*", "-1", nil)
	mustEqual(t, "aaa/bbb:*.*-dev", "aaa/bbb:0.1-dev", Version{0, 1})
	mustEqual(t, "aaa/bbb:*.*-dev", "aaa/bb:0.1-dev", nil)
	mustEqual(t, "*/*", "1/2", Version{1, 2})
	mustEqual(t, "*,*,*,*", "1,2,3,4", Version{1, 2, 3, 4})
}
