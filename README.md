# edgex-snap-metadata
This repo contains the metadata of EdgeX snaps along with utility scripts to generate dummy snaps that can used to update the metadata of those snaps in the snap store.

## Usage
Login once:
```
snapcraft login
```

Create the `snapcraft.yaml` file:
```bash
go run create-snapcraft.go --name=edgex-device-gpio
```

Build snap and upload metadata:
```bash
./upload-metadata edgex-device-gpio
```