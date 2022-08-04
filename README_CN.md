# <center> ddns-golang: 动态IP绑定域名服务(Golang版本) </center>

<br/><br/>
用于绑定动态IP与域名的服务。

使用场景:  
当前网络环境无固定IP时，IP处于不定期变化的情况下，需要确定当前网络新的IP,而后在域名提供商平台进行绑定。  
这一连串操作比较繁琐，并且我们无法及时感知到当前网络的IP已经变化。

此项目的诞生就是为了解决这一问题，及时有效的感知当前网络IP的变化，并且自动在域名提供商平台进行自动绑定。

当前支持的运营商:Aliyun（阿里云）、Tencent（腾讯云）

<br/><br/>
## 文档

点击 [Document in English](./README.md) 查看英文文档.


<br/><br/>
## 快速开始

### 获取安装包
>##### &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;方式一：直接下载安装包
> 
> - 1、进入 Github 的 Code 页面找到 Tags 的 [超链接](https://github.com/jnan806/ddns-golang/tags) 点击进入
> - 2、在 Releases 或 Tags 下找到想要下载的版本
> - 3、进入版本详情，在 Assets 下找到 ddns-golang.zip，点击下载  
>
>##### &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;方式二：源码编译打包（ GO SDK >=1.16 ）
>
> - 1、下载项目源码
> - 2、进入项目根目录 ddns-golang
> - 3、执行 ```go mod tidy``` 整理并下载相关的第三方包 
> - 4、执行 ```make install``` 编译、打包
> - 5、项目根目录会生成 ddns-golang.zip 包


<br/><br/>
### 目录结构
    解压 ddns-golang.zip，得到如下结构

    - ddns-golang
        |- bin
            |- ddns_386.exe         windows   386平台 可执行文件
            |- ddns_amd64.exe       windows amd64平台 可执行文件
            |- ddns_darwin          macOs   amd64平台 可执行文件
            |- ddns_darwin_arm64    macOs   arm64平台 可执行文件
            |- ddns_linux_386       linux     386平台 可执行文件
            |- ddns_linux_amd64     linux   amd64平台 可执行文件
            |- ddns_linux_arm       linux     arm平台 可执行文件
            |- ddns_linux_arm64     linux   arm64平台 可执行文件
        |- conf
            |- ddns.conf  dddns的配置文件


<br/><br/>
### 配置
> - 配置文件采用 ini 风格
>```
>   [sectionName]
>   keyName.1=value1
>   keyName.2=value2
>```
>  | 配置项         | 说明 
>  | :---          | :---
>  | [sectionName] | 节点名称,在整个配置文件中 **【 不可重复 】**
>  | keyName.1     | sectionName节点下配置项名称
>  | keyName.2     | sectionName节点下配置项名称
> 
> - 配置示例: 域名服务提供商 ( 简称ISP, ISP **【 可以有多个 】** )
>```
>   [my-aliyun] 
>   ispType=aliyun
>   regionId=cn-hangzhou
>   accessKeyId=accessKeyId
>   accessKeySecret=accessKeySecret
>```
>  | 配置项           | 说明 
>  | :---            | :---
>  | [my-aliyun]     | 节点名称,在整个配置文件中 **【 不可重复 】**
>  | ispType         | ISP类型: 目前支持 aliyun(阿里云) tencent(腾讯云)
>  | regionId        | ISP地区, 具体值需在对应 ISP 查询
>  | accessKeyId     | 访问ISP 的 key，由 ISP 提供
>  | accessKeySecret | 访问ISP 的 secret，由 ISP 提供
> 
> - 配置示例: 域名解析A记录配置 ( 域名A记录映射 **【 可以有多个 】** )
>```
>   [domainRecord.1]
>   ispId=my-aliyun
>   domain=bbb.com
>   recordMapping=map1:192.168.0.1
>   isUnique=true
>```
>  | 配置项            | 说明 
>  | :---             | :---
>  | [domainRecord.1] |  节点名称,在整个配置文件中 **【 不可重复 】**
>  | ispId            |  ISP 配置对应的 sectionName
>  | domain           |  主域名
>  | recordMapping    |  记录值映射，示例中表示 map1.bbb.com 对应 192.168.0.1,可为多组以英文逗号分割，如果ip不写则自动获取当前网络外网ip
>  | isUnique         |  记录值是否唯一, true时，同一个记录下只有一个ip，记录重复时用新的ip替换，false时,同一个记录下可存在多个ip，记录重复时不替换

    
<br/><br/>  
### 运行
进入 ddns-golang/bin 目录，直接执行对应操作系统与平台的可执行文件即可完成一次绑定。
```
注: 
    出于对设备性能的考虑，此项目并未集成定时任务，因此可执行文件仅为一次有效，
    如需实现实时绑定，建议配合电脑本身提供的定时任务脚本使用。
```
