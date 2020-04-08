Exploring Go Modules
====================

Helpful commands:
- `go help mod`
- `go list -m all`
- `go list -m -versions <module>`
- `go mod verify`
- `go mod tidy`


Semantic Versioning
-------------------

[https://semver.org](https://semver.org)

ex. v1.5.3-pre1

- __V__: Version prefix (required)
- __1__: Major revision (likely to break backward comatibility (BC))
- __5__: Minor revision (new features, doesn't break BC)
- __3__: Patch (bug fixes, no new features, and doesn't break BC)
- __pre1__: Pre-release of new version, if applicable (text is arbitrary)


Module Versioning Rules
-----------------------

__v1 and earlier__
- No promise of backward compatibility prior to v1.0.0
- Precendence is determined by major, minor, then patch versions

__v2+__
- Backward compatibility should be preserved within a major version
- Each major version has unique import path
- Tags or vX directory

__Unversioned commits__
- Still uses semver
- Prerelease identifier (timestamp and commit hash)


Module Queries
--------------
- Specific version: @v1.7.2
- Version prefix: @v1
- Latest: @latest (release or tag)
- Specific commit: @c956192
- Specific branch: @master
- Comparsion: @>=1.7.2

__Comparison Rules__
- Closet match wins
- Prereleases have lower precendence


Advanced Module Management Tools
--------------------------------
- `go mod why`
- `go mod graph`
- `go mod edit`
- `go mod download`
- `go mod vendor`
- `go build -mod=readonly .`
