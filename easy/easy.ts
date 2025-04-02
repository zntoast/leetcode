/**
 * LC 3200. 三角形的最大高度
 * @param {number} red
 * @param {number} blue
 * @return {number}
 */
var MaxHeightOfTriangle = function (red: number, blue: number) {
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


// LC 999. 可以被一步捕获的棋子数
function numRookCaptures(board: string[][]): number {
  let rus: number = 0
  let R: Array<number> = new Array(2)
  for (let i = 0; i < 8; i++) {
    for (let j = 0; j < 8; j++) {
      if (board[i][j] === 'R') {
        R = [i, j]
      }
    }
  }

  for (let k = 1; k <= 8; k++) {
    if (R[0] - k >= 0 && board[R[0] - k][R[1]] === 'B') {
      rus++
    }
    if (R[0] + k < 8 && board[R[0] + k][R[1]] === 'B') {
      rus++
    }
    if (R[1] - k >= 0 && board[R[0]][R[1] - k] === 'B') {
      rus++
    }
    if (R[1] + k < 8 && board[R[0]][R[1] + k] === 'B') {
      rus++
    }
  }

  return rus
};

// LC 2873. 有序三元组中的最大值 I
function maximumTripletValue(nums: number[]): number {
  if (nums.length < 3) {
    return 0
  }
  var rus: number = 0
  for (var i = 0; i < nums.length - 2; i++) {
    for (var j = i + 1; j < nums.length - 1; j++) {
      for (var k = j + 1; k < nums.length ; k++) {
        var temp = (nums[i] - nums[j]) * nums[k]
        rus = Math.max(rus, temp)
      }
    }
  }
  return rus
};
