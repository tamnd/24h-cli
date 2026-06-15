---
title: "Configuration"
description: "Environment variables, defaults, and the data directory."
weight: 20
---

24h needs almost no configuration: it runs anonymously against public
data out of the box. The settings below let you tune politeness and storage.

## Defaults

| Setting | Default | Flag |
|---|---|---|
| Requests | paced and retried on 429/5xx | `--rate`, `--retries` |
| Per-request timeout | 30s | `--timeout` |
| On-disk cache | under the data directory | `--no-cache` to bypass |

## The data directory

Caches and any record store live under one data directory, chosen in this order:

1. `--data-dir`
2. `24H_DATA_DIR`
3. `$XDG_DATA_HOME/24h`
4. `~/.local/share/24h`

## Environment variables

Every flag has an environment fallback, prefixed `24H_` in
upper case with dashes as underscores. For example:

```bash
export 24H_RATE=1s        # same as --rate 1s
export 24H_DATA_DIR=~/data/24h
```

Flags win over environment variables, which win over the built-in defaults.

## Sending records to a store

`--db` tees every emitted record into a store as a side effect of reading, so a
session fills a local database without a separate import step:

```bash
24h page <path> --db out.db        # SQLite file
24h page <path> --db 'postgres://...'
```
