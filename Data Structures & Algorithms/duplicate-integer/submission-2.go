func hasDuplicate(nums []int) bool {
    memory := make(map[int]struct{}, len(nums))

    for _, num := range nums {
        if _, exists := memory[num]; exists {
            return true
        }
        memory[num] = struct{}{}
    }
    return false
}
