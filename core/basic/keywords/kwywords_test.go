package keywords

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAllKeywords(t *testing.T) {
	// This is the set of 25 keywords or reserved words used in Go-code
	validKeywords := []string{
		"break",
		"default",
		"func",
		"interface",
		"select",
		"case",
		"defer",
		"go",
		"map",
		"struct",
		"chan",
		"else",
		"goto",
		"package",
		"switch",
		"const",
		"fallthrough",
		"if",
		"range",
		"type",
		"continue",
		"for",
		"import",
		"return",
		"var",
	}

	var keywords string

	for _, keyword := range validKeywords {
		keywords += keyword + " "
	}

	assert.Equal(t,
		"break default func interface select case defer go map struct chan else "+
			"goto package switch const fallthrough if range type continue for import return var",
		strings.Trim(keywords, " "),
	)
}
