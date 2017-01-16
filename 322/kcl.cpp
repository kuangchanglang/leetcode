#include<vector>
#include<iostream>
using namespace std;

class Solution {
public:
    int coinChange(vector<int>& coins, int amount) {
        const int INF = 0x7fffffff;
        vector<int> dp(amount+1, INF);
        dp[0] = 0;
        for(int i=0;i<=amount; i++){
            for(int j=0;j<coins.size();j++){
                int idx = i-coins[j];
                if(idx<0){
                    continue;
                }
                // dp[idx]>0 reachable
                if(dp[idx]>=0 && dp[idx]<dp[i]){
                    dp[i] = dp[idx]+1;
                }
            }
        }
        return dp[amount]==INF ? -1: dp[amount];
    }
};
