# version-object

version-object - go library for working with versions

## Usage

### Parse Version

```go
package main

import (
	"fmt"
	versionobject "github.com/ci-space/version-object"
)

func main() {
	version, _ := versionobject.ParseVersion("v1.2.3")
	fmt.Println(version.Major)
	fmt.Println(version.Minor)
	fmt.Println(version.Patch)
}
```

### Up major Version

```go
package main

import (
	"fmt"
	versionobject "github.com/ci-space/version-object"
)

func main() {
	version, _ := versionobject.ParseVersion("v1.2.3")
	newVersion := version.UpMajor()
	fmt.Println(newVersion.String()) // v2.0.0
}
```

### Up minor Version

```go
package main

import (
	"fmt"
	versionobject "github.com/ci-space/version-object"
)

func main() {
	version, _ := versionobject.ParseVersion("v1.2.3")
	newVersion := version.UpMinor()
	fmt.Println(newVersion.String()) // v1.3.0
}
```

### Up patch Version

```go
package main

import (
	"fmt"
	versionobject "github.com/ci-space/version-object"
)

func main() {
	version, _ := versionobject.ParseVersion("v1.2.3")
	newVersion := version.UpPatch()
	fmt.Println(newVersion.String()) // v1.3.4
}
```
