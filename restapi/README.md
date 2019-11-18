# 博客网站 Restful API 设计

## 用户注册

```
POST /api/v1/users
```

### Input

| 名称        | 类型   | 描述          |
| ----------- | ------ | ------------- |
| username    | string | 用户名        |
| nickname    | string | 昵称          |
| password    | string | 密码          |
| captchaId   | string | 图形验证码 id |
| captchaCode | string | 用户识别结果  |

#### Example

```
{
	"username": "huangyuhui",
	"nickname": "huangyuhui",
	"password": "123456",
	"captchaId": "这里是一个 UUID",
	"captchaCode": "auid6A",
	"createdAt": "2019-10-31T00:00:00Z"
}
```

### Response

> Status: 201 Created
>
> Location: /api/v1/users

```
{
	"id": 1,
	"username": "huangyuhui",
	"nickname": "huangyuhui",
	"createdAt": "2019-10-31T00:00:00Z"
}
```

## 用户登录

```
GET /api/v1/users/login
```

### Input

| 名称        | 类型   | 描述          |
| ----------- | ------ | ------------- |
| username    | string | 用户名        |
| password    | string | 密码          |
| captchaId   | string | 图形验证码 id |
| captchaCode | string | 用户识别结果  |

#### Example

```
{
	"username": "huangyuhui",
	"password": "123456",
	"captchaId": "这里是一个 UUID",
	"captchaCode": "auid6A"
}
```

### Response

> Status: 200 OK
>
> Location: /api/v1/users/login

```
{
	"accessToken": "string",
	"expiresAt": "date"
}
```

## 获取图形验证码 id

```
GET /api/v1/users/captcha
```

### Response

> Status: 200 OK
>
> Location: /api/v1/users/captcha

```json
{
    "captchaId": "UUID"
}
```

## 获取图形验证码

```
GET /api/v1/users/captcha/:captchaId
```

### Response

> Status: 200 OK
>
> Location: /api/v1/users/captcha/UUID
>
> Content-Type: application/png

```json
这里是一个图片的二进制
```

## 获取分类列表

```
GET /api/v1/categories
```

### Parameters

| 名称 | 类型   | 描述                                    |
| ---- | ------ | --------------------------------------- |
| sort | string | 排序方式<br />sort=id,DESC 表示 id 逆序 |
| page | number | 页码                                    |
| size | number | 页内项数                                |

### Response

> Status: 200 OK
>
> Location: /api/v1/categories

```json
[
    {
        "id": "0", // 分类 id
        "name": "程序设计", // 分类名称
        "count": 22 // 分类内文章数
    },
    {
        "id": "1",
        "name": "数据结构",
        "count": 7
    }
]
```

## 创建分类

创建新分类，并返回新分类的信息

仅管理员可以创建分类，否则返回 `401 Unauthorized`

```
POST /api/v1/categories
```

### Input

| 名称 | 类型   | 描述     |
| ---- | ------ | -------- |
| name | string | 分类名称 |

#### Example

```json
{
    "name": "程序设计"
}
```

### Response

> Status: 201 Created
>
> Location: /api/v1/categories

```json
{
    "id": 0,
    "name": "程序设计"
}
```

## 更新分类

仅管理员可以修改分类信息，否则返回 `401 Unauthorized`

```
PUT /api/v1/categories/:categoryId
```

### Input

| 名称 | 类型   | 描述     |
| ---- | ------ | -------- |
| name | string | 分类名称 |

#### Example

```json
{
    "name": "程序设计"
}
```

### Response

> Status: 200 OK
>
> Location: /api/v1/categories/0

```json
{
    "id": 0,
    "name": "程序设计"
}
```

## 获取分类内的文章

根据分类 id 获取文章列表

```
GET /api/v1/categories/:categoryId
```

### Parameters

| 名称 | 类型   | 描述                                    |
| ---- | ------ | --------------------------------------- |
| sort | string | 排序方式<br />sort=id,DESC 表示 id 逆序 |
| page | number | 页码                                    |
| size | number | 页内项数                                |

### Response

> Status: 200 OK
>
> Location: /api/v1/categories/0

```json
{
    "category": "程序设计",
    "articles": [
        {
            "id": 0,
            "title": "C 语言的内存管理机制",
            "createdAt": "2019-11-01T22:00:01Z"
        },
        {
            "id": 3,
            "title": "C 语言的预处理器指令",
            "createdAt": "2019-11-18T22:00:01Z"
        }
    ]
}
```

## 创建文章

仅管理员可以创建文章，否则返回 `401 Unauthorized`

```
POST /api/v1/articles
```

### Input

| 名称       | 类型   | 描述                    |
| ---------- | ------ | ----------------------- |
| categoryId | number | 分类 id                 |
| title      | string | 文章标题                |
| content    | string | 文章内容，Markdown 格式 |

#### Example

```json
{
    "categoryId": 0,
    "title": "C 语言的内存管理机制",
    "content": "malloc 就完事了！"
}
```

### Response

> Status: 201 Created
>
> Location: /api/v1/articles

```json
{
    "id": 0,
    "categoryId": 0,
    "title": "C 语言的内存管理机制",
    "content": "malloc 就完事了！"
}
```

## 更新文章

仅管理员可以更新文章，否则返回 `401 Unauthorized`

```
PUT /api/v1/articles/3
```

### Input

| 名称       | 类型   | 描述                    |
| ---------- | ------ | ----------------------- |
| categoryId | number | 分类 id                 |
| title      | string | 文章标题                |
| content    | string | 文章内容，Markdown 格式 |

#### Example

```json
{
    "categoryId": 0,
    "title": "C 语言的预处理器指令:",
    "content": "用得最多的还是 \\#include"
}
```

### Response

> Status: 200 OK
>
> Location: /api/v1/articles/3

```json
{
    "id": 3,
    "categoryId": 0,
    "title": "C 语言的预处理器指令:",
    "content": "用得最多的还是 \\#include"
}
```

## 删除文章

仅管理员可以删除文章，否则返回 `401 Unauthorized`

```
DELETE /api/v1/articles/1
```

### Response

> Status: 200 OK
>
> Location: /api/v1/articles/1

## 获取文章的评论

```
GET /api/v1/articles/:articleId/comments
```