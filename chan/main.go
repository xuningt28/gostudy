package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	//
	Id int
	//
	RandNum int
}

type Result struct {
	//
	job *Job
	// sum
	sum int
}

func main() {
	//  需要两个管道
	// 1 job 管道
	jobChan := make(chan *Job, 128)

	// 2 结果管道
	resultChan := make(chan *Result, 128)

	// 3 工作池
	createPool(1,jobChan,resultChan)

	// 4 打印协程
	go func(resultChan chan *Result){
		// 遍历结果管道打印
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n",result.job.Id,result.job.RandNum,result.sum)
		}
	}(resultChan)

	// 循环创建job,输入到管道
	var id int
	for{
		id++
		//
		r_num := rand.Int()
		job:=&Job{
			Id:id,
			RandNum:r_num,
		}

        jobChan <- job


	}

	
}

// 创建工作池
// 参数1 ，开几个工作池
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	//
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
           // 执行运算
           // 遍历job数据，进行相加
           for job:= range jobChan{
           	// 接随机数
           	r_num:=job.RandNum
           	// 随机数每一位相加
           	// 定义返回值
           	var sum int
           	for r_num!=0 {
           		tmp := r_num %10
           		sum += tmp
           		r_num /= 10
			}

           	// 想要的结果是Result
           	r := &Result{
           		job:job,
           		sum:sum,
			}
           	// 运算结果扔到管道
           	resultChan<-r
		   }
		}(jobChan, resultChan)
	}
}
