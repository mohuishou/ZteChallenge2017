package utils

//MInt 对int型切片进行扩充
type MInt []int

//Swap 交换
func (mint MInt) Swap(i, j int) {
	tmp := mint[i]
	mint[i] = mint[j]
	mint[j] = tmp
}

var seq = make([][]int, 0)

//Permutation 全排列
func Permutation(vexs MInt, k, m int) [][]int {
	permutation(vexs, k, m)
	return seq
}

//permutation 对序列进行全排列
func permutation(vexs MInt, k, m int) {
	if k == m {
		// seq[0] = arr
		tmp := make([]int, len(vexs))
		copy(tmp, vexs)
		seq = append(seq, tmp)
	} else {
		for i := k; i < m; i++ {
			vexs.Swap(i, k)
			permutation(vexs, k+1, m)
			vexs.Swap(i, k)
		}
	}
}
