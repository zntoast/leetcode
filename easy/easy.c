#include <stdio.h>
int LargestSumAfterKNegations(int* nums, int numsSize, int k);

int main(){
	int nums[3] = {4,2,3};
	int ans = 0;
	ans = LargestSumAfterKNegations(nums, 3, 1);
	printf("和为：%d\n",ans);
}
/*
1005. K 次取反后最大化的数组和
*/
int LargestSumAfterKNegations(int* nums, int numsSize, int k){
	int result = 0;
	while (k--) {
		int min_index = 0;
		for (int i=0;i<numsSize;i++){
			if (nums[min_index] > nums[i]){
				min_index = i;
			}
		}
		nums[min_index] = -nums[min_index];
	}
	for (int i = 0;i<numsSize;i++){
		result += nums[i];
	}
	return  result;
}
