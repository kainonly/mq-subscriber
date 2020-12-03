# Mq Subscriber

Microservices with automatic message queue consumption and network callback

[![Github Actions](https://img.shields.io/github/workflow/status/kain-lab/mq-subscriber/release?style=flat-square)](https://github.com/kain-lab/mq-subscriber/actions)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kain-lab/mq-subscriber?style=flat-square)](https://github.com/kain-lab/mq-subscriber)
[![Image Size](https://img.shields.io/docker/image-size/kainonly/mq-subscriber?style=flat-square)](https://hub.docker.com/r/kainonly/mq-subscriber)
[![Docker Pulls](https://img.shields.io/docker/pulls/kainonly/mq-subscriber.svg?style=flat-square)](https://hub.docker.com/r/kainonly/mq-subscriber)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://raw.githubusercontent.com/kainonly/mq-subscriber/master/LICENSE)

## Setup

Example using docker compose

```yaml
version: "3.8"
services: 
  subscriber:
    image: kainonly/mq-subscriber
    restart: always
    volumes:
      - ./subscriber/config:/app/config
      - ./subscriber/logs:/app/logs
    ports:
      - 6000:6000
      - 8080:8080
```

## Configuration

For configuration, please refer to `config/config.example.yml`

- **debug** `string` Start debugging, ie `net/http/pprof`, access address is`http://localhost:6060`
- **listen** `string` grpc server listening address
- **gateway** `string` API gateway server listening address
- **queue** `object`
    - **drive** `string` Contains: `amqp`
    - **option** `object` (amqp) 
        - **url** `string` E.g `amqp://guest:guest@localhost:5672/`
- **filelog** `string` file log
- **transfer** `object` [elastic-transfer](https://github.com/kain-lab/elastic-transfer) service
  - **listen** `string` host
  - **pipe** `string` id
    
## Service

The service is based on gRPC to view `api/api.proto`

```proto
syntax = "proto3";
package mq.subscriber;
option go_package = "mq-subscriber/gen/go/mq/subscriber";
import "google/protobuf/empty.proto";
import "google/api/annotations.proto";

service API {
  rpc Get (ID) returns (Option) {
    option (google.api.http) = {
      get: "/subscriber",
    };
  }
  rpc Lists (IDs) returns (Options) {
    option (google.api.http) = {
      post: "/subscribers",
      body: "*"
    };
  }
  rpc All (google.protobuf.Empty) returns (IDs) {
    option (google.api.http) = {
      get: "/subscribers",
    };
  }
  rpc Put (Option) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      put: "/subscriber",
      body: "*",
    };
  }
  rpc Delete (ID) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/subscriber",
    };
  }
}

message ID {
  string id = 1;
}

message IDs {
  repeated string ids = 1;
}

message Option {
  string id = 1;
  string queue = 2;
  string url = 3;
  string secret = 4;
}

message Options {
  repeated Option options = 1;
}
```

## Get (ID) returns (Option)

Get subscriber configuration

### RPC

- **ID**
  - **id** `string` subscriber id
- **Option**
  - **id** `string` subscriber id
  - **queue** `string` consumption queue
  - **url** `string` callback hook url
  - **secret** `string` hook secret

```golang
client := pb.NewAPIClient(conn)
response, err := client.Get(context.Background(), &pb.ID{Id: "debug"})
```

### API Gateway

- **GET** `/subscriber`

```http
GET /subscriber?id=debug HTTP/1.1
Host: localhost:8080
```

## Lists (IDs) returns (Options)

Lists subscriber configuration

### RPC

- **IDs**
  - **ids** `[]string` subscriber IDs
- **Options**
  - **options** `[]Option` result
    - **id** `string` subscriber id
    - **queue** `string` consumption queue
    - **url** `string` callback hook url
    - **secret** `string` hook secret 

```golang
client := pb.NewAPIClient(conn)
response, err := client.Lists(
  context.Background(),
  &pb.IDs{Ids: []string{"debug"}},
)
```

### API Gateway

- **POST** `/subscribers`

```http
POST /subscribers HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "ids": [
        "debug"
    ]
}
```

## All (google.protobuf.Empty) returns (IDs)

Get all subscriber configuration identifiers

### RPC

- **IDs**
  - **ids** `[]string` subscriber IDs

```golang
client := pb.NewAPIClient(conn)
response, err := client.All(context.Background(), &empty.Empty{})
```

### API Gateway

- **GET** `/subscribers`

```http
GET /subscribers HTTP/1.1
Host: localhost:8080
```

## Put (Option) returns (google.protobuf.Empty)

Put subscriber configuration

### RPC

- **Option**
  - **id** `string` subscriber id
  - **queue** `string` consumption queue
  - **url** `string` callback hook url
  - **secret** `string` hook secret

```golang
client := pb.NewAPIClient(conn)
_, err := client.Put(context.Background(), &pb.Option{
  Id:     "debug",
  Queue:  "subscriber.debug",
  Url:    "http://mac:3000/subscriber",
  Secret: "fq7K8EsCMjrv4wOB",
})
```

### API Gateway

- **PUT** `/subscriber`

```http
PUT /subscriber HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "id": "debug",
    "queue": "subscriber.debug",
    "url": "http://mac:3000/subscriber",
    "secret": "fq7K8EsCMjrv4wOB"
}
```

## Delete (ID) returns (google.protobuf.Empty)

Remove subscriber configuration

### RPC

- **ID**
  - **id** `string` subscriber id

```golang
client := pb.NewAPIClient(conn)
_, err := client.Delete(context.Background(), &pb.ID{Id: "debug"})
```

### API Gateway

- **DELETE** `/subscriber`

```http
DELETE /subscriber?id=debug HTTP/1.1
Host: localhost:8080
```
