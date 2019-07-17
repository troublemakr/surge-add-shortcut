# surge-add-shortcut

日常使用经常遇到访问缓慢的网站，这时就需要添加规则，而每次添加时用 GUI 都觉得繁琐，所以写了个命令行简化工作。

使用方法很简单：

1. 设置环境变量；

  - 设置 `SURGE_CONF` 为你的设置文件地址，比如我的为`/Users/myname/Library/Mobile Documents/iCloud~com~nssurge~inc/Documents/macOS.conf`；
  - 设置 `SURGE_PROXY` 为你的代理名称或者代理组名称，比如我的就是 `PROXY` 而已。

2. 下载`surge-add`，然后添加到执行目录，如`/usr/local/bin`，然后运行 `chmod +x surge-add` 添加权限；

3. 执行命令`surge-add example.com`，然后就往配置文件添加 `DOMAIN-SUFFIX,example.com,PROXY` 规则。

## Alfred Workflow 使用方法

> as example.com

下载本项目提供的 Alfred Worflow 并安装，然后在设置里打开，修改 shell script（即`SURGE_CONF`，`SURGE_PROXY` 以及 `surge-add` 的路径）。

如果仅在 Alfred 中使用此命令，那么环境变量 `SURGE_CONF` 和 `SURGE_PROXY` 不需要额外设置，Alred Workflow 内已包括这些设置。

## License
MIT
