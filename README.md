# goTools For Windows

基于go的windows实用GUI工具集

- Tools/GetIpApp 获取本机IP
- Tools/PlaceToArray 提取指定省市区以及省市区编码转化为三维数组并复制到剪贴板
- Tools/TableToArray 将表格数据提取转化为一维数组并复制到剪贴板
- Tools/TimedShutdown 定时睡眠以及关机
- Tools/continued 待完成工具

# Tools文件夹为脚本工具
只适用于windows系统，内含已打包的可执行二进制文件

如需修改以及查看报错可执行`go build`命令打包

如修改完成不需要命令行显示以及打印输出，可以执行`go build -ldflags="-H windowsgui`进行打包

修改图标先替换相应脚本文件夹的myicon.ico然后执行进到相应文件夹`rsrc -manifest main.manifest -ico myicon.ico -o app.syso
`
最后重新打包即可



# Project supported by JetBrains
***
Many thanks to Jetbrains for kindly providing a license for me to work on this and other open-source projects.
![img.png](img.png)
