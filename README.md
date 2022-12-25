# Go-kaohe
后端考核

  
# t5-api文档   
     
## Register 注册

---
### POST /Register

> 请求参数

|    名称    |  位置  |   类型   | 必选  | 说明  |
|:--------:|:----:|:------:|:---:|:---:|
| username | body | string |  是  | 用户名 |
| password | body | string |  是  | 密码  |

> 成功
```json
{
	"msg": "注册成功",
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzMTIxNiIsImlzcyI6IiIsImp0aSI6IjEwMCJ9.b7354eb226beb2c369eae32bacd223b4a601f28158843022157686d81c6e68ba"
}
```

## Login 登录

---
### POST /Login

> 请求参数

|    名称    |  位置  |   类型   | 必选  |       说明       |
|:--------:|:----:|:------:|:---:|:--------------:|
|  token   | body | string |  否  | 不用token就写用户名密码 |
| username | body | string |  否  |      用户名       |
| password | body | string |  否  |       密码       |

> 成功
```json
{
  "msg": "登陆成功"
}
```
```json
{
  "msg": "登陆成功",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzMTIxNiIsImlzcyI6InhpYW90ZTIiLCJqdGkiOiIxMDAifQ==.e607ae5a90f581a58fd29adfd3414940aae6e9e4fdcd838827d7ca011a89f393"
}
```

## AddWarehouse 新增仓库

---
### POST /AddWarehouse

> 请求参数

|  名称   |  位置  |   类型   | 必选  | 说明  |
|:-----:|:----:|:------:|:---:|:---:|
| token | body | string |  是  |     |

> 成功
```json
{
  "msg": "仓库添加成功"
}
```

## AddObject 上架货物

---
### POST /AddObject

> 请求参数

|  名称   |  位置  |   类型   | 必选  |  说明  |
|:-----:|:----:|:------:|:---:|:----:|
| token | body | string |  是  |      |
|  oid  | body | string |  是  | 货物id |
|  wid  | body | string |  是  | 仓库id |

> 成功
```json
{
	"msg": "货物上架成功"
}
```

## PutObject 增加货物数量

---
### POST /PutObject

> 请求参数

|   名称   |  位置  |   类型   | 必选  |  说明  |
|:------:|:----:|:------:|:---:|:----:|
| token  | body | string |  是  |      |
|  oid   | body | string |  是  | 货物id |
| amount | body |  int   |  是  |  数量  |

> 成功
```json
{
	"msg": "货物增加成功"
}
```

## TakeAwayObject 货物出库

---
### POST /TakeAwayObject

> 请求参数

|   名称   |  位置  |   类型   | 必选  |  说明  |
|:------:|:----:|:------:|:---:|:----:|
| token  | body | string |  是  |      |
|  oid   | body | string |  是  | 货物id |
| amount | body |  int   |  是  |  数量  |

> 成功
```json
{
	"msg": "货物出库成功"
}
```

## TakeAwayObject 货物变更

---
### POST /TakeAwayObject

> 请求参数

|  名称   |  位置  |   类型   | 必选  |   说明   |
|:-----:|:----:|:------:|:---:|:------:|
| token | body | string |  是  |        |
|  oid  | body | string |  是  |  货物id  |
|  wid  | body | string |  是  | 变更仓库id |

> 成功
```json
{
	"msg": "货物变更成功"
}
```