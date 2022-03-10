# rabbitmq-pool
golang rabbitmq 连接池

# 实现的功能
1、复用connection、channel  
2、channel最大连接数限制  
3、默认创建的channel不足时动态创建channel至 channel最大连接数  
4、动态创建的channel用完自动销毁恢复默认channel数  
