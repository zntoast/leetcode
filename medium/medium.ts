
// LC 3195. 包含所有 1 的最小矩形面积 I
function MinimumArea(grid: number[][]): number {
  let ans: number = 0;
  let left_x: number, right_x: number, left_y: number, right_y: number;

  // 初始化边界变量
  left_x = Number.MAX_VALUE;
  right_x = 0;
  left_y = Number.MAX_VALUE;
  right_y = 0;

  for (let i: number = 0; i < grid.length; i++) {
    for (let j: number = 0; j < grid[i].length; j++) {
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
};

// LC 3255. 长度为 K 的子数组的能量值 II
function ResultsArray(nums: number[], k: number): number[] {
  let n :number = nums.length,cnt :number= 0;
  const ans: number[] = Array(n-k+1).fill(-1);
  for (let i: number = 0; i < n; i++) {
    cnt = (i === 0 || nums[i] - nums[i-1] !==1)?1:cnt+1;
    if (cnt >= k ){
      ans[i-k+1] = nums[i];
    }
  }
  return ans;
};

ResultsArray([1,2,3,4,3,2,5], 3);