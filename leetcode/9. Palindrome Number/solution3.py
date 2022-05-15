class Solution(object):
    def isPalindrome(self, x: int) -> bool:
        """
        :type x: int
        :rtype: bool
        """
        if x < 0 or (x % 10 == 0 and x != 0):
            return False

        reversed_x = 0
        while x > reversed_x:
            reversed_x = 10*reversed_x + x % 10
            x = int(x / 10)

        return x == reversed_x or x == int(reversed_x/10)
