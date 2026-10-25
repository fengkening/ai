# ai

一个"推 tag 就自动编译发版"的模板：代码在手机上写，编译在 GitHub Actions 上跑，
产物直接从 release 下载回手机。

## 用法

```bash
# 1. 改代码（main.go 换成你自己的）
git add -A && git commit -m "你的改动"
git push

# 2. 打标签 = 触发发版
git tag v0.1.1
git push origin v0.1.1

# 3. 等 Actions 跑完，把产物拉回手机
gh release download v0.1.1 -R fengkening/ai -D ./bin
chmod +x ./bin/hello-linux-arm64
./bin/hello-linux-arm64
```

## workflow 做了什么

`.github/workflows/release.yml`，推 `v*` 标签时：

1. 矩阵编译 `linux/arm64` 和 `linux/amd64`
2. 静态链接（`CGO_ENABLED=0`）、`-trimpath`、去符号表
3. 生成 `SHA256SUMS.txt`
4. 自动创建 release 并把产物作为附件上传

也支持 `workflow_dispatch` 手动触发，此时只编译、不发版。

## 为什么是 Go

手机上的容器没有编译器，本地编不了任何东西。Go 交叉编译不需要额外工具链，
`GOOS=linux GOARCH=arm64` 一行就能出静态二进制，在 Android 容器里直接能跑。

## 注意

本仓库的 `.git` 是一个指针文件（`gitdir: /root/work/ai.git`），因为 `/sdcard`
是 FUSE 挂载、不支持硬链接，git 无法在那里存放对象库。工作区文件在 `/sdcard`，
对象库在 `/root/work`。搬运请用 `git bundle` 或重新 clone。
