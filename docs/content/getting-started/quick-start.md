---
title: "Quick start"
description: "Fetch your first record with 24h."
weight: 30
---

Once `24h` is on your `PATH`, fetch a page. The argument is the path
of the page on 24h.com (everything after the host), or a full URL:

```bash
24h page <path>
```

By default you get an aligned table. Ask for JSON when you want to pipe it:

```bash
$ 24h page <path> -o json
[
  {
    "id": "<path>",
    "url": "https://24h.com/<path>",
    "title": "<path>",
    "body": "..."
  }
]
```

## Shape the output

The same flags work on every command:

```bash
24h page <path> --fields id,url        # keep only these columns
24h page <path> --template '{{.Body}}' # just the body text
24h page <path> -o jsonl | jq .url     # one object per line, into jq
```

`-o` takes `table`, `json`, `jsonl`, `csv`, `tsv`, `url`, or `raw`. Left to
`auto`, it prints a table to a terminal and JSONL into a pipe, so the same
command reads well by hand and parses cleanly downstream. See
[output formats](/reference/output/) for the full contract.

## Follow the links

`links` lists the pages a page links to, and each one is a path you can fetch in
turn:

```bash
24h links <path> -n 10                 # the first ten links
24h links <path> -o url                # just the URLs
24h links <path> -o url | head -3 | xargs -n1 24h page
```

## Serve it instead

The same operations are available over HTTP and to agents over MCP:

```bash
24h serve --addr :7777 &
curl -s 'localhost:7777/v1/page/<path>'          # NDJSON, one record per line
24h mcp                                # MCP over stdio: page, links
```

## What to build next

This scaffold ships one example type, `page`, wired end to end so the whole
chain works today. To make it really about 24h, model the records you
care about in `24h/` and declare their operations in
`24h/domain.go`. Each one you add shows up as a command here, a route
under `serve`, and a tool under `mcp`, with no extra wiring. The
[guides](/guides/) cover the common jobs.
