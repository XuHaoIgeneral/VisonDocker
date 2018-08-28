package test

import (
	"testing"
	"visonDocker/components/docker"
)

func TestTopDocker(t *testing.T) {
	test :=docker.TopDocker("as")
	if test!="" {
		t.Log(test)
	}
}

func BenchmarkTopDocker(b *testing.B)  {
	b.StopTimer() //调用该函数停止压力测试的时间计数
	//
	////做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	////这样这些时间不影响我们测试函数本身的性能
	b.StartTimer() //重新开始时间
	for i := 0; i < 1000; i++ {
		docker.TopDocker("aa")
	}
}

