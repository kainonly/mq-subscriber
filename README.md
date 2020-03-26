# AMQP Subscriber

Configure the broker to subscribe to the AMQP message queue and trigger the microservice to the network callback interface

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kainonly/amqp-subscriber?style=flat-square)](https://github.com/kainonly/amqp-subscriber)
[![Travis](https://img.shields.io/travis/kainonly/amqp-subscriber?style=flat-square)](https://www.travis-ci.org/kainonly/amqp-subscriber)
[![Docker Pulls](https://img.shields.io/docker/pulls/kainonly/amqp-subscriber.svg?style=flat-square)](https://hub.docker.com/r/kainonly/amqp-subscriber)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://raw.githubusercontent.com/kainonly/samqp-subscriber/master/LICENSE)

## Setup

Example using docker compose

```yaml
version: "3.7"
services: 
  subscriber:
    image: kainonly/amqp-subscriber
    restart: always
    volumes:
      - ./subscriber/config:/app/config
      - ./subscriber/log:/app/log
    ports:
      - 6000:6000
      - 6001:6001
```

## Configuration

For configuration, please refer to `config/config.example.yml`

- **debug** `bool` Start debugging, ie `net/http/pprof`, access address is`http://localhost:6060`
- **listen** `string` Microservice listening address
- **amqp** `string` AMQP uri `amqp://guest:guest@localhost:5672/`
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

#### rpc Put (PutParameter) returns (Response) {}

Add or update a subscriber

- PutParameter
  - **identity** `string` subscriber id
  - **queue** `string` consumption queue
  - **url** `string` callback hook url
  - **secret** `string` hook secret
- Response
  - **error** `uint32` error code, `0` is normal
  - **msg** `string` error feedback

```golang
client := pb.NewRouterClient(conn)
response, err := client.Put(
    context.Background(),
    &pb.PutParameter{
        Identity: "a1",
        Queue:    "test",
        Url:      "http://localhost:3000",
        Secret:   "123",
    },
)
```

#### rpc Delete (DeleteParameter) returns (Response) {}

remove subscriber

- DeleteParameter
  - **identity** `string` subscriber id
- Response
  - **error** `uint32` error code, `0` is normal
  - **msg** `string` error feedback

```golang
client := pb.NewRouterClient(conn)
response, err := client.Delete(
    context.Background(),
    &pb.DeleteParameter{
        Identity: "a1",
    },
)
```

#### rpc Get (GetParameter) returns (GetResponse) {}

Get Subscriber Information

- GetParameter
  - **identity** `string` subscriber id
- GetResponse
  - **error** `uint32` error code, `0` is normal
  - **msg** `string` error feedback
  - **data** `Option` result
    - **identity** `string` subscriber id
    - **queue** `string` consumption queue
    - **url** `string` callback hook url
    - **secret** `string` hook secret

```golang
client := pb.NewRouterClient(conn)
response, err := client.Get(
    context.Background(),
    &pb.GetParameter{
        Identity: "a1",
    },
)
```

#### rpc Lists (ListsParameter) returns (ListsResponse) {}

Get subscriber information in batches

- ListsParameter
  - **identity** `string` subscriber IDs
- ListsResponse
  - **error** `uint32` error code, `0` is normal
  - **msg** `string` error feedback
  - **data** `[]Option` result
    - **identity** `string` subscriber id
    - **queue** `string` consumption queue
    - **url** `string` callback hook url
    - **secret** `string` hook secret 

```golang
client := pb.NewRouterClient(conn)
response, err := client.Lists(
    context.Background(),
    &pb.ListsParameter{
        Identity: []string{"a1", "a2"},
    },
)
```

#### rpc All (NoParameter) returns (AllResponse) {}

Get all subscriber IDs

- NoParameter
- AllResponse
  - **error** `uint32` error code, `0` is normal
  - **msg** `string` error feedback
  - **data** `[]string` subscriber IDs

```golang
client := pb.NewRouterClient(conn)
response, err := client.All(
    context.Background(),
    &pb.NoParameter{},
)
```