# Go Utilities Library

A modular Go utilities repository designed for squad autonomy and organic standardization. This repository provides both company-wide standard utilities and squad-specific commons modules.

## 🏗️ Repository Structure

```
go-utils/
├── libman                     # Library management tool
├── logger/                    # Standard module: Company-wide logging utilities
│   ├── go.mod
│   └── logger.go
└── commons/                   # Squad-owned utilities
    ├── esign-core/           # eSign core squad utilities
    │   └── stringutils/
    │       ├── go.mod
    │       ├── stringutils.go
    │       └── stringutils_test.go
    └── talenta-tm/           # Talenta time management squad utilities
        └── mathutils/
            ├── go.mod
            ├── mathutils.go
            └── mathutils_test.go
```

## 📦 Module Tiers

### Standard Tier (Company-wide)
- **Location**: Root directory
- **Approval**: Company-wide consensus required
- **Scope**: Universal patterns (config, logger, database)
- **Example**: `go get bitbucket.org/mid-kelola-indonesia/go-utils/logger`

### Squad-Commons Tier (Squad-owned)
- **Location**: `commons/<squad-name>/<module-name>/`
- **Approval**: Squad autonomy only
- **Scope**: Squad-specific business utilities
- **Example**: `go get bitbucket.org/mid-kelola-indonesia/go-utils/commons/talenta-tm/mathutils`

## 🚀 Quick Start

### Using Existing Modules

```bash
# Use a standard module
go get bitbucket.org/mid-kelola-indonesia/go-utils/logger

# Use a squad commons module
go get bitbucket.org/mid-kelola-indonesia/go-utils/commons/esign-core/stringutils@v1.0.0
```

### Creating a New Commons Module

1. **Create your module directory**:
   ```bash
   mkdir -p commons/<your-squad>/<module-name>
   cd commons/<your-squad>/<module-name>
   ```

2. **Initialize the Go module**:
   ```bash
   go mod init bitbucket.org/mid-kelola-indonesia/go-utils/commons/<your-squad>/<module-name>
   ```

3. **Implement your utilities**:
   ```go
   // your-utility.go
   package yourmodule

   func YourFunction() string {
       return "Hello from your squad!"
   }
   ```

4. **Add tests**:
   ```go
   // your-utility_test.go
   package yourmodule

   import "testing"

   func TestYourFunction(t *testing.T) {
       result := YourFunction()
       if result != "Hello from your squad!" {
           t.Errorf("Expected 'Hello from your squad!', got %s", result)
       }
   }
   ```

5. **Version and release** (see [Versioning](#-versioning) section below)


## 🏷️ Versioning

### Version Format
All versions MUST follow semantic versioning: `v<MAJOR>.<MINOR>.<PATCH>` (e.g., `v1.0.0`)

### Git Tagging Schema

**Standard modules**:
```bash
git tag logger/v1.0.0
git tag logger/v1.1.0
```

**Commons modules**:
```bash
git tag commons/talenta-tm/mathutils/v1.0.0
git tag commons/esign-core/stringutils/v2.1.0
```

### Release Workflow

1. **Develop and test your changes**
2. **Use libman to create a version tag**:
   ```bash
   ./libman tag commons/your-squad/your-module v1.0.0
   ```
3. **Publish the tag**:
   ```bash
   ./libman publish
   ```
4. **Consumers can now use the new version**:
   ```bash
   go get bitbucket.org/mid-kelola-indonesia/go-utils/commons/your-squad/your-module@v1.0.0
   ```

## 🏛️ Governance Model

### Standard Tier Governance
- **Decision Making**: Company-wide consensus required
- **Review Process**: Technical committee review for universal applicability
- **Use Cases**: Patterns used across multiple squads (logging, config, database)
- **Responsibility**: Shared ownership with change impact assessment

### Squad-Commons Tier Governance
- **Decision Making**: Squad autonomy - no external approval needed
- **Review Process**: Internal squad review only
- **Use Cases**: Squad-specific business logic and utilities
- **Responsibility**: Full squad ownership and maintenance

### Module Promotion Path

Commons modules can be promoted to Standard tier through this process:

1. **Recognition**: Module gains adoption across multiple squads
2. **Proposal**: Squad proposes promotion with usage evidence
3. **Review**: Technical committee evaluates for universal applicability
4. **Consensus**: Company-wide agreement on standardization
5. **Migration**: Gradual migration from commons to standard location


## 🛠️ Library Management with `libman`

The repository includes a `libman` Ruby script to help manage modules and versions.

### Available Commands

#### List All Modules
```bash
# List all available modules
$ ./libman list
Found 3 module(s):
[standard] logger
commons/esign-core/stringutils
commons/talenta-tm/mathutils
```

#### List Module Versions
```bash
# List versions for a standard module
$ ./libman versions logger

# List versions for a commons module
$ ./libman versions commons/talenta-tm/mathutils
```

#### Create Version Tags
```bash
# Tag a standard module
$ ./libman tag logger v1.2.0

# Tag a commons module
$ ./libman tag commons/esign-core/stringutils v2.1.0
```

#### Publish All Tags
```bash
# Pushes all local tags to the remote repository.
$ ./libman publish
```
