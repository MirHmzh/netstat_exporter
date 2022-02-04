# Netstat exporter
Prometheus exporter for exposing reserved ports and it's mapped process.
## System Requirements
- Linux OS (Ubuntu 20.04 & Debian Stretch tested)
- netstat installed
## Quick Start
To quickly scrape reserved ports & it's process
```
./netstat_exporter
```
## Flags
``--port`` Exporter listened ports, default to 9119

## Metrics
``node_port_used{port=80,process="nginx",type="TCP",pid=81} 1.0``
- port is reserved port for specific process
- process is process name that used the port
- type is protocol type of reserved port. Either TCP or UDP
- pid is process ID of reserving port