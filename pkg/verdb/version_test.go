package verdb

import (
	"testing"
)

func mustNewer(t *testing.T, old, new Version) {
	if !new.IsNewerThan(old) {
		t.Errorf("%v should be newer than %old", new, old)
	}
}

func mustNotNewer(t *testing.T, old, new Version) {
	if new.IsNewerThan(old) {
		t.Errorf("%v should be newer than %old", new, old)
	}
}

func Test(t *testing.T) {
	mustNewer(t, Version{0, 1}, Version{0, 2})
	mustNewer(t, Version{0, 1}, Version{1, 2})
	mustNewer(t, Version{1, 2}, Version{1, 3})

	mustNotNewer(t, Version{1, 2}, Version{1, 2})
	mustNotNewer(t, Version{1, 2}, Version{1, 1})
	mustNotNewer(t, Version{1, 2}, Version{10, 10, 10})
	mustNotNewer(t, Version{}, Version{10, 10, 10})
	mustNotNewer(t, Version{}, Version{})
}
