class Solution:
    def isPowerOfThree(self, n: int) -> bool:
        if n%3 == 0 and n!=0:
            return self.isPowerOfThree(n//3)
        if n==1:
            return True
        return False