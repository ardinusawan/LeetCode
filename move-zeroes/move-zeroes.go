func moveZeroes(nums []int)  {
    for i:=len(nums)-1;i>=0;i-- {
        if nums[i] == 0 {
            nums = append(nums[:i], nums[i+1:]...)
            nums = append(nums, 0)
        }
    }
}