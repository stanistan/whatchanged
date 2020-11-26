package formatter

import (
	"github.com/axetroy/whatchanged/option"
	"github.com/pkg/errors"
	"github.com/shurcooL/markdownfmt/markdown"
)

func Format(src []byte, format option.Format) ([]byte, error) {

	switch format {
	case option.FormatJSON:
		return src, nil
	case option.FormatMarkdown:
		return markdown.Process("CHANGELOG.md", src, &markdown.Options{
			Terminal: false,
		})
	default:
		return nil, errors.Errorf("invalid format '%s'", format)
	}

}
