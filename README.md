# edgex-snap-metadata
This repo contains the metadata of EdgeX snaps along with utility scripts to generate appstream metadata files.

The output files along with the EdgeX snaps icon are stored in the [appstream] branch.

## Metadata

The [metadata](metadata) directory contains the metadata of snaps.

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


## Usage

Generate AppStream file from markdown:
```bash
go run generate-appstream.go --input=metadata/edgex-device-gpio.md
```

## CI
Markdown files in main get converted automatically and added to the [appstream] branch.

See [.github/workflows/appstream.yaml](.github/workflows/appstream.yaml)

[appstream]: https://github.com/canonical/edgex-snap-metadata/tree/appstream