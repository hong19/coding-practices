class Solution(object):
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        if x < 0:
            return False

        s = self.intToString(x)


    def intToString(self, x):
        s = ""
        while x > 0:
            digit = x % 10
            s += digit
            x = x/10

        return s

    def isPalindromeStr(self, s):
        idx = 0
        while len(s):
            if s[i] != s[:-i]:




