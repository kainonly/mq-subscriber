# AMQP Subscriber

Configure the broker to subscribe to the AMQP message queue and trigger the microservice to the network callback interface

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kainonly/amqp-subscriber?style=flat-square)](https://github.com/kainonly/amqp-subscriber)
[![Travis](https://img.shields.io/travis/kainonly/amqp-subscriber?style=flat-square)](https://www.travis-ci.org/kainonly/amqp-subscriber)
[![Docker Pulls](https://img.shields.io/docker/pulls/kainonly/amqp-subscriber.svg?style=flat-square)](https://hub.docker.com/r/kainonly/amqp-subscriber)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://raw.githubusercontent.com/kainonly/samqp-subscriber/master/LICENSE)

## Configuration

For configuration, please refer to `config/config.example.yml`

- **debug** `bool` Start debugging, ie `net/http/pprof`, access address is`http://localhost:6060`
- **listen** `string` Microservice listening address
- **amqp** `object` AMQP configuration
    - **host** `string` Connection address
    - **port** `int` port
    - **username** `string` username
    - **password** `string` password
    - **vhost** `string` vhost
- **log** `object` Log configuration
    - **storage** `bool` Turn on local logs
    - **storage_dir** `string` Local log storage directory
    - **socket** `bool` Enable remote log transfer
    - **socket_port** `int` Define the socket listening port
    
## Service

The service is based on gRPC and you can view `router/router.proto`

```
syntax = "proto3";

service Router {
    rpc Put (PutParameter) returns (Response) {
    }

    rpc Delete (DeleteParameter) returns (Response) {
    }

    rpc Get (GetParameter) returns (GetResponse) {
    }

    rpc Lists (ListsParameter) returns (ListsResponse) {
    }

    rpc All (NoParameter) returns (AllResponse) {
    }
}

message NoParameter {
}

message Response {
    uint32 error = 1;
    string msg = 2;
}

message Option {
    string identity = 1;
    string queue = 2;
    string url = 3;
    string secret = 4;
}

message PutParameter {
    string identity = 1;
    string queue = 2;
    string url = 3;
    string secret = 4;
}

message DeleteParameter {
    string identity = 1;
}

message GetParameter {
    string identity = 1;
}

message GetResponse {
    uint32 error = 1;
    string msg = 2;
    Option data = 3;
}

message ListsParameter {
    repeated string identity = 1;
}

message ListsResponse {
    uint32 error = 1;
    string msg = 2;
    repeated Option data = 3;
}

message AllResponse {
    uint32 error = 1;
    string msg = 2;
    repeated string data = 3;
}
```