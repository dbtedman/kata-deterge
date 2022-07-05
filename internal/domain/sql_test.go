package domain_test

import (
	"bytes"
	"github.com/dbtedman/kata-deterge/internal/domain"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCanHandleEmpty(t *testing.T) {
	in := strings.NewReader("")
	out := bytes.Buffer{}

	result := domain.DetergeSql(in, &out)

	assert.Equal(t, domain.EmptyInputError, result)
	assert.Equal(t, 0, out.Len())
}

func TestCanHandleSingleCharacterInput(t *testing.T) {
	in := strings.NewReader(" ")
	out := bytes.Buffer{}

	result := domain.DetergeSql(in, &out)

	assert.Nil(t, result)
	assert.Equal(t, 1, out.Len())
}
