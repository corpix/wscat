wscp
---------

[![Build Status](https://travis-ci.org/corpix/wscp.svg?branch=master)](https://travis-ci.org/corpix/wscp)

## Usage

```console
$ ./wscp -h
NAME:
   wscp - WebSocket Consumer & Producer
USAGE:
   wscp [global options] command [command options] [arguments...]
VERSION:
   development
COMMANDS:
     help, h  Shows a list of commands or help for one command
GLOBAL OPTIONS:
   --debug        add this flag to enable debug mode
   --help, -h     show help
   --version, -v  print the version
   
$ ./wscp wss://cryptounicorns.io/api/v1/events/stream | head -n 5
{"type":"tickers","payload":{"buy":308.85,"high":318.99,"last":307.61,"low":275.51,"market":"bitfinex","sell":307.77,"symbolPair":"ZEC-USD","tags":["mole","* * * * * *"],"timestamp":1520815645173712100,"vol":11637.10892408}}
{"type":"tickers","payload":{"buy":0.0015303,"high":0.0015553,"last":0.0015273,"low":0.0014944,"market":"bitfinex","sell":0.0015246,"symbolPair":"OMG-BTC","tags":["mole","* * * * * *"],"timestamp":1520815645301125400,"vol":72863.29011334}}
{"type":"tickers","payload":{"buy":0.00063447,"high":0.00065544,"last":0.00063443,"low":0.00063443,"market":"bitfinex","sell":0.00063404,"symbolPair":"EOS-BTC","tags":["mole","* * * * * *"],"timestamp":1520815645308929300,"vol":274006.63606891}}
{"type":"tickers","payload":{"buy":0.000175,"high":0.00018042,"last":0.00017425,"low":0.00016707,"market":"bitfinex","sell":0.0001741,"symbolPair":"SAN-BTC","tags":["mole","* * * * * *"],"timestamp":1520815645311848400,"vol":188912.06904485}}
{"type":"tickers","payload":{"buy":1.6841,"high":1.7347,"last":1.6801,"low":1.4324,"market":"bitfinex","sell":1.6801,"symbolPair":"SAN-USD","tags":["mole","* * * * * *"],"timestamp":1520815645315781000,"vol":1415995.45142326}}
```
