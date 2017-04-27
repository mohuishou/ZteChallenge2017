package main

import (
	"flag"

	"fmt"

	"time"

	"github.com/mohuishou/ZteChallenge2017/deploy"
	"github.com/mohuishou/ZteChallenge2017/utils"
)

func main() {

	var path string
	//flag.StringVar这样的函数第一个参数换成了变量地址，后面的参数和flag.String是一样的。
	flag.StringVar(&path, "path", "./example/case_1.txt", "Where is your graph text?")

	flag.Parse() //解析输入的参数

	g := utils.BuildGraph(path)

	//记录开始时间
	start := time.Now()

	p, d := deploy.Deploy(g)

	//记录结束时间
	end := time.Now()

	if len(p) > g.MaxVexNum {
		fmt.Println("错误：当前不存在", g.MaxVexNum, "个节点以内的路径")
	}

	fmt.Print("路径为：")
	for i := range p {
		fmt.Print(p[i])
		if i < len(p)-1 {
			fmt.Print("->")
		}
	}
	fmt.Println()
	fmt.Println("节点数目：", len(p))
	fmt.Println("路径长度：", d)

	//输出执行时间，单位为毫秒。
	fmt.Println("运行时间：", end.Sub(start))
}
