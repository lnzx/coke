# headscale web ui

简单的 Headscale Web UI

### 支持功能:

- 安装headscale
- 自定义节点IP
- 管理节点
- 管理namespacce
- 

## api

### login

`POST /api/login`

`POST /api/logout`

### node

`GET /api/nodes`

`POST /api/nodes`

`PUT /api/nodes/:id`

`DELETE /api/nodes/:id`

### user namespace

`GET /api/ns`

`POST /api/ns`

`PUT /api/ns`

`DELETE /api/ns`

### 使用第三方库

- https://github.com/bytedance/sonic
- 