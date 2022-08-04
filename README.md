# ddns: Dynamic Domain Name Server


<br/><br/>
To bind between dynamic IP and domain.

Usage scenarios:  
When there is no fixed IP in the current network environment, IP changes irregularly, You need to determine the new IP of the current network, and then bind it on platform of ISP (Internet Service Provider).
But this series of operations is cumbersome, And we can't timely perceive that the IP of the current network has changed.

This project was born to solve this problem, it can timely and effectively perceive the changes of IP, and it can automatically bind on the domain name provider platform.

Currently supported ISP: Alibaba Cloud、Tencent Cloud

<br/><br/>
## Document

See the [中文文档](./README_CN.md) for document in Chinese.


<br/><br/>
## Quick Start

### Get Installation File

>##### &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Plan 1：down the installation file
>
> - 1、go into the [page of Tags](https://github.com/jnan806/ddns-golang/tags) in Github 
> - 2、choose the version you needed in the Releases or Tags
> - 3、go to the detail page of the choosed version，find ’ddns-golang.zip‘ in the Assets part, click to download  
>
>##### &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Plan 2：build source code（ GO SDK >=1.16 ）
>
> - 1、download source code
> - 2、go to the root dir of ddns-golang 
> - 3、execute commamd ```go mod tidy``` to download any necessary modules and removes unused modules 
> - 4、execute commamd ```make install``` to package installation file
> - 5、the installation file named 'ddns-golang.zip' will in the root dir


<br/><br/>
### Directory
   Unzip ddns-golang.zip to get the following directory

    - ddns-golang
        |- bin
            |- ddns_386.exe         // windows   386 Executable file
            |- ddns_amd64.exe       // windows amd64 Executable file
            |- ddns_darwin          // macOs   amd64 Executable file
            |- ddns_darwin_arm64    // macOs   arm64 Executable file
            |- ddns_linux_386       // linux     386 Executable file
            |- ddns_linux_amd64     // linux   amd64 Executable file
            |- ddns_linux_arm       // linux     arm Executable file
            |- ddns_linux_arm64     // linux   arm64 Executable file
        |- conf
            |- ddns.conf  // config file


<br/><br/>
### Config Settings
> - The configuration file adopts ini style
>```
>   [sectionName]
>   keyName.1=value1
>   keyName.2=value2
>```
>  | Config-Item   | Description
>  | :---          | :---
>  | [sectionName] | section name, in the whole configuration file **【 cannot be repeated 】**
>  | keyName.1     | config-item in this section
>  | keyName.2     | config-item in this section
>
> - example: Internet Service Provider of domain service ( Abbreviated as 'ISP', ISP **【 can have multiple 】** )
>```
>   [my-aliyun] 
>   ispType=aliyun
>   regionId=cn-hangzhou
>   accessKeyId=accessKeyId
>   accessKeySecret=accessKeySecret
>```
>  | Config-Item     | Description
>  | :---            | :---
>  | [my-aliyun]     | Node name, in the whole configuration file **【 can't be repeated 】**
>  | ispType         | type of ISP: aliyun、tencent is supported. aliyun(Alibaba Cloud) tencent(Tencent Cloud)
>  | regionId        | regionId, provided by in ISP
>  | accessKeyId     | the key to access ISP, provided by ISP
>  | accessKeySecret | the secret to access ISP, provided by ISP
>
> - example: Type 'A' of Record in Domain ( Record Mappings **【 can have multiple 】** )
>```
>   [domainRecord.1]
>   ispId=my-aliyun
>   domain=bbb.com
>   recordMapping=map1:192.168.0.1
>   isUnique=true
>```
>  | Config-Item      | Description
>  | :---             | :---
>  | [domainRecord.1] |  Node name, in the whole configuration file **【 can't be repeated 】**
>  | ispId            |  matched to sectionName of ISP in the above 
>  | domain           |  DomainName
>  | recordMapping    |  IP mapped to the record. In example, 192.168.0.1 is mapped to map1.bbb.com, multiple mappings can be set joined by comma,if IP isn't be given, the Internet IP will be mapped to the record
>  | isUnique         |  is Mapping Unique, when given true，there's one IP mapped to the same record，new IP instead of the old，when given false,several IP can be mapped to the same record，new IP will not instead  of the old



<br/><br/>
### How To Run
Enter the ddns-golang/bin directory and execute the executable files of the corresponding operating system and platform to complete once binding.
```
attention: 
    In consideration of equipment performance, this project does not integrate scheduled tasks, 
    so the executable file is only valid once time.
    
    To realize real-time binding, it is recommended to 
    use it in conjunction with the scheduled task script provided by the computer itself.
```
