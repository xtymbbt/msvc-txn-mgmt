#### 本目录结构：

go-server & myproto-go 两个均用于充当客户端和服务端。

java-client只用于充当客户端，myproto-java既可充当客户端又可充当服务端。

go-server与java-client使用的为同一个simple.proto文件来通信。

myproto-go与myproto-java使用的为同一个HelloWorld.proto文件来通信。

#### 关于GRPC通信需要注意的点：

注意：在使用protoc生成代码的时候，若指定包名，则java端和golang端包名应当相同，否则在grpc通信时，会找不到服务地址。出现unknown service "xxx"的错误。

因此，在生成代码的时候，我们可以选择不指定options和package。

grpc在不同语言间通信的方法：

1.在java方，需预先在pom.xml文件中引入grpc自动生成代码的插件，将编写好的xxx.proto文件放在src/main/proto/目录下，然后，使用mvn compile自动将src/main/proto/目录下的xxx.proto生成代码，然后在target目录下找到generated-sources目录，然后将该目录下的文件拷贝到java客户端的项目目录中。注意包名及其目录位置即可。

2.在golang方，需在项目目录建立proto文件夹，将proto文件放在该目录中，然后再该目录执行

```
protoc --go_out=plugins=grpc:. xxx.proto
```

命令即可。

**注意，该proto文件中，其package需要与java端统一，方可进行grpc通信。**

例如：

java端

```protobuf
package com.bridge.grpc.helloworld;
```

则go端也应为

```protobuf
package com.bridge.grpc.helloworld;
```