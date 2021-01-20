# go-consul-cleanup                                                 ![Go](https://github.com/ICEBERG98/Go-Consul-Cleanup/workflows/Go/badge.svg?branch=main)
###  **Go Application for cleaning up Failing health checks and Redundant Services in Consul**

I've Written this application to allow cleanup of Services with failing Healtcheks in consul. This application also
 cleans up Redundant Checks and Services that come up in Consul.
 
## Usage
- To build the Binary-
```bash
cd cmd
go build .
```
- To run the Binary
```bash
./cmd -config <configFilePath>
```

## Configuration

- A sample config File looks like this-
```yaml
logging:
  logfile: "/Users/kashishsoni/go/logs/consul_cleanup.log"
  level: "INFO"
  logToStdout: true

defaults:
  port: 8600

bootstrap:
  node:
    address: ""
    port: 8400
  datacenter: ""

```
1. I have tried to keep only the necessary params in config
1. The logging config in the configuration Defines the Defaults for logging
1. The defaults subsection deals with default consul node ports when not specified in the IP
1. The bootstrap subsection defines the bootstrap datacenter and the bootstrap node that will be used to start
 cleanup in the configuration.
    1. Bootstrap Node: This Defines the Ip and Port of the consul node from which the script will fetch the Address
     of all the other nodes to connect to. This will also be the first node the program will connect to.
    1.  Bootstrap Datacenter: This Defines the current Datacenter on which the cleanup is being performed. This
     Cleanup will be performed only on the nodes in this specified Datacenter. This done in order to prevent
      unexpected behaviour.
