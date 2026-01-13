# GitHub Actions Workflows for Go Multi-Major Modules

This directory contains GitHub Actions workflows for managing tags and releases for the Go multi-major modules in this repository.

## Workflows

### 1. Create Module Tags (`tag-release.yml`)

**Purpose**: Create Git tags for Go module releases with `workflow_dispatch` trigger.

**Trigger**: Manual via GitHub Actions UI

**Inputs**:
- `policy_name`: Choose the module (`add-header`, `cors`, `rate-limit`)
- `version`: Choose the major version (`v1`, `v2`, `v3`)
- `patch_version`: Specify patch version (e.g., `0.0`, `0.1`, `1.0`)
- `dry_run`: Preview what would be tagged without creating tags

**What it does**:
1. Validates the module exists and has a valid `go.mod`
2. Runs tests if they exist
3. Checks if the tag already exists
4. Creates an annotated Git tag with format: `{policy_name}/{version}.{patch_version}`
5. Pushes the tag to the repository

**Example tags created**:
- `add-header/v1.0.0`
- `cors/v2.1.0`
- `rate-limit/v3.0.1`

### 2. List Module Tags (`list-tags.yml`)

**Purpose**: List and validate existing module tags.

**Trigger**: Manual via GitHub Actions UI

**Inputs**:
- `policy_name`: Filter by specific module (optional)
- `action`: Choose action (`list`, `list-latest`, `validate-modules`)

**Actions**:
- `list`: Show all tags or tags for a specific module
- `list-latest`: Show the latest tag for each major version of each module
- `validate-modules`: Validate all modules and their go.mod files

## Usage

### Creating a New Release

1. Go to **Actions** tab in your GitHub repository
2. Select **"Create Module Tags"** workflow
3. Click **"Run workflow"**
4. Fill in the parameters:
   - **Policy name**: Select the module you want to release
   - **Version**: Select the major version
   - **Patch version**: Enter the patch version (e.g., "0.0" for first release, "0.1" for patch)
   - **Dry run**: Check this to preview without creating tags
5. Click **"Run workflow"**

### Example Scenarios

#### First release of add-header v1
- Policy name: `add-header`
- Version: `v1`
- Patch version: `0.0`
- Result: Creates tag `add-header/v1.0.0`

#### Bug fix for rate-limit v2
- Policy name: `rate-limit`
- Version: `v2`
- Patch version: `1.1` (if previous was v2.1.0)
- Result: Creates tag `rate-limit/v2.1.1`

#### Preview without creating tags
- Enable "Dry run" to see what would be created

### Viewing Existing Tags

1. Go to **Actions** tab
2. Select **"List Module Tags"** workflow
3. Choose action:
   - **list**: See all tags
   - **list-latest**: See latest tag per major version
   - **validate-modules**: Check all modules are valid

## Tag Format

Tags follow the format: `{module-name}/{major-version}.{patch-version}`

Examples:
- `add-header/v1.0.0`
- `add-header/v1.0.1`
- `add-header/v2.0.0`
- `cors/v1.0.0`
- `rate-limit/v3.1.0`

## Go Module Usage

After creating tags, users can import specific versions:

```go
// Import specific version
import "github.com/DakshithaS/go-multi-major-modules/add-header/v1"

// Or get specific tagged version
// go get github.com/DakshithaS/go-multi-major-modules/add-header/v1@v1.0.0
```

## Validation

The workflows perform several validations:
- ✅ Module directory exists
- ✅ `go.mod` file exists and is valid
- ✅ Module name in `go.mod` matches expected pattern
- ✅ Tests pass (if present)
- ✅ Tag doesn't already exist
- ✅ Git repository is clean

## Security

The workflows use:
- `contents: write` permission for creating tags
- `contents: read` permission for listing tags
- GitHub Actions bot for commits
- No external dependencies beyond official GitHub actions

## Troubleshooting

### Tag already exists
If a tag already exists, the workflow will fail and show existing tags for the module.

### Module validation fails
Check that:
- The module directory exists
- `go.mod` file is present and valid
- Tests pass (if any exist)
- Module name in `go.mod` matches expected pattern

### Permission denied
Ensure the repository has Actions enabled and the workflow has necessary permissions.