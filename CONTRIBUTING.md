# Contributing to Go Arcade

Thank you for your interest in contributing! We welcome bug reports, feature suggestions, and pull requests. This document provides guidelines to make the contribution process smooth and enjoyable.

## Code of Conduct

Please read our [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) before participating.

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork locally:**
   ```powershell
   git clone https://github.com/<YOUR-USERNAME>/go-arcade.git
   cd go-arcade
   ```
3. **Add upstream remote:**
   ```powershell
   git remote add upstream https://github.com/<ORIGINAL-OWNER>/go-arcade.git
   ```

## Development Workflow

### Setting Up Your Environment

```powershell
# Download dependencies
go mod tidy

# Run tests to ensure everything works
go test ./... -v

# Format code
go fmt ./...
```

### Creating a Feature Branch

```powershell
git checkout -b feat/your-feature-name
# or for bug fixes:
git checkout -b fix/bug-description
```

Use descriptive branch names:
- `feat/` for new features
- `fix/` for bug fixes
- `chore/` for maintenance tasks
- `docs/` for documentation updates

### Making Changes

1. **Edit files** in your preferred editor
2. **Format code** before committing:
   ```powershell
   go fmt ./...
   ```
3. **Run tests** to ensure no regressions:
   ```powershell
   go test ./... -v
   go vet ./...
   ```
4. **Lint your code** (optional but recommended):
   ```powershell
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
   golangci-lint run ./...
   ```

### Committing Changes

Write clear, concise commit messages using imperative mood:

```
Add player collision detection with enemies
Fix audio initialization on Windows
Update documentation for audio module
```

**Bad commit messages:**
```
Fixed stuff
Update
Added new features
```

If your commit fixes an issue, reference it:
```
Fix crash when no audio device available (fixes #42)
```

### Pushing and Creating a Pull Request

```powershell
# Push your branch
git push origin feat/your-feature-name

# Visit GitHub and create a Pull Request
```

**In your PR description, include:**

- A clear title summarizing the change
- Description of what changed and why
- Reference to related issues (e.g., "Closes #42")
- Screenshots or GIFs for visual changes (if applicable)
- Any testing you performed

**Example PR Description:**
```
## Description
Adds collision detection system for player-enemy interactions.

## Changes
- Implemented bounding box collision checking in `game.go`
- Added collision response (player takes damage)
- Updated tests to cover collision scenarios

## Related Issues
Closes #18

## Testing
- Tested player collision with multiple enemy types
- Verified no false positives with overlapping sprites
```

## Code Style

### Go Conventions

- Follow [Effective Go](https://golang.org/doc/effective_go)
- Use `gofmt` to format code (automatic with `go fmt`)
- Document exported functions and types:
  ```go
  // Player represents the player entity in the game
  type Player struct {
      X, Y float64
  }
  ```

### Naming Conventions

- Use camelCase for variables and functions
- Use PascalCase for exported types and functions
- Use ALL_CAPS for constants
- Use descriptive names (prefer `playerPosition` over `pp`)

### Code Organization

- Keep functions small and focused (ideally < 30 lines)
- Put related functionality in the same file
- Export only necessary functions/types
- Group imports: stdlib, then third-party, then local

## Testing

### Writing Tests

```powershell
# All test files should end with _test.go
# Example: game_test.go
```

Example test:
```go
func TestPlayerMove(t *testing.T) {
    p := NewPlayer()
    p.Move(10, 20)
    if p.X != 10 || p.Y != 20 {
        t.Errorf("Expected (10, 20), got (%v, %v)", p.X, p.Y)
    }
}
```

### Running Tests

```powershell
# Run all tests with verbose output
go test ./... -v

# Run specific package tests
go test ./game -v

# Run with coverage
go test ./... -cover
```

## Documentation

- Update `README.md` for user-facing changes
- Update `CHANGELOG.md` for all notable changes
- Add inline comments for complex logic
- Document public APIs with godoc comments

## Pull Request Review Process

1. **CI checks must pass:** Build, tests, and linting
2. **Code review:** Maintainers will review for quality and fit
3. **Requested changes:** Address feedback and push updates
4. **Approval:** Once approved, a maintainer will merge

## Performance Considerations

- Profile code for performance-critical paths
- Avoid unnecessary allocations in game loops
- Use appropriate data structures (slices vs maps)
- Document performance implications if relevant

## Security

If you discover a security vulnerability, **do not** open a public issue. Instead:

1. Email the maintainers with details
2. Include steps to reproduce (if possible)
3. Allow time for a fix before public disclosure

## Questions or Need Help?

- Check existing [GitHub Issues](https://github.com/<OWNER>/<REPO>/issues) and [Discussions](https://github.com/<OWNER>/<REPO>/discussions)
- Open a new issue to discuss before large changes
- Ask in the PR review if anything is unclear

## Recognition

Contributors will be recognized in:
- Commit history (Git)
- GitHub contributors page
- Potential mention in `CHANGELOG.md` for major contributions

Thank you for contributing to Go Arcade! 🎮
