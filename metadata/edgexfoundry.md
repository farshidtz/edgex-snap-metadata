Open-source framework for IoT edge computing
---
EdgeX Foundry is a vendor-neutral, highly flexible and scalable open-source framework hosted by The Linux Foundation. It enables developers to build apps that run at the edge, acting as a middleware between the things and the cloud.

It provides the components to develop microservices responsible for data acquisition, data analytics and cloud connectors, exposing a rich set of APIs to allow full control of the system and configuration. The reference architecture allows scaling out to thousands of devices and sensors.

This snap is a part of EdgeX Foundry and contains the following reference services:

**Core Services**
* Core Data and Message Bus (Redis)
* Core Command
* Core Metadata
* Registry and Config (Consul)

**Supporting Services**
* Rules Engine (eKuiper)
* Scheduling
* Alerts & Notifications

**Application Services**
* Configurable Application Service

**Security Services**
* API Gateway a.k.a. Reverse Proxy (Kong)
* Secret Store (Vault)

**Management Services**
* Management Service Agent (deprecated)

**Device Services**
* Virtual

Note that not all the above services are enabled and started by default.

=============

**Further Reading**

Getting started with snaps (v2.1)  
https://docs.edgexfoundry.org/2.1/getting-started/Ch-GettingStartedSnapUsers

Snap usage instructions  
https://github.com/edgexfoundry/edgex-go/blob/main/snap/README.md

EdgeX documentation  
https://docs.edgexfoundry.org

Source code  
https://github.com/edgexfoundry/edgex-go

====================

**Development Tool Snaps**
* EdgeX UI: https://snapcraft.io/edgex-ui
* EdgeX CLI: https://snapcraft.io/edgex-cli

**Other Supported EdgeX Snaps**
* App Service Configurable: https://snapcraft.io/edgex-app-service-configurable
* Device Camera: https://snapcraft.io/edgex-device-camera
* Device Modbus: https://snapcraft.io/edgex-device-modbus
* Device MQTT: https://snapcraft.io/edgex-device-mqtt
* Device REST: https://snapcraft.io/edgex-device-rest
* Device SNMP: https://snapcraft.io/edgex-device-snmp
* Device Grove: https://snapcraft.io/edgex-device-grove

**Additional References**
* Redis: https://github.com/redis/redis
* Consul: https://github.com/hashicorp/consul
* eKuiper: https://github.com/lf-edge/ekuiper
* Kong: https://github.com/Kong/kong
* Vault: https://github.com/hashicorp/vault