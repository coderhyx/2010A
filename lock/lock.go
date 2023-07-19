package main

func lock() {
	//conf := RedisConf{
	//	Host: "127.0.0.1:55000",
	//	Type: "node",
	//	Pass: "123456",
	//	Tls:  false,
	//}
	//
	//rds := MustNewRedis(conf)
	//
	//lock := NewRedisLock(rds, "test")
	//
	//// 设置过期时间
	//lock.SetExpire(10)
	//
	//// 尝试获取锁
	//acquire, err := lock.Acquire()
	//
	//switch {
	//case err != nil:
	//	// deal err
	//case acquire:
	//	// 获取到锁
	//	defer lock.Release() // 释放锁
	//	// 业务逻辑
	//
	//case !acquire:
	//	// 没有拿到锁 wait?
	//}
}
