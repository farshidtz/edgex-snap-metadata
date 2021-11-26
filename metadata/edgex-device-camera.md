EdgeX Camera Device Service
---
The EdgeX Camera Device Service is developed to control/communicate ONVIF-compliant
cameras accessible via http in an EdgeX deployment.
Initially the daemon in the snap is disabled - a device profile must be
provisioned externally with core-metadata or provided to device-camera inside
"$SNAP_DATA/config/device-camera/res" before starting.