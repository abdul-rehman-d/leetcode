package trapping_rain_water

func Trap(height []int) int {
    total := 0

    leftIdx, rightIdx := 0, len(height)-1
    left, right := height[leftIdx], height[rightIdx]

    for leftIdx < rightIdx {
        if right > left {
            leftIdx += 1
            left = max(height[leftIdx], left)
            total += (left - height[leftIdx])
        } else {
            rightIdx -= 1
            right = max(height[rightIdx], right)
            total += (right - height[rightIdx])
        }
    }
    return total
}

