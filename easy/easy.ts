/**
 * LC 3200. 三角形的最大高度
 * @param {number} red
 * @param {number} blue
 * @return {number}
 */
var MaxHeightOfTriangle = function (red, blue) {
  const masHeight = (x: number, y: number) => {
    for (let i = 1; ; i++) {
      if (i % 2 == 1) {
        x -= i
        if (x < 0) {
          return i - 1
        }
      } else {
        y -= i
        if (y < 0) {
          return i - 1
        }
      }
    }
  }
  return Math.max(masHeight(red, blue), masHeight(blue, red))
};

// console.log(MaxHeightOfTriangle(2, 4))
// console.log(MaxHeightOfTriangle(2, 1))
// console.log(MaxHeightOfTriangle(1, 1))
// console.log(MaxHeightOfTriangle(10, 1))



// LC 3194. 最小元素和最大元素的最小平均值
function MinimumAverage(nums: number[]): number {
  nums.sort((a, b) => a - b)
  const len = nums.length
  let ans = nums[len - 1]
  for (let i = 0; i < len / 2; i++) {
    let temp = (nums[i] + nums[len - i - 1]) / 2
    ans = Math.min(ans, temp)
  }
  return ans
};
