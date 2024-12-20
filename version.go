package versionobject

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int

	Prefix string
}

func ParseVersion(v string) (*Version, error) {
	const minVersionParts = 3

	prefix := ""
	if strings.HasPrefix(v, "v") {
		prefix = "v"
		v = strings.TrimPrefix(v, "v")
	}

	parts := strings.Split(v, ".")
	if len(parts) != minVersionParts {
		return nil, errors.New("invalid version")
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid major version: %w", err)
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid minor version: %w", err)
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, fmt.Errorf("invalid patch version: %w", err)
	}

	return &Version{
		Major:  major,
		Minor:  minor,
		Patch:  patch,
		Prefix: prefix,
	}, nil
}

func (v *Version) UpMajor() *Version {
	major := v.Major
	major++

	return &Version{
		Major:  major,
		Minor:  0,
		Patch:  0,
		Prefix: v.Prefix,
	}
}

func (v *Version) UpMinor() *Version {
	minor := v.Minor
	minor++

	return &Version{
		Major:  v.Major,
		Minor:  minor,
		Patch:  0,
		Prefix: v.Prefix,
	}
}

func (v *Version) UpPatch() *Version {
	patch := v.Patch
	patch++

	return &Version{
		Major:  v.Major,
		Minor:  v.Minor,
		Patch:  patch,
		Prefix: v.Prefix,
	}
}

func (v *Version) String() string {
	return fmt.Sprintf("%s%d.%d.%d", v.Prefix, v.Major, v.Minor, v.Patch)
}
