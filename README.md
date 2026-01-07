# go-multi-major-modules

This repository demonstrates how to maintain multiple major versions of multiple Go modules in the same Git repository using Go's Semantic Import Versioning (SIV).

## Why /vN in Module Paths?

Go requires the `/vN` suffix in module paths for major versions greater than 1 to enable semantic versioning. This allows different major versions to coexist without conflicts, as each version is treated as a separate module. For example, `github.com/dakshithas/go-multi-major-modules/rate-limit/v2` is distinct from `github.com/dakshithas/go-multi-major-modules/rate-limit/v1`.

## Why Version Folders Instead of Branches or Separate Repositories?

Using version folders in a single repository provides several advantages:

- **Simplicity**: Users can clone one repository and access all versions of all modules.
- **CI/CD**: Easier to run tests and builds for all versions in a single pipeline.
- **Maintenance**: Shared tooling and documentation can be maintained centrally.
- **Visibility**: All versions are visible in one place, making it easier to track changes across versions.

Branches would not work well because Go modules are tied to specific commits/tags, and branches don't provide the clean separation needed for semantic versioning. Separate repositories would increase maintenance overhead and make it harder for users to discover and use different versions.

## When This Approach Is Appropriate

This approach is suitable when:

- You need to support multiple major versions simultaneously for multiple modules.
- The modules share some common structure or tooling.
- You want to minimize repository proliferation.
- The modules are relatively stable, and major versions don't change frequently.

## When This Approach Is Not Appropriate

Avoid this approach when:

- Modules are radically different and share no common code or structure.
- You have many major versions (e.g., v10+), which could make the repository unwieldy.
- You prefer separate repositories for better isolation and access control.
- The modules are in very early stages with frequent breaking changes.

## Modules

- `rate-limit`: Provides rate limiting functionality.
- `add-header`: Provides HTTP header addition functionality.

## Usage

To use a specific version of a module, import it with the appropriate `/vN` suffix:

```go
import "github.com/dakshithas/go-multi-major-modules/rate-limit/v1"
// or
import "github.com/dakshithas/go-multi-major-modules/rate-limit/v2"
// or
import "github.com/dakshithas/go-multi-major-modules/add-header/v3"
```

Each version is independent and can be used separately.