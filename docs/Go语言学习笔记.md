# 简介

> 官网地址：<https://golang.google.cn/>
>
> Go（又称Golang）是Google开发的一种`静态强类型`、`编译型`、`并发型`，并具有`垃圾回收功能`的编程语言。
>
> > 静态可以简单理解为在编译时就确定了具体类型，而不能在运行时动态绑定。
>
> Go的语法接近`C语言`，所以很多概念与C语言保持一致。但GO又是一门十分年轻的语言，因此借鉴了很多C语言的相关设计经验，简化了很多。但对于变量的声明有所不同,有人说GO语言是C语言的网络版。
>
> Go支持垃圾回收功能。
>
> Go即是面向过程语言也是面向对象语言(通过Struct结构实现)。

## 历史版本

**2007年**，[谷歌](https://baike.baidu.com/item/%E8%B0%B7%E6%AD%8C)工程师Rob Pike, Ken Thompson和Robert Griesemer开始设计一门全新的语言，这是Go语言的最初原型。

**2009**年11月10日，Go语言以开放源代码的方式向全球发布。 [1] 

**2011**年3月16日，Go语言的第一个稳定(stable)版本r56发布。 [2] 

**2012**年3月28日，Go语言的第一个正式版本Go1发布。 [2] 

**2013**年4月04日，Go语言的第一个Go 1.1beta1测试版发布。 [3] 

2013年4月08日，Go语言的第二个Go 1.1beta2测试版发布。 [3] 

2013年5月02日，Go语言Go 1.1RC1版发布。 [4] 

2013年5月07日，Go语言Go 1.1RC2版发布。 [5] 

2013年5月09日，Go语言Go 1.1RC3版发布。 [6] 

2013年5月13日，Go语言Go 1.1正式版发布。

2013年9月20日，Go语言Go 1.2RC1版发布。 [7] 

2013年12月1日，Go语言Go 1.2正式版发布。 [8] 

**2014**年6月18日，Go语言Go 1.3版发布。 [9] 

**2014**年12月10日，Go语言Go 1.4版发布。 [10] 

**2015**年8月19日，Go语言Go 1.5版发布，本次更新中移除了”最后残余的C代码”。 [11] 

2016年2月17日，Go语言Go 1.6版发布。 [12] 

2016年8月15日，Go语言Go 1.7版发布。 [13] 

2017年2月17日，Go语言Go 1.8版发布。 [14] 

**2017**年8月24日，Go语言Go 1.9版发布。 [15] 

**2018**年2月16日，Go语言Go 1.10版发布。 [16]

2018年x月x日，Go语言Go 1.11版发布。 [17]

**2019**年x月x日，Go语言Go 1.12版发布。 [18]

2019年x月x日，Go语言Go 1.13版发布。 [19]

## Goland的环境搭建

https://blog.csdn.net/guidao13/article/details/81839771

### 下载安装

> <https://golang.google.cn/dl/>

### 配置IDEA

不官是IDEA还是其它环境，GO语言最重的参数只有两个，GOROOT和GOPATH。其中GOROOT是指你的GO SDK安装路径，这个可以类型JavaJDK。

 GO语言的项目管理与Java不同，没有Maven等类似工具插件，而是在语言层面规定了一般项目应该有的结构。首先，最关键一的是通过环境变量GOPATH指定了工作空间。注意：此变量可以分为全局的和局部的。

#### GOROOT

![1](https://github.com/roboslyq/go_study_proj/tree/master/docs/images/1.jpg)

#### GOPATH

![2](docs/images/2.jpg)

##　GO语言的标准目录　

# GO语言标准库

| Go语言标准包名 | 功能                                                         |                        |
| :------------- | :----------------------------------------------------------- | ---------------------- |
| bufio          | 带缓冲的IO操作                                               | 所有语言通用           |
| bytes          | 字节码操作                                                   | 所有语言通用           |
| container      | 容器（封装了堆，队列，环形列表等容器）                       | 与Java的Collection类似 |
| crypto         | 加密算法                                                     | Java的一个工具包       |
| database       | 数据库驱动                                                   | Java的JDBC             |
| debug          | 各种调试文件格式访问及调试功能                               |                        |
| encoding       | 编码解码，如常见的JSON,XML,Base64等                          | Java的各种分散包       |
| flag           | 命令行解析                                                   |                        |
| fmt            | 格式化                                                       |                        |
| go             | Go语言的词法，语法树，类型等。可通过这个包进行代码信息提取和修改 |                        |
| html           | HTML转义及模板系统                                           |                        |
| image          | 图片操作                                                     |                        |
| io             | I/O接口的                                                    |                        |
| math           | 数学算法库                                                   |                        |
| net            | 网络库，支持Socket，HTTP,邮件，RPC,SMTP等                    |                        |
| os             | 操作系统平台不依赖于平台封装                                 |                        |
| path           | 兼容各操作系统路径操作实用函数                               |                        |
| plugin         | Go1.7 加入插件系统，支持将代码插件，按需加载                 |                        |
| reflect        | 反射，可以动态获取代码中的类型信息，获取和修改变量值         |                        |
| regexp         | 正常表达式                                                   |                        |
| runtime        | 运行时接口                                                   |                        |
| sort           | 排序接口                                                     |                        |
| string         | 字符串接口：转换，解析及其它常用函数                         |                        |
| time           | 时间接口                                                     |                        |
| text           | 文本模板及Token词法器等                                      |                        |
|                |                                                              |                        |
|                |                                                              |                        |
|                |                                                              |                        |
|                |                                                              |                        |

## 标准库



| 名称                                                         |      | 摘要                                                         |
| :----------------------------------------------------------- | :--- | :----------------------------------------------------------- |
| [archive](https://studygolang.com/static/pkgdoc/main.html#archive) |      |                                                              |
| [tar](https://studygolang.com/static/pkgdoc/pkg/archive_tar.htm) |      | tar包实现了tar格式压缩文件的存取.                            |
| [zip](https://studygolang.com/static/pkgdoc/pkg/archive_zip.htm) |      | zip包提供了zip档案文件的读写服务.                            |
| [bufio](https://studygolang.com/static/pkgdoc/pkg/bufio.htm) |      | bufio 包实现了带缓存的I/O操作.                               |
| [builtin](https://studygolang.com/static/pkgdoc/pkg/builtin.htm) |      | builtin 包为Go的预声明标识符提供了文档.                      |
| [bytes](https://studygolang.com/static/pkgdoc/pkg/bytes.htm) |      | bytes包实现了操作[]byte的常用函数.                           |
| [compress](https://studygolang.com/static/pkgdoc/main.html#compress) |      |                                                              |
| [bzip2](https://studygolang.com/static/pkgdoc/pkg/compress_bzip2.htm) |      | bzip2包实现bzip2的解压缩.                                    |
| [flate](https://studygolang.com/static/pkgdoc/pkg/compress_flate.htm) |      | flate包实现了deflate压缩数据格式，参见RFC 1951.              |
| [gzip](https://studygolang.com/static/pkgdoc/pkg/compress_gzip.htm) |      | gzip包实现了gzip格式压缩文件的读写，参见RFC 1952.            |
| [lzw](https://studygolang.com/static/pkgdoc/pkg/compress_lzw.htm) |      | lzw包实现了Lempel-Ziv-Welch数据压缩格式，这是一种T. A. Welch在“A Technique for High-Performance Data Compression”一文（Computer, 17(6) (June 1984), pp 8-19）提出的一种压缩格式. |
| [zlib](https://studygolang.com/static/pkgdoc/pkg/compress_zlib.htm) |      | zlib包实现了对zlib格式压缩数据的读写，参见RFC 1950.          |
| [container](https://studygolang.com/static/pkgdoc/main.html#container) |      |                                                              |
| [heap](https://studygolang.com/static/pkgdoc/pkg/container_heap.htm) |      | heap包提供了对任意类型（实现了heap.Interface接口）的堆操作.  |
| [list](https://studygolang.com/static/pkgdoc/pkg/container_list.htm) |      | list包实现了双向链表.                                        |
| [ring](https://studygolang.com/static/pkgdoc/pkg/container_ring.htm) |      | ring实现了环形链表的操作.                                    |
| [context](https://studygolang.com/static/pkgdoc/pkg/context.htm) |      | Package context defines the Context type, which carries deadlines, cancelation signals, and other request-scoped values across API boundaries and between processes. |
| [crypto](https://studygolang.com/static/pkgdoc/pkg/crypto.htm) |      | crypto包搜集了常用的密码（算法）常量.                        |
| [aes](https://studygolang.com/static/pkgdoc/pkg/crypto_aes.htm) |      | aes包实现了AES加密算法，参见U.S. Federal Information Processing Standards Publication 197. |
| [cipher](https://studygolang.com/static/pkgdoc/pkg/crypto_cipher.htm) |      | cipher包实现了多个标准的用于包装底层块加密算法的加密算法实现. |
| [des](https://studygolang.com/static/pkgdoc/pkg/crypto_des.htm) |      | des包实现了DES标准和TDEA算法，参见U.S. Federal Information Processing Standards Publication 46-3. |
| [dsa](https://studygolang.com/static/pkgdoc/pkg/crypto_dsa.htm) |      | dsa包实现FIPS 186-3定义的数字签名算法（Digital Signature Algorithm），即DSA算法. |
| [ecdsa](https://studygolang.com/static/pkgdoc/pkg/crypto_ecdsa.htm) |      | ecdsa包实现了椭圆曲线数字签名算法，参见FIPS 186-3.           |
| [elliptic](https://studygolang.com/static/pkgdoc/pkg/crypto_elliptic.htm) |      | elliptic包实现了几条覆盖素数有限域的标准椭圆曲线.            |
| [hmac](https://studygolang.com/static/pkgdoc/pkg/crypto_hmac.htm) |      | hmac包实现了U.S. Federal Information Processing Standards Publication 198规定的HMAC（加密哈希信息认证码）. |
| [md5](https://studygolang.com/static/pkgdoc/pkg/crypto_md5.htm) |      | md5包实现了MD5哈希算法，参见RFC 1321.                        |
| [rand](https://studygolang.com/static/pkgdoc/pkg/crypto_rand.htm) |      | rand包实现了用于加解密的更安全的随机数生成器.                |
| [rc4](https://studygolang.com/static/pkgdoc/pkg/crypto_rc4.htm) |      | rc4包实现了RC4加密算法，参见Bruce Schneier's Applied Cryptography. |
| [rsa](https://studygolang.com/static/pkgdoc/pkg/crypto_rsa.htm) |      | rsa包实现了PKCS#1规定的RSA加密算法.                          |
| [sha1](https://studygolang.com/static/pkgdoc/pkg/crypto_sha1.htm) |      | sha1包实现了SHA1哈希算法，参见RFC 3174.                      |
| [sha256](https://studygolang.com/static/pkgdoc/pkg/crypto_sha256.htm) |      | sha256包实现了SHA224和SHA256哈希算法，参见FIPS 180-4.        |
| [sha512](https://studygolang.com/static/pkgdoc/pkg/crypto_sha512.htm) |      | sha512包实现了SHA384和SHA512哈希算法，参见FIPS 180-2.        |
| [subtle](https://studygolang.com/static/pkgdoc/pkg/crypto_subtle.htm) |      | Package subtle implements functions that are often useful in cryptographic code but require careful thought to use correctly. |
| [tls](https://studygolang.com/static/pkgdoc/pkg/crypto_tls.htm) |      | tls包实现了TLS 1.2，细节参见RFC 5246.                        |
| [x509](https://studygolang.com/static/pkgdoc/pkg/crypto_x509.htm) |      | x509包解析X.509编码的证书和密钥.                             |
| [pkix](https://studygolang.com/static/pkgdoc/pkg/crypto_x509_pkix.htm) |      | pkix包提供了共享的、低层次的结构体，用于ASN.1解析和X.509证书、CRL、OCSP的序列化. |
| [database](https://studygolang.com/static/pkgdoc/main.html#database) |      |                                                              |
| [sql](https://studygolang.com/static/pkgdoc/pkg/database_sql.htm) |      | sql 包提供了通用的SQL（或类SQL）数据库接口.                  |
| [driver](https://studygolang.com/static/pkgdoc/pkg/database_sql_driver.htm) |      | driver包定义了应被数据库驱动实现的接口，这些接口会被sql包使用. |
| [debug](https://studygolang.com/static/pkgdoc/main.html#debug) |      |                                                              |
| [dwarf](https://studygolang.com/static/pkgdoc/pkg/debug_dwarf.htm) |      | Package dwarf provides access to DWARF debugging information loaded from executable files, as defined in the DWARF 2.0 Standard at http://dwarfstd.org/doc/dwarf-2.0.0.pdf |
| [elf](https://studygolang.com/static/pkgdoc/pkg/debug_elf.htm) |      | Package elf implements access to ELF object files.           |
| [gosym](https://studygolang.com/static/pkgdoc/pkg/debug_gosym.htm) |      | Package gosym implements access to the Go symbol and line number tables embedded in Go binaries generated by the gc compilers. |
| [macho](https://studygolang.com/static/pkgdoc/pkg/debug_macho.htm) |      | Package macho implements access to Mach-O object files.      |
| [pe](https://studygolang.com/static/pkgdoc/pkg/debug_pe.htm) |      | Package pe implements access to PE (Microsoft Windows Portable Executable) files. |
| [plan9obj](https://studygolang.com/static/pkgdoc/pkg/debug_plan9obj.htm) |      | Package plan9obj implements access to Plan 9 a.out object files. |
| [encoding](https://studygolang.com/static/pkgdoc/pkg/encoding.htm) |      | encoding包定义了供其它包使用的可以将数据在字节水平和文本表示之间转换的接口. |
| [ascii85](https://studygolang.com/static/pkgdoc/pkg/encoding_ascii85.htm) |      | ascii85 包是对 ascii85 的数据编码的实现.                     |
| [asn1](https://studygolang.com/static/pkgdoc/pkg/encoding_asn1.htm) |      | asn1包实现了DER编码的ASN.1数据结构的解析，参见ITU-T Rec X.690. |
| [base32](https://studygolang.com/static/pkgdoc/pkg/encoding_base32.htm) |      | base32包实现了RFC 4648规定的base32编码.                      |
| [base64](https://studygolang.com/static/pkgdoc/pkg/encoding_base64.htm) |      | base64实现了RFC 4648规定的base64编码.                        |
| [binary](https://studygolang.com/static/pkgdoc/pkg/encoding_binary.htm) |      | binary包实现了简单的数字与字节序列的转换以及变长值的编解码.  |
| [csv](https://studygolang.com/static/pkgdoc/pkg/encoding_csv.htm) |      | csv读写逗号分隔值（csv）的文件.                              |
| [gob](https://studygolang.com/static/pkgdoc/pkg/encoding_gob.htm) |      | gob包管理gob流——在编码器（发送器）和解码器（接受器）之间交换的binary值. |
| [hex](https://studygolang.com/static/pkgdoc/pkg/encoding_hex.htm) |      | hex包实现了16进制字符表示的编解码.                           |
| [json](https://studygolang.com/static/pkgdoc/pkg/encoding_json.htm) |      | json包实现了json对象的编解码，参见RFC 4627.                  |
| [pem](https://studygolang.com/static/pkgdoc/pkg/encoding_pem.htm) |      | pem包实现了PEM数据编码（源自保密增强邮件协议）.              |
| [xml](https://studygolang.com/static/pkgdoc/pkg/encoding_xml.htm) |      | Package xml implements a simple XML 1.0 parser that understands XML name spaces. |
| [errors](https://studygolang.com/static/pkgdoc/pkg/errors.htm) |      | error 包实现了用于错误处理的函数.                            |
| [expvar](https://studygolang.com/static/pkgdoc/pkg/expvar.htm) |      | expvar包提供了公共变量的标准接口，如服务的操作计数器.        |
| [flag](https://studygolang.com/static/pkgdoc/pkg/flag.htm)   |      | flag 包实现命令行标签解析.                                   |
| [fmt](https://studygolang.com/static/pkgdoc/pkg/fmt.htm)     |      | fmt 包实现了格式化I/O函数，类似于C的 printf 和 scanf.        |
| [go](https://studygolang.com/static/pkgdoc/main.html#go)     |      |                                                              |
| [ast](https://studygolang.com/static/pkgdoc/pkg/go_ast.htm)  |      | Package ast declares the types used to represent syntax trees for Go packages. |
| [build](https://studygolang.com/static/pkgdoc/pkg/go_build.htm) |      | Package build gathers information about Go packages.         |
| [constant](https://studygolang.com/static/pkgdoc/pkg/go_constant.htm) |      | Package constant implements Values representing untyped Go constants and their corresponding operations. |
| [doc](https://studygolang.com/static/pkgdoc/pkg/go_doc.htm)  |      | Package doc extracts source code documentation from a Go AST. |
| [format](https://studygolang.com/static/pkgdoc/pkg/go_format.htm) |      | Package format implements standard formatting of Go source.  |
| [importer](https://studygolang.com/static/pkgdoc/pkg/go_importer.htm) |      | Package importer provides access to export data importers.   |
| [parser](https://studygolang.com/static/pkgdoc/pkg/go_parser.htm) |      | Package parser implements a parser for Go source files.      |
| [printer](https://studygolang.com/static/pkgdoc/pkg/go_printer.htm) |      | Package printer implements printing of AST nodes.            |
| [scanner](https://studygolang.com/static/pkgdoc/pkg/go_scanner.htm) |      | Package scanner implements a scanner for Go source text.     |
| [token](https://studygolang.com/static/pkgdoc/pkg/go_token.htm) |      | Package token defines constants representing the lexical tokens of the Go programming language and basic operations on tokens (printing, predicates). |
| [types](https://studygolang.com/static/pkgdoc/pkg/go_types.htm) |      | Package types declares the data types and implements the algorithms for type-checking of Go packages. |
| [hash](https://studygolang.com/static/pkgdoc/main.html#hash) |      | hash包提供hash函数的接口.                                    |
| [adler32](https://studygolang.com/static/pkgdoc/pkg/hash_adler32.htm) |      | adler32包实现了Adler-32校验和算法，参见RFC 1950.             |
| [crc32](https://studygolang.com/static/pkgdoc/pkg/hash_crc32.htm) |      | crc32包实现了32位循环冗余校验（CRC-32）的校验和算法.         |
| [crc64](https://studygolang.com/static/pkgdoc/pkg/hash_crc64.htm) |      | Package crc64 implements the 64-bit cyclic redundancy check, or CRC-64, checksum. |
| [fnv](https://studygolang.com/static/pkgdoc/pkg/hash_fnv.htm) |      | fnv包实现了FNV-1和FNV-1a（非加密hash函数）.                  |
| [html](https://studygolang.com/static/pkgdoc/pkg/html.htm)   |      | html包提供了用于转义和解转义HTML文本的函数.                  |
| [template](https://studygolang.com/static/pkgdoc/pkg/html_template.htm) |      | template包（html/template）实现了数据驱动的模板，用于生成可对抗代码注入的安全HTML输出. |
| [image](https://studygolang.com/static/pkgdoc/pkg/image.htm) |      | image实现了基本的2D图片库.                                   |
| [color](https://studygolang.com/static/pkgdoc/pkg/image_color.htm) |      | color 包实现了基本的颜色库。                                 |
| [palette](https://studygolang.com/static/pkgdoc/pkg/image_color_palette.htm) |      | palette包提供了标准的调色板.                                 |
| [draw](https://studygolang.com/static/pkgdoc/pkg/image_draw.htm) |      | draw 包提供组装图片的方法.                                   |
| [gif](https://studygolang.com/static/pkgdoc/pkg/image_gif.htm) |      | gif 包实现了GIF图片的解码.                                   |
| [jpeg](https://studygolang.com/static/pkgdoc/pkg/image_jpeg.htm) |      | jpeg包实现了jpeg格式图像的编解码.                            |
| [png](https://studygolang.com/static/pkgdoc/pkg/image_png.htm) |      | png 包实现了PNG图像的编码和解码.                             |
| [index](https://studygolang.com/static/pkgdoc/main.html#index) |      |                                                              |
| [suffixarray](https://studygolang.com/static/pkgdoc/pkg/index_suffixarray.htm) |      | suffixarrayb包通过使用内存中的后缀树实现了对数级时间消耗的子字符串搜索. |
| [io](https://studygolang.com/static/pkgdoc/pkg/io.htm)       |      | io 包为I/O原语提供了基础的接口.                              |
| [ioutil](https://studygolang.com/static/pkgdoc/pkg/io_ioutil.htm) |      | ioutil 实现了一些I/O的工具函数。                             |
| [log](https://studygolang.com/static/pkgdoc/pkg/log.htm)     |      | log包实现了简单的日志服务.                                   |
| [syslog](https://studygolang.com/static/pkgdoc/pkg/log_syslog.htm) |      | syslog包提供一个简单的系统日志服务的接口.                    |
| [math](https://studygolang.com/static/pkgdoc/pkg/math.htm)   |      | math 包提供了基本常数和数学函数。                            |
| [big](https://studygolang.com/static/pkgdoc/pkg/math_big.htm) |      | big 包实现了（大数的）高精度运算.                            |
| [cmplx](https://studygolang.com/static/pkgdoc/pkg/math_cmplx.htm) |      | cmplx 包为复数提供了基本的常量和数学函数.                    |
| [rand](https://studygolang.com/static/pkgdoc/pkg/math_rand.htm) |      | rand 包实现了伪随机数生成器.                                 |
| [mime](https://studygolang.com/static/pkgdoc/pkg/mime.htm)   |      | mime实现了MIME的部分规定.                                    |
| [multipart](https://studygolang.com/static/pkgdoc/pkg/mime_multipart.htm) |      | multipart实现了MIME的multipart解析，参见RFC 2046.            |
| [quotedprintable](https://studygolang.com/static/pkgdoc/pkg/mime_quotedprintable.htm) |      | Package quotedprintable implements quoted-printable encoding as specified by RFC 2045. |
| [net](https://studygolang.com/static/pkgdoc/pkg/net.htm)     |      | net包提供了可移植的网络I/O接口，包括TCP/IP、UDP、域名解析和Unix域socket. |
| [http](https://studygolang.com/static/pkgdoc/pkg/net_http.htm) |      | http包提供了HTTP客户端和服务端的实现.                        |
| [cgi](https://studygolang.com/static/pkgdoc/pkg/net_http_cgi.htm) |      | cgi 包实现了RFC3875协议描述的CGI（公共网关接口）.            |
| [cookiejar](https://studygolang.com/static/pkgdoc/pkg/net_http_cookiejar.htm) |      | cookiejar包实现了保管在内存中的符合RFC 6265标准的http.CookieJar接口. |
| [fcgi](https://studygolang.com/static/pkgdoc/pkg/net_http_fcgi.htm) |      | fcgi 包实现了FastCGI协议.                                    |
| [httptest](https://studygolang.com/static/pkgdoc/pkg/net_http_httptest.htm) |      | httptest 包提供HTTP测试的单元工具.                           |
| [httptrace](https://studygolang.com/static/pkgdoc/pkg/net_http_httptrace.htm) |      | Package httptrace provides mechanisms to trace the events within HTTP client requests. |
| [httputil](https://studygolang.com/static/pkgdoc/pkg/net_http_httputil.htm) |      | httputil包提供了HTTP公用函数，是对net/http包的更常见函数的补充. |
| [pprof](https://studygolang.com/static/pkgdoc/pkg/net_http_pprof.htm) |      | pprof 包通过提供HTTP服务返回runtime的统计数据，这个数据是以pprof可视化工具规定的返回格式返回的. |
| [mail](https://studygolang.com/static/pkgdoc/pkg/net_mail.htm) |      | mail 包实现了解析邮件消息的功能.                             |
| [rpc](https://studygolang.com/static/pkgdoc/pkg/net_rpc.htm) |      | rpc 包提供了一个方法来通过网络或者其他的I/O连接进入对象的外部方法. |
| [jsonrpc](https://studygolang.com/static/pkgdoc/pkg/net_rpc_jsonrpc.htm) |      | jsonrpc 包使用了rpc的包实现了一个JSON-RPC的客户端解码器和服务端的解码器. |
| [smtp](https://studygolang.com/static/pkgdoc/pkg/net_smtp.htm) |      | smtp包实现了简单邮件传输协议（SMTP），参见RFC 5321.          |
| [textproto](https://studygolang.com/static/pkgdoc/pkg/net_textproto.htm) |      | textproto实现了对基于文本的请求/回复协议的一般性支持，包括HTTP、NNTP和SMTP. |
| [url](https://studygolang.com/static/pkgdoc/pkg/net_url.htm) |      | url包解析URL并实现了查询的逸码，参见RFC 3986.                |
| [os](https://studygolang.com/static/pkgdoc/pkg/os.htm)       |      | os包提供了操作系统函数的不依赖平台的接口.                    |
| [exec](https://studygolang.com/static/pkgdoc/pkg/os_exec.htm) |      | exec包执行外部命令.                                          |
| [signal](https://studygolang.com/static/pkgdoc/pkg/os_signal.htm) |      | signal包实现了对输入信号的访问.                              |
| [user](https://studygolang.com/static/pkgdoc/pkg/os_user.htm) |      | user包允许通过名称或ID查询用户帐户.                          |
| [path](https://studygolang.com/static/pkgdoc/pkg/path.htm)   |      | path实现了对斜杠分隔的路径的实用操作函数.                    |
| [filepath](https://studygolang.com/static/pkgdoc/pkg/path_filepath.htm) |      | filepath包实现了兼容各操作系统的文件路径的实用操作函数.      |
| [plugin](https://studygolang.com/static/pkgdoc/pkg/plugin.htm) |      | Package plugin implements loading and symbol resolution of Go plugins. |
| [reflect](https://studygolang.com/static/pkgdoc/pkg/reflect.htm) |      | reflect包实现了运行时反射，允许程序操作任意类型的对象.       |
| [regexp](https://studygolang.com/static/pkgdoc/pkg/regexp.htm) |      | regexp包实现了正则表达式搜索.                                |
| [syntax](https://studygolang.com/static/pkgdoc/pkg/regexp_syntax.htm) |      | Package syntax parses regular expressions into parse trees and compiles parse trees into programs. |
| [runtime](https://studygolang.com/static/pkgdoc/pkg/runtime.htm) |      | TODO(osc): 需更新 runtime 包含与Go的运行时系统进行交互的操作，例如用于控制Go程的函数. |
| [cgo](https://studygolang.com/static/pkgdoc/pkg/runtime_cgo.htm) |      | cgo 包含有 cgo 工具生成的代码的运行时支持.                   |
| [debug](https://studygolang.com/static/pkgdoc/pkg/runtime_debug.htm) |      | debug 包含有程序在运行时调试其自身的功能.                    |
| [pprof](https://studygolang.com/static/pkgdoc/pkg/runtime_pprof.htm) |      | pprof 包按照可视化工具 pprof 所要求的格式写出运行时分析数据. |
| [race](https://studygolang.com/static/pkgdoc/pkg/runtime_race.htm) |      | race 包实现了数据竞争检测逻辑.                               |
| [trace](https://studygolang.com/static/pkgdoc/pkg/runtime_trace.htm) |      | Go execution tracer.                                         |
| [sort](https://studygolang.com/static/pkgdoc/pkg/sort.htm)   |      | sort 包为切片及用户定义的集合的排序操作提供了原语.           |
| [strconv](https://studygolang.com/static/pkgdoc/pkg/strconv.htm) |      | strconv包实现了基本数据类型和其字符串表示的相互转换.         |
| [strings](https://studygolang.com/static/pkgdoc/pkg/strings.htm) |      | strings包实现了用于操作字符的简单函数.                       |
| [sync](https://studygolang.com/static/pkgdoc/pkg/sync.htm)   |      | sync 包提供了互斥锁这类的基本的同步原语.                     |
| [atomic](https://studygolang.com/static/pkgdoc/pkg/sync_atomic.htm) |      | atomic 包提供了底层的原子性内存原语，这对于同步算法的实现很有用. |
| [syscall](https://studygolang.com/static/pkgdoc/pkg/syscall.htm) |      | Package syscall contains an interface to the low-level operating system primitives. |
| [testing](https://studygolang.com/static/pkgdoc/pkg/testing.htm) |      | Package testing provides support for automated testing of Go packages. |
| [iotest](https://studygolang.com/static/pkgdoc/pkg/testing_iotest.htm) |      | Package iotest implements Readers and Writers useful mainly for testing. |
| [quick](https://studygolang.com/static/pkgdoc/pkg/testing_quick.htm) |      | Package quick implements utility functions to help with black box testing. |
| [text](https://studygolang.com/static/pkgdoc/main.html#text) |      |                                                              |
| [scanner](https://studygolang.com/static/pkgdoc/pkg/text_scanner.htm) |      | scanner包提供对utf-8文本的token扫描服务.                     |
| [tabwriter](https://studygolang.com/static/pkgdoc/pkg/text_tabwriter.htm) |      | tabwriter包实现了写入过滤器（tabwriter.Writer），可以将输入的缩进修正为正确的对齐文本. |
| [template](https://studygolang.com/static/pkgdoc/pkg/text_template.htm) |      | template包实现了数据驱动的用于生成文本输出的模板.            |
| [parse](https://studygolang.com/static/pkgdoc/pkg/text_template_parse.htm) |      | Package parse builds parse trees for templates as defined by text/template and html/template. |
| [time](https://studygolang.com/static/pkgdoc/pkg/time.htm)   |      | time包提供了时间的显示和测量用的函数.                        |
| [unicode](https://studygolang.com/static/pkgdoc/pkg/unicode.htm) |      | unicode 包提供了一些测试Unicode码点属性的数据和函数.         |
| [utf16](https://studygolang.com/static/pkgdoc/pkg/unicode_utf16.htm) |      | utf16 包实现了对UTF-16序列的编码和解码。                     |
| [utf8](https://studygolang.com/static/pkgdoc/pkg/unicode_utf8.htm) |      | utf8 包实现了支持UTF-8文本编码的函数和常量.                  |
| [unsafe](https://studygolang.com/static/pkgdoc/pkg/unsafe.htm) |      | unsafe 包含有关于Go程序类型安全的所有操作.                   |

## 其它包

### 子代码库

这些包是 Go 项目的一部分，但并未在主源码树中。它们在比 Go 核心库更加宽松的[兼容性需求](https://studygolang.com/doc/go1compat)下开发。 可通过“[go get](https://studygolang.com/cmd/go/#hdr-Download_and_install_packages_and_dependencies)”安装它们，子代码库的[文档](http://godoc.org/-/subrepo)和[源码](https://github.com/golang)可通过相应的链接访问

- [crypto](https://github.com/golang/crypto) — 附加的加密包。
- [image](https://github.com/golang/image) — 附加的图像包。
- [net](https://github.com/golang/net) — 附加的网络包。
- [sys](https://github.com/golang/sys) — 系统调用包。
- [text](https://github.com/golang/text) — 文本处理包。
- [tools](https://github.com/golang/tools) — godoc、vet、cover 及其它工具。
- [exp](https://github.com/golang/exp) — 实验性代码（可能不经警告就更改，请小心对待）。

## 参考资料
https://studygolang.com/
http://docscn.studygolang.com/
https://books.studygolang.com/