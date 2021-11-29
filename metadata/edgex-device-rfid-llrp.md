EdgeX Device LLRP Micro Service
---
The EdgeX Device LLRP service for connecting LLRP based RFID devices to EdgeX. Initially the daemon in the snap is disabled - a device profile must be provisioned externally with core-metadata or provided to device-rfid-llrp-go inside "$SNAP_DATA/config/device-rfid-llrp-go/res" before starting.