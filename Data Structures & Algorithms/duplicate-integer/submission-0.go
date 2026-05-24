func hasDuplicate(nums []int) bool {
    memory := make(map[int]bool)

    for i := 0; i < len(nums); i++ {
        _, ok := memory[nums[i]]
        if ok {
            return true
        }
        memory[nums[i]] = true
    }

    return false
}
