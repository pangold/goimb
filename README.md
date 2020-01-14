#### GoIM Backend Service

GoIM 的后端服务，用来处理、并分发 GoIM-Front API 服务传递过来的聊天消息，并提供 API

##### 会话（保存到 Redis）

该怎么通过 UserId 找到这个连接所在的节点？以下几个点影响会话的多样性

* 存在多个 GoIM Front 服务（Front 服务集群）
* 存在多种客户端类型：PC/Android/IOS/Web，每个用户都能同时使用每种客户端，即：每个用户最多有4个 token

##### 消息获取（Front GRPC API / Kafka）

通过适配器主动/被动获取被分发到当前节点的消息，涉及 Front 服务集群

* **主动获取消息**：接收 GoIM-Front GRPC API 分发的推送（处理 Front 集群对接逻辑）
* **被动获取消息**：接收 Kakfa 分发的推送（无需处理 Front 集群对接逻辑）

##### 消息处理（历史记录保存到 MongoDB）

管理 middleware 链来处理消息（接收到消息时同步处理），每个消息都包含多个处理方法：

* 处理前 - 消息过滤：Middleware[Before]
* 处理中 - 消息处理并分发：Middleware[对点对、群聊]
* 处理后 - 保存历史消息（异步）：Middleware[After]

> 处理完后，根据 Target 是否在线来进行分发，并保存历史记录，并标记已发送的消息的id（nano second）的位置 -- 比较时间值

##### 消息分发

* **点对点聊天**：分发逻辑1
* **群发消息**：点对点的变种，点对点中 TargetId 为数组','分割
* **群聊**：分发逻辑2

##### 离线消息分发

Session In 触发离线消息推送：查询历史消息[UserId-ClientId-LastDispatchedMessage]，然后发送

> Session In 时传递的参数包括 Token，通过这个 Token 直接发送

##### 系统消息推送

分离到独立的服务？