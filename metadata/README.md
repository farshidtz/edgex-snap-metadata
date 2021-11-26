# Metadata

This directory contains the metadata of snaps.

The default snap icon is [edgex-snap-icon.png](edgex-snap-icon.png).

Each `.md` file contains the summary and description, separated by a markdown horizontal line (`---`):
```
summary
---
description
```

Notes on the formatting:
* Use `**bold**` for section header
* Section headers do not need a trailing colon
* Use double-spaces at the end of a line to make a line-break
* Use two line-break to start a new paragraph in text or after the section headers
* URLs get converted to hyperlinks automatically. Adding hyperlinks as `[title](url)` work on Github and in the snapcraft.io preview but will not be rendered on the final snapcraft.io listing!
