# Skland-Task

![golang](https://img.shields.io/github/actions/workflow/status/starudream/skland-task/golang.yml?style=for-the-badge&logo=github&label=golang)
![release](https://img.shields.io/github/v/release/starudream/skland-task?style=for-the-badge)
![license](https://img.shields.io/github/license/starudream/skland-task?style=for-the-badge)

## Config

- `global` [doc](https://github.com/starudream/go-lib/blob/v2/README.md) - [example](https://github.com/starudream/go-lib/blob/v2/app.example.yaml)

以下参数无需手动增加，可通过下方 [Account Manager](#account) 通过发送验证码自动获取

```yaml
accounts:
  - phone: "手机号码"
    hypergryph:
      token: "手机登录后的 token，skland.cred 失效时需要使用该 token 重新登陆"
      code: "授权森空岛后的 code"
    skland:
      cred: "通过上方 hypergryph.code 获取到的凭证"
      token: "通过上方 hypergryph.code 获取到的 token，用于签名，会过期，需要手动 refresh"

cron:
  attend:
    spec: "签到奖励执行时间，默认 0 0 9 * * * 即每天 09:00"
    startup: "是否启动时执行一次，默认 false"
```

## Usage

```
> skland-task -h
Usage of :
  -c, --config string
Usage:
  skland-task [command]

Available Commands:
  account     Manage accounts
  attend      Attend skland
  cron        Run as cron job
  notify      Manage notify

Flags:
  -c, --config string   path to config file
  -h, --help            help for skland-task
  -v, --version         version for skland-task

Use "skland-task [command] --help" for more information about a command.
```

### Account

```shell
# list accounts
skland-task account list
# add account by send phone validate code
skland-task account login
```

### Attend

```shell
# manual attend for input account phone
skland-task attend <account phone>
```

### Notify

```shell
# test notify
skland-task notify send <message>
```

## Docker

```shell
mkdir skland && touch skland/app.yaml
docker run -it --rm -v $(pwd)/skland:/skland -e DEBUG=true starudream/skland-task /skland-task -c /skland/app.yaml account login
docker run -it --rm -v $(pwd)/skland:/skland -e DEBUG=true starudream/skland-task /skland-task -c /skland/app.yaml attend <account phone>
```

## Docker Compose

```yaml
version: "3"
services:
  skland:
    image: starudream/skland-task
    container_name: skland
    restart: always
    command: /skland-task -c /skland/app.yaml cron
    volumes:
      - "./skland/:/skland"
    environment:
      DEBUG: "true"
      app.log.console.level: "info"
      app.log.file.enabled: "true"
      app.log.file.level: "debug"
      app.log.file.filename: "/skland/app.log"
      app.cron.attend.spec: "0 40 8 * * *"
```

## Thanks

- [skland-sign-header](https://gitee.com/FancyCabbage/skyland-auto-sign#sign-header)

## [License](./LICENSE)
