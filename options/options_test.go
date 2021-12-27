package options

import (
	"strings"
	"testing"

	gem "git.sr.ht/~kota/goldmark-gemtext"
	"github.com/google/go-cmp/cmp"
)

func TestOptions(t *testing.T) {
	tests := []struct {
		args []string
		opts []gem.Option
	}{
		{
			[]string{"-e", "markdown"},
			[]gem.Option{gem.WithEmphasis(gem.EmphasisMarkdown), gem.WithCodeSpan(gem.CodeSpanMarkdown), gem.WithStrikethrough(gem.StrikethroughMarkdown), gem.WithHorizontalRule("~~~")},
		},
		{
			[]string{"-e", "unicode"},
			[]gem.Option{gem.WithEmphasis(gem.EmphasisUnicode), gem.WithStrikethrough(gem.StrikethroughUnicode), gem.WithHorizontalRule("~~~")},
		},
		{
			[]string{"-A"},
			[]gem.Option{gem.WithHeadingSpace(gem.HeadingSpaceSingle), gem.WithHorizontalRule("~~~")},
		},
		{
			[]string{"-a", "off"},
			[]gem.Option{gem.WithHeadingLink(gem.HeadingLinkOff), gem.WithHorizontalRule("~~~")},
		},
		{
			[]string{"-a", "below"},
			[]gem.Option{gem.WithHeadingLink(gem.HeadingLinkBelow), gem.WithHorizontalRule("~~~")},
		},
		{
			[]string{"-p", "off"},
			[]gem.Option{gem.WithParagraphLink(gem.ParagraphLinkOff), gem.WithHorizontalRule("~~~")},
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			opts, output, err := ParseArgs("gemgen", tt.args)
			if err != nil {
				t.Errorf("err got %v, want nil", err)
			}
			if output != "" {
				t.Errorf("output got %q, want empty", output)
			}
			want := gem.NewConfig()
			for _, opt := range tt.opts {
				opt.SetConfig(want)
			}
			got := gem.NewConfig()
			for _, opt := range opts.GemOptions {
				opt.SetConfig(got)
			}
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("TestOptions() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
