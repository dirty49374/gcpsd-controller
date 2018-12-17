package verdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionDB(t *testing.T) {
	db := NewDeployDB()

	db.OnNewImage("nginx:1.1")
	assert.Equal(t, 0, len(db.patterns))

	db.OnNewDeploySpec("nginx:*.*")
	assert.Equal(t, 1, len(db.patterns))

	db.OnNewDeploySpec("nginx:*.*")
	assert.Equal(t, 1, len(db.patterns))

	db.OnNewDeploySpec("nginx:1.*")
	assert.Equal(t, 2, len(db.patterns))

	db.OnPod("nginx:*.*")
}
