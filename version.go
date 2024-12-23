package versionobject

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	Major  int
	Minor  int
	Patch  int
	Prefix string

	Namespace string
}

func ParseVersion(v string) (*Version, error) {
	const (
		minParts              = 1
		numberPartsNeedLength = 3
	)

	parts := strings.Split(v, "/")
	if len(parts) < minParts {
		return nil, errors.New("version string too short")
	}

	numberParts := strings.Split(parts[len(parts)-1], ".")
	if len(numberParts) != numberPartsNeedLength {
		return nil, errors.New("invalid version")
	}

	prefix, major, minor, patch, err := parseNumbers(numberParts)
	if err != nil {
		return nil, err
	}

	return &Version{
		Major:     major,
		Minor:     minor,
		Patch:     patch,
		Prefix:    prefix,
		Namespace: strings.Join(parts[0:len(parts)-1], "/"),
	}, nil
}

func (v *Version) UpMajor() *Version {
	major := v.Major
	major++

	return &Version{
		Major:     major,
		Minor:     0,
		Patch:     0,
		Prefix:    v.Prefix,
		Namespace: v.Namespace,
	}
}

func (v *Version) UpMinor() *Version {
	minor := v.Minor
	minor++

	return &Version{
		Major:     v.Major,
		Minor:     minor,
		Patch:     0,
		Prefix:    v.Prefix,
		Namespace: v.Namespace,
	}
}

func (v *Version) UpPatch() *Version {
	patch := v.Patch
	patch++

	return &Version{
		Major:     v.Major,
		Minor:     v.Minor,
		Patch:     patch,
		Prefix:    v.Prefix,
		Namespace: v.Namespace,
	}
}

func (v *Version) String() string {
	namespace := ""
	if v.Namespace != "" {
		namespace = v.Namespace + "/"
	}

	return fmt.Sprintf("%s%s%d.%d.%d", namespace, v.Prefix, v.Major, v.Minor, v.Patch)
}

func parseNumbers(parts []string) (string, int, int, int, error) {
	majorStr := parts[0]
	prefix := ""
	if majorStr[0] == 'v' {
		prefix = "v"
		majorStr = majorStr[1:]
	}

	major, err := strconv.Atoi(majorStr)
	if err != nil {
		return "", 0, 0, 0, fmt.Errorf("invalid major version: %w", err)
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, 0, 0, fmt.Errorf("invalid minor version: %w", err)
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return "", 0, 0, 0, fmt.Errorf("invalid patch version: %w", err)
	}

	return prefix, major, minor, patch, nil
}
