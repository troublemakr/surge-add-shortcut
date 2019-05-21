# surge-add-shortcut

日常使用经常遇到访问缓慢的网站，这时就需要添加规则，而每次添加时用 GUI 都觉得繁琐，所以写了个命令行简化工作。

使用方法很简单：

1. 设置环境变量

  - 设置 `SURGE_CONF` 为你的设置文件地址，比如我的为`/Users/myname/Library/Mobile Documents/iCloud~com~nssurge~inc/Documents/macOS.conf`；
  - 设置 `SURGE_PROXY` 为你的代理名称，比如我的就是 `PROXY` 而已。

2. 下载`surge-add`，然后添加到执行目录，如`/usr/local/bin`

3. 执行命令`surge-add example.com`，然后就往配置文件添加 `DOMAIN-SUFFIX,example.com,PROXY` 规则

你还可以基于这命令，添加 LaunchBar 的 Action 或者 Alfred 的 Workflow。

## License
MIT
