#include <vector>
using namespace std;

class Solution {
public:
    int lengthOfLIS(vector<int>& nums) {
        int * dp = new int[nums.size()];
        for(int i=0;i<nums.size();i++){
            int max = 0;
            for(int j=0;j<i;j++){
                if(nums[j]<nums[i] && max<dp[j]){
                    max = dp[j];
                }
            }
            dp[i] = max+1;
        }

        int max = 0;
        for(int i=0;i<nums.size();i++){
            if(dp[i]>max){
                max = dp[i];
            }
        }
        delete dp;
        return max;
    }
};
