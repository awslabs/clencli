# Versioning

This document describes how clencli is versioned and released.

## Version scheme

clencli uses semantic versioning in the form `MAJOR.MINOR.PATCH`. The `box/resources/VERSION` file holds the current version, and the `clencli version` command reports it together with the Go version, operating system, and architecture:

```text
clencli v0.3.6 go1.16 linux amd64
```

Increment the version as follows:

- `MAJOR` for incompatible changes to commands, flags, or output.
- `MINOR` for new commands or flags that remain backward compatible.
- `PATCH` for backward-compatible bug fixes.

## Branching

This project follows the [GitFlow](https://nvie.com/posts/a-successful-git-branching-model/) branching model. Develop features on branches and integrate them through the project's main line before tagging a release.

## Release process

Releases are driven by Git tags. The [release workflow](.github/workflows/release.yml) runs on every pushed tag.

1. Update `box/resources/VERSION` to the new version.
2. Update [CHANGELOG.md](CHANGELOG.md) with the changes in the release.
3. Commit the changes and push them.
4. Create and push a tag that matches the version, for example:

   ```bash
   git tag 0.4.0
   git push origin 0.4.0
   ```

When the tag is pushed, the release workflow compiles clencli for each supported operating system and architecture, creates a GitHub release, and uploads the compiled binaries as release assets. The release notes are taken from `CHANGELOG.md`.
