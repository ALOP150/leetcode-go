package sumrange

type NumArray struct {
	vals []int
	size int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		vals: nums,
		size: len(nums),
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i >= this.size || j >= this.size || i > j {
		return 0
	}

	res := 0
	for i < j {
		res += this.vals[i]
	}
	return res
}
