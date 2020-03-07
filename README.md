# AMQP Subscriber

配置代理订阅 AMQP 消息队列并触发至网络回调接口的微服务

## 配置

配置请参考 `config/config.example.yml` 并创建 `config/config.yml`

- **debug** `bool` 开启调试，即 `net/http/pprof`，访问地址为 `http://localhost:6060` 
- **listen** `string` 微服务监听地址
- **amqp** `object` AMQP 配置
    - **host** `string` 连接地址
    - **port** `int` 端口
    - **username** `string` 用户名
    - **password** `string` 密码
    - **vhost** `string` 虚拟空间
- **log** `object` 日志配置
    - **storage** `bool` 开启本地日志
    - **storage_dir** `string` 本地日志存储目录
    - **socket** `bool` 开启日志远程传输
    - **socket_port** `int` 定义socket监听端口
    
## 服务

服务基于 gRPC 可查看 `router/router.proto`

```
syntax = "proto3";

message NoParameter {
}

message Response {
    uint32 error = 1;
    string msg = 2;
}

message AllResponse {
    uint32 error = 1;
    repeated string data = 2;
}

message Option {
    string identity = 1;
    string queue = 2;
    string url = 3;
    string secret = 4;
}

message GetParameter {
    string identity = 1;
}

message GetResponse {
    uint32 error = 1;
    Option data = 2;
}

message ListsParameter {
    repeated string identity = 1;
}

message ListsResponse {
    uint32 error = 1;
    repeated Option data = 2;
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

service Router {
    rpc All (NoParameter) returns (AllResponse) {
    }

    rpc Get (GetParameter) returns (GetResponse) {
    }

    rpc Lists (ListsParameter) returns (ListsResponse) {
    }

    rpc Put (PutParameter) returns (Response) {
    }

    rpc Delete (DeleteParameter) returns (Response) {
    }
}
```