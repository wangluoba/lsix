<img src="image.gif" alt="image">

## **概述**：
只需单击即可激活 JetBrains IDE 的单文件工具 - 无需手动输入激活码。

### ✅ Features：
- 支持激活**付费插件**，例如**Rainbow Brackets**
- 自动**备份并恢复修补**前的原始状态
- **跨平台**：兼容 macOS、Linux 和 Windows

### 🔧 v3.1 修复和改进：
- **提高了与其他激活脚本留下的残留配置**的兼容性（例如环境变量和不正确的权限处理）
- 增加了**并行插件请求支持**以加快启动速度
- 引入**插件缓存**，即使在离线或网络状况不佳的情况下也可以使用
- 更改了**文件提取路径**以避免污染当前工作目录
- 增加了对 **`plugin-privacy`** 的支持
- 支持激活**Rainbow Brackets**等插件（某些插件内置了基于时间的许可证检查 - 为避免被标记为异常，建议将到期日期设置为从**今天起 2 年**，而不是过长的期限）

## 如何使用

**只需点击一次，自动激活**

- 点击ide卡自动注入ja-netfilter和激活密钥，重启IDE就可以看到已经激活了。
- 点击插件或IDE卡将激活密钥复制到剪贴板，您可以手动输入密钥进行激活
- 还支持一键移除激活配置
- 单个bin文件，执行后会在当前目录释放激活相关文件，激活完成后IDE会引用该目录，不会删除。
- 警告：某些插件内置了时间检测机制。设置过长的到期时间可能会导致许可证被标记为例外。请考虑将这些插件的到期日期调整为自今日起两年后。
- 警告：如果通过其他激活方式激活时环境变量以_VM_OPTIONS结尾，请参考提示将其移除
- 警告：使用 scoop 安装的软件无法激活，提示“破解卸载失败”。您需要手动创建 %appdata%\JetBrains\IntelliJIdea2025.1 文件夹，然后才能成功激活。其他操作系统也一样。请阅读软件提示，了解目录。


## 开发

安装 go-bindata

```bash
go install github.com/go-bindata/go-bindata/v3/go-bindata@latest
export PATH=$PATH:$(go env GOPATH)/bin
go-bindata --version
go-bindata -o internal/util/access.go -pkg util static/... templates/... cache/...
go run cmd/main.go
```

## 运行它 !

mac linux windows

```
make run
```

## 构建它 !

mac or linux ：

```bash
make run
make build-all
make clean
```

windows use powershell run:

```powershell
.\build.ps1
```

## 关注历史

[![Stargazers over time](https://starchart.cc/saxpjexck/lsix.svg?variant=adaptive)](https://starchart.cc/saxpjexck/lsix)
