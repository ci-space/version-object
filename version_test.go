package versionobject_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	versionobject "github.com/ci-space/version-object"
)

func TestObject_UpMajor(t *testing.T) {
	cases := []struct {
		Title    string
		Input    *versionobject.Version
		Expected *versionobject.Version
	}{
		{
			Title: "v1.2.3 -> v2.0.0",
			Input: &versionobject.Version{
				Major:  1,
				Minor:  2,
				Patch:  3,
				Prefix: "v",
			},
			Expected: &versionobject.Version{
				Major:  2,
				Minor:  0,
				Patch:  0,
				Prefix: "v",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Title, func(t *testing.T) {
			got := c.Input.UpMajor()

			assert.Equal(t, c.Expected, got)
		})
	}
}

func TestObject_UpMinor(t *testing.T) {
	cases := []struct {
		Title    string
		Input    *versionobject.Version
		Expected *versionobject.Version
	}{
		{
			Title: "v1.2.3 -> v1.3.0",
			Input: &versionobject.Version{
				Major:  1,
				Minor:  2,
				Patch:  3,
				Prefix: "v",
			},
			Expected: &versionobject.Version{
				Major:  1,
				Minor:  3,
				Patch:  0,
				Prefix: "v",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Title, func(t *testing.T) {
			got := c.Input.UpMinor()

			assert.Equal(t, c.Expected, got)
		})
	}
}

func TestObject_UpPatch(t *testing.T) {
	cases := []struct {
		Title    string
		Input    *versionobject.Version
		Expected *versionobject.Version
	}{
		{
			Title: "v1.2.3 -> v1.2.4",
			Input: &versionobject.Version{
				Major:  1,
				Minor:  2,
				Patch:  3,
				Prefix: "v",
			},
			Expected: &versionobject.Version{
				Major:  1,
				Minor:  2,
				Patch:  4,
				Prefix: "v",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Title, func(t *testing.T) {
			got := c.Input.UpPatch()

			assert.Equal(t, c.Expected, got)
		})
	}
}

func TestParseVersion(t *testing.T) {
	cases := []struct {
		Title    string
		Input    string
		Expected *versionobject.Version
	}{
		{
			Title: "v1.2.3",
			Input: "v1.2.3",
			Expected: &versionobject.Version{
				Major:  1,
				Minor:  2,
				Patch:  3,
				Prefix: "v",
			},
		},
		{
			Title: "1.2.3",
			Input: "1.2.3",
			Expected: &versionobject.Version{
				Major:     1,
				Minor:     2,
				Patch:     3,
				Namespace: "",
			},
		},
		{
			Title: "pkg/github/v0.1.0",
			Input: "pkg/github/v0.1.0",
			Expected: &versionobject.Version{
				Major:     0,
				Minor:     1,
				Patch:     0,
				Prefix:    "v",
				Namespace: "pkg/github",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Title, func(t *testing.T) {
			got, err := versionobject.ParseVersion(c.Input)
			assert.NoError(t, err)
			assert.Equal(t, c.Expected, got)
		})
	}
}
