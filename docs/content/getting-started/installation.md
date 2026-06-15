---
title: "Installation"
description: "Install 24h from a release, with go install, or from source."
weight: 20
---

## Prebuilt binaries

Every [release](https://github.com/tamnd/24h-cli/releases) carries archives for Linux, macOS,
and Windows on amd64 and arm64, plus deb, rpm, and apk packages for Linux.
Download, unpack, put `24h` on your `PATH`, done. The `checksums.txt`
on each release is signed with keyless [cosign](https://docs.sigstore.dev/) if
you want to verify before running.

## With Go

```bash
go install github.com/tamnd/24h-cli/cmd/24h@latest
```

That puts `24h` in `$(go env GOPATH)/bin`, which is `~/go/bin` unless
you moved it. Make sure that directory is on your `PATH`.

## From source

```bash
git clone https://github.com/tamnd/24h-cli
cd 24h-cli
make build        # produces ./bin/24h
./bin/24h version
```

## Container image

```bash
docker run --rm ghcr.io/tamnd/24h:latest --help
```

## Checking the install

```bash
24h version
```

prints the version and exits.
