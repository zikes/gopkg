package gopkg

import (
	"testing"

	"github.com/caddyserver/caddy"
)

func TestGopkgConfig(t *testing.T) {
	tests := []struct {
		input     string
		shouldErr bool
		expect    []Config
	}{
		// Single config
		{
			`gopkg /chrisify https://github.com/zikes/chrisify`,
			false,
			[]Config{
				{
					Path: "/chrisify",
					Vcs:  "git",
					Uri:  "https://github.com/zikes/chrisify",
				},
			},
		},
		// Multiple config
		{
			`
			gopkg /chrisify https://github.com/zikes/chrisify
			gopkg /multistatus https://github.com/zikes/multistatus
			`,
			false,
			[]Config{
				{
					Path: "/chrisify",
					Vcs:  "git",
					Uri:  "https://github.com/zikes/chrisify",
				},
				{
					Path: "/multistatus",
					Vcs:  "git",
					Uri:  "https://github.com/zikes/multistatus",
				},
			},
		},
		// Mercurial
		{
			`gopkg /myrepo hg https://bitbucket.org/zikes/myrepo`,
			false,
			[]Config{
				{
					Path: "/myrepo",
					Vcs:  "hg",
					Uri:  "https://bitbucket.org/zikes/myrepo",
				},
			},
		},
	}

	for _, test := range tests {
		c := caddy.NewTestController("http", test.input)
		actual, err := parse(c)
		if !test.shouldErr && err != nil {
			t.Errorf("Unexpected error with %v:\n  %v\n", test.input, err)
		}
		if test.shouldErr && err == nil {
			t.Errorf("Expected error with %v but got none\n", test.input)
		}

		for idx, cfg := range test.expect {
			actualCfg := actual[idx]
			if cfg.Path != actualCfg.Path {
				t.Errorf(
					"Mismatched Path config in %v, expected\n  %v\ngot\n  %v\n",
					test.input,
					cfg.Path,
					actualCfg.Path,
				)
			}
			if cfg.Vcs != actualCfg.Vcs {
				t.Errorf(
					"Mismatched Vcs config in %v, expected\n  %v\ngot\n  %v\n",
					test.input,
					cfg.Vcs,
					actualCfg.Vcs,
				)
			}
			if cfg.Uri != actualCfg.Uri {
				t.Errorf(
					"Mismatched Uri config in %v, expected\n  %v\ngot\n  %v\n",
					test.input,
					cfg.Uri,
					actualCfg.Uri,
				)
			}
		}
	}
}
