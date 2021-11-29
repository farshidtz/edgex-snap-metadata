Connect data MQTT to EdgeX using device-mqtt reference Device Service
---
The official reference EdgeX device-mqtt Device Service built using the device-sdk-go to interact with MQTT brokers. Initially the daemon in the snap is disabled - a device profile must be provisioned externally with core-metadata or provided to device-mqtt inside "$SNAP_DATA/config/device-mqtt/res" before starting.