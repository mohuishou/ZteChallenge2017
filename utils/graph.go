package utils

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

//INF 无穷大
const INF = 0xfffff

//Graph 图
type Graph struct {
	VNum, ENum int     //顶点、边的个数
	MaxVexNum  int     //最多可以经过的顶点数目
	MustVex    []int   //必过点集合
	MustEdge   [][]int //必过边集合
	G          [][]int //邻接矩阵
}

//CreateGraph 创建一个图并且初始化邻接矩阵
func CreateGraph(n int) (graph Graph) {
	graph = Graph{VNum: n}
	graph.G = make([][]int, n)
	for i := 0; i < n; i++ {
		graph.G[i] = make([]int, n)
		for j := 0; j < n; j++ {
			graph.G[i][j] = INF
			if i == j {
				graph.G[i][j] = 0
			}
		}
	}
	return graph
}

//BuildGraph 创建图
func BuildGraph(path string) (graph Graph) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)

	i, j := 0, 0
	//边的数目
A:
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return graph
			}
			panic(err)
		}
		line = strings.TrimSpace(line)
		dataStr := strings.Split(line, " ")

		if dataStr == nil || (len(dataStr) == 1 && dataStr[0] == "") {
			i++
			j = 0
			continue
		}

		//转成int之后的数据
		data := stringData2int(dataStr)

		switch i {
		case 0:
			if len(data) != 6 {
				panic("数据格式错误，第一行需要6个数据，现有：" + strconv.Itoa(len(data)) + "个")
			}
			graph = CreateGraph(data[4])
			graph.MaxVexNum = data[0]
			graph.MustVex = make([]int, data[1])
			graph.MustEdge = make([][]int, data[2])
			graph.ENum = data[5]
		case 1:
			if len(data) != 3 {
				panic("数据格式错误，食蚁兽边需要3个数据，现有：" + strconv.Itoa(len(data)) + "个")
			}
			graph.G[data[0]][data[1]] = data[2]
			graph.G[data[1]][data[0]] = data[2]
		case 2:
			graph.MustVex[j] = data[0]
			j++
		case 3:
			if len(data) != 2 {
				panic("数据格式错误，食蚁兽边需要2个数据，现有：" + strconv.Itoa(len(data)) + "个")
			}
			graph.MustEdge[j] = data
			j++
		case 4:
			if len(data) != 2 {
				panic("数据格式错误，食蚁兽边需要2个数据，现有：" + strconv.Itoa(len(data)) + "个")
			}
			graph.G[data[0]][data[1]] = INF
		default:
			break A
		}

	}
	return graph
}

//stringData2int 字符串切片转int切片
func stringData2int(s []string) []int {
	data := make([]int, len(s))
	for i := range s {
		val, err := strconv.Atoi(s[i])
		if err != nil {
			panic("错误：数据转换错误" + err.Error())
		}
		data[i] = val
	}
	return data
}
