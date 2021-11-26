# edgex-snap-metadata
This repo contains the metadata of EdgeX snaps along with utility scripts to generate dummy snaps that can used to update the metadata of those snaps on https://snapcraft.io.

Utility scripts:
* [create-snapcraft](create-snapcraft.go): Reads [template.yaml](template.yaml) and populates it with the [metadata](metadata) of a snap.
* [upload-metadata](upload-metadata.sh): Builds a snap out of the resulting `snap/snapcraft.yaml` file and uploads it to the store, replacing any existing data.

## Requirements
Snapcraft and Go:
```bash
sudo snap install snapcraft go
```

## Usage
Login once:
```
snapcraft login
```

Create the `snapcraft.yaml` file:
```bash
go run create-snapcraft.go --name=edgex-device-gpio
```

Set `--upload-icon` to include the default icon in case it needs to be updated as part of the metadata upload.

Build snap and upload metadata:
```bash
./upload-metadata edgex-device-gpio
```