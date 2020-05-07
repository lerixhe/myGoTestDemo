
代码段如下：

	// 2020/04/29 11:46:58.675 [I] [serverCheck.go:29]  total speedtest servers count: 26693
	logs.Info("total speedtest servers count:", len(speedtestIPs))
	// use channel to limit the max request count as the same time
	limitChan := make(chan int, maxRequestLimit)
	// use waitgroup to wait all goroutines finished
	wg := &sync.WaitGroup{}
	wg.Add(len(speedtestIPs))
	for _, speedtestIP := range speedtestIPs {
		limitChan <- 1
		go func(speedtestIP models.SpeedtestInfo) {
			speedtestStatus := models.SpeedtestStatus{
				ID:        speedtestIP.ID,
				Status:    1,
				UpdatedAt: now,
			}
			if !CheckHost(speedtestIP, time.Duration(timeout)) {
				speedtestStatus.Status = 0
			}
			<-limitChan
			wg.Done()
			speedtestIPsStatus = append(speedtestIPsStatus, speedtestStatus)
		}(speedtestIP)
	}
	wg.Wait()
	//插入
	// 2020/04/29 11:48:09.091 [I] [serverCheck.go:55]  finish check total 26686 speedtest servers
	logs.Info("finish check total %d speedtest servers", len(speedtestIPsStatus))



	总结：切片的append函数并非并发安全的。
	本代码段，起n个goroutine并发的向slice中append数据，n个goroutine都结束后，打印slice的长度。
	发现len(slice)<n.
	原因：slice是对数组一个连续片段的引用，当slice长度增加的时候，可能底层的数组会被换掉。
	当出在换底层数组之前，切片同时被多个goroutine拿到，并执行append操作。那么很多goroutine的append结果会被覆盖，导致n个gouroutine append后，长度小于n。


	解决方案：
	1.加锁，性能有影响
	2.使用生产者消费者的多对一模型，所有生产者不自己append，而是写入管道。由一个专门的消费者线程执行append