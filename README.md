#####项目结构介绍
1. bin目录存放可执行程序  
2. cmd目录可执行程序的入口目录  
3. core目录判断提交字符串数组的核心校验方法  
4. memory目录数据的内存存储  
5. tls证书生成目录,目录中makessl.sh可以生成证书  
#####安装
* 执行make all会自动把编译后在可以执行文件放到bin目录下，会生成server和client两个可执行文件。  
由于没有配置文件，要改变程序的执行路径，需要同时复制bin和tls两个目录可以执行server程序就可以正常启动。  
如果win系统没有安装make直接进行cmd目录进行手动编译。



