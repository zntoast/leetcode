// LC 3195. 包含所有 1 的最小矩形面积 I
function MinimumArea(grid) {
    var ans = 0;
    var left_x, right_x, left_y, right_y;
    // 初始化边界变量
    left_x = Number.MAX_VALUE;
    right_x = 0;
    left_y = Number.MAX_VALUE;
    right_y = 0;
    for (var i = 0; i < grid.length; i++) {
        for (var j = 0; j < grid[i].length; j++) {
            if (grid[i][j] != 1) {
                continue;
            }
            // 更新左边界
            if (j < left_x) {
                left_x = j;
            }
            // 更新右边界
            if (j > right_x) {
                right_x = j;
            }
            // 更新上边界
            if (i < left_y) {
                left_y = i;
            }
            // 更新下边界
            if (i > right_y) {
                right_y = i;
            }
        }
    }
    // 计算矩形的面积
    if (left_x !== Number.MAX_VALUE) {
        ans = (right_x - left_x + 1) * (right_y - left_y + 1);
    }
    return ans;
}
;
// LC 910. 最小差值 II
function SmallestRangeII(nums, k) {
    var total = 0;
    for (var i = 0; i < nums.length; i++) {
        total += nums[i];
    }
    var meat = total / nums.length;
    for (var i = 0; i < nums.length; i++) {
        if (nums[i] >= total / nums.length) {
            nums[i] -= k;
        }
        else {
            nums[i] += k;
        }
    }
    console.log(nums);
    console.log(meat);
    var max = Math.max.apply(Math, nums);
    var min = Math.min.apply(Math, nums);
    return max - min;
}
;
console.log(SmallestRangeII([7, 8, 8], 5));
