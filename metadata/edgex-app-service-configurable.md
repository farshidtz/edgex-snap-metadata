The reference EdgeX App Service Configurable
---
The reference EdgeX app-service-configurable is built using the App Functions SDK. This service is provided as an easy way to get started with processing data flowing through EdgeX. This service leverages the App Functions SDK and provides a way for developers to use configuration instead of having to compile standalone services to utilize built in functions in the SDK. For a full list of supported/built-in functions view the documentation located here:

https://docs.edgexfoundry.org/1.2/microservices/application/ApplServices/

Initially the daemon in the snap is disabled - this allows the configuration to be modified and provided to app-service-config in "$SNAP_DATA/config/app-service-configurable/res" before starting.