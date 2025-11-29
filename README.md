# Go Arcade

[![Go Version](https://img.shields.io/badge/Go-1.20%2B-blue)](https://golang.org/doc/go1.20)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)


A lightweight 2D arcade game scaffold written in Go using [Ebiten](https://ebitenengine.org/). No external assets required — all sprites and assets are generated procedurally. Perfect for learning Go game development or as a starting point for arcade-style games.

<br>
<img style="width:100%" src="demo.gif">

## Features

- ✨ Minimal, self-contained codebase
- 🎮 Ready-to-play demo game with keyboard controls
- 🎨 Procedurally generated sprites (no asset files needed)
- 🔊 Audio scaffolding for sound implementation
- 📦 Small, modular project structure
- 🚀 Cross-platform: Windows, macOS, Linux

## Prerequisites

- **Go 1.20+** (recommended: 1.21 or newer)
- Platform-specific requirements for Ebiten (see [Ebiten documentation](https://ebitenengine.org/en/documents/install.html))

## Quick Start

### Run Directly

```powershell
cd "c:\Users\resfa\OneDrive\Desktop\go-arcade"
go mod tidy
go run ./cmd/play
```

### Build Standalone Executable

```powershell
go build -o go-arcade.exe ./cmd/play
.\go-arcade.exe
```

### Cross-Compile for Other Platforms

Linux:
```powershell
$env:GOOS = 'linux'; $env:GOARCH = 'amd64'; go build -o go-arcade-linux ./cmd/play
```

macOS:
```powershell
$env:GOOS = 'darwin'; $env:GOARCH = 'amd64'; go build -o go-arcade-macos ./cmd/play
```

## Game Controls

| Key | Action |
|-----|--------|
| **Arrow Keys** / **WASD** | Move player |
| **Space** | Shoot |

## Project Structure

```
go-arcade/
├── cmd/
│   └── play/          # Main game entry point
│       └── main.go
├── game/              # Core game logic
│   └── game.go
├── audio/             # Audio helpers and stubs
│   └── audio.go
├── internal/          # Internal UI utilities
│   └── ui.go
├── go.mod             # Module definition and dependencies
├── go.sum             # Dependency checksums
├── README.md          # This file
├── CONTRIBUTING.md    # Contribution guidelines
├── CODE_OF_CONDUCT.md # Community code of conduct
├── LICENSE            # MIT License
├── CHANGELOG.md       # Version history
└── .github/
    └── workflows/
        └── ci.yml     # GitHub Actions CI configuration
```


## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

## Code of Conduct

We are committed to providing a welcoming and inspiring community. Please read our [CODE OF CONDUCT](CODE_OF_CONDUCT.md).

## Contributing

Contributions are welcome! Please see [CONTRIBUTING](CONTRIBUTING.md) for guidelines on:

- Reporting bugs
- Suggesting enhancements
- Submitting pull requests
- Commit message conventions

## Changelog

See [CHANGELOG](CHANGELOG.md) for a detailed list of changes in each release.

## Resources

- [Ebiten Official Documentation](https://ebitenengine.org/)
- [Ebiten GitHub Repository](https://github.com/hajimehoshi/ebiten)
- [Go Documentation](https://golang.org/doc/)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

## Troubleshooting

### "module not found" errors

Run `go mod tidy` to download and verify all dependencies.

### Build fails on Linux

Ensure you have the required system packages (see [Ebiten Linux installation](https://ebitenengine.org/en/documents/install.html#Linux)).

### Game runs slowly

- Close unnecessary background applications
- Ensure your GPU drivers are up to date
- Check CPU/memory usage with system monitor

## Future Enhancements

- [ ] Implement procedural audio generation
- [ ] Add particle effects
- [ ] Create level system
- [ ] Add scoring and persistence
- [ ] Implement enemy AI
- [ ] Add animations and tweening
- [ ] Create settings/options menu

## Support

If you encounter issues or have questions:

1. Check existing [GitHub Issues](https://github.com/<OWNER>/<REPO>/issues)
2. Search the [Ebiten community](https://github.com/hajimehoshi/ebiten/discussions)
3. Open a new issue with a clear description and reproduction steps

## Acknowledgments

- Built with [Ebiten](https://ebitenengine.org/) game engine
- Inspired by classic arcade games

---

**Happy coding! 🎮**
