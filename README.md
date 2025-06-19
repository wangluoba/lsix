<img src="image.gif">

## **Overview**:
A single-file tool that activates JetBrains IDEs with one click — no manual activation code input required.

### ✅ Features:
- Supports activation of **paid plugins**, such as **Rainbow Brackets**
- Automatically **backs up and restores** the original state before patching
- **Cross-platform**: compatible with macOS, Linux, and Windows

### 🔧 v3.1 Fixes & Improvements:
- Improved compatibility with **residual configurations** left by other activation scripts (e.g., environment variables and incorrect permission handling)
- Added **parallel plugin request support** to speed up startup
- Introduced **plugin caching**, allowing usage even when offline or under poor network conditions
- Changed the **file extraction path** to avoid polluting the current working directory
- Added support for the **`plugin-privacy`** plugin
- Supports activating plugins like **Rainbow Brackets** (some plugins have built-in time-based license checks — to avoid being flagged as abnormal, it's recommended to set the expiration date to **2 years from today** rather than an excessively long period)

## How 2 Use

**Only once click, automatically activated**

- Click the ide card to automatically inject ja-netfilter and activation key, and restart the IDE to see that it has been activated.
- Click the plugin or ide cards to copy the activation key to the clipboard, and you can manually enter the key for activation
- It also supports one-click removal of activation configuration
- It is single bin file, after execution, the relevant files for activation will be released in the current directory. After activation, the ide will reference the directory and do not delete it.
- warning: Some plugins have built-in time detection mechanisms. Setting an expiration time too long may cause the license to be marked as an exception. Consider adjusting the expiration date of these plugins to two years from today
-  warning: If the environment variable ends with _VM_OPTIONS when activated by other activation methods, please refer to the prompts to remove it
-  warning: The software installed with scoop cannot be activated, and the message "crack failed uninstall" is displayed. You need to manually create the %appdata%\JetBrains\IntelliJIdea2025.1 folder, and then you can successfully activate it. The same is true for other operating systems. Please read the software prompts for the directory.


## dev

install go-bindata

```bash
go install github.com/go-bindata/go-bindata/v3/go-bindata@latest
export PATH=$PATH:$(go env GOPATH)/bin
go-bindata --version
go-bindata -o internal/util/access.go -pkg util static/... templates/... cache/...
go run cmd/main.go
```

## run it !

mac linux windows

```
make run
```

## make it !

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

## Star History

[![Stargazers over time](https://starchart.cc/saxpjexck/lsix.svg?variant=adaptive)](https://starchart.cc/saxpjexck/lsix)
