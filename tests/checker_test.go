package tests

import (
	"github.com/exts/BTagChecker"
	"testing"
)

func TestStringWithInvalidClosingTags(t *testing.T) {
	strs := []string {
		`</blah><p>Links:</p><ul><li><a href="foo">Foo</a></li><li><a href="/bar/baz">BarBaz</a></li></ul>Test`,
		`<p>Links:</p><ul><li><a href="foo">Foo</a></li><li><a href="/bar/baz">BarBaz</a></li></ul><blah>Test`,
		`<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>Test`,
	}

	for _, s := range strs {
		if BTagChecker.HasValidClosingTags(s) != false {
			t.Errorf("String should be invalid, but it's matching as true: [%s]", s)
		}
	}
}

func TestStringWithValidClosingTags(t *testing.T) {
	strs := []string {
		"empty string",
		`<p>Links:</p><ul><li><a href="foo">Foo</a></li><li><a href="/bar/baz">BarBaz</a></li></ul>Test`,
		`<p>Links:</p><ul><li><a href="foo">Foo</a></li><li><a href="/bar/baz">BarBaz</a></li></ul><blah>Test</blah>`,
	}

	for _, s := range strs {
		if BTagChecker.HasValidClosingTags(s) != true {
			t.Errorf("String should be valid, but it's matching as false: [%s]", s)
		}
	}
}