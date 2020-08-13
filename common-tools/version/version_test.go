package version

import (
	"fmt"
	"sort"
	"testing"

	"github.com/hashicorp/go-version"
)

/*
 * @author taohu
 * @date 2020/8/12
 * @time 下午7:42
 * @desc Versioning Library for Go
**/

func TestCompareVersion(t *testing.T) {
	v1, err := version.NewVersion("1.2")
	if err != nil {
		fmt.Println(err)
	}
	v2, err := version.NewVersion("1.5+metadata")
	if err != nil {
		fmt.Println(err)
	}

	if v1.LessThan(v2) {
		fmt.Printf("%v is less than %v\n", v1, v2)
	}
}

func TestVersionRange(t *testing.T) {
	v1, err := version.NewVersion("1.2")
	if err != nil {
		fmt.Println(err)
	}
	// Constraints example.
	constraints, err := version.NewConstraint(">= 1.0, < 1.4")
	if err != nil {
		fmt.Println(err)
	}
	if constraints.Check(v1) {
		fmt.Printf("%s satisfies constraints %s\n", v1, constraints)
	}
}

func TestVersionSort(t *testing.T) {
	versionsRaw := []string{"1.1", "0.7.1", "1.4-beta", "1.4", "2"}
	versions := make([]*version.Version, len(versionsRaw))
	for i, raw := range versionsRaw {
		v, _ := version.NewVersion(raw)
		versions[i] = v
	}

	// After this, the versions are properly sorted
	sort.Sort(version.Collection(versions))
	fmt.Println(versions)
}
