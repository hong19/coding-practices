class Solution(object):
    def isPalindrome(self, x: int) -> bool:
        """
        :type x: int
        :rtype: bool
        """
        if x < 0:
            return False

        result = self.intToString(x)
        return self.isPalindromeStr(result)


    def intToString(self, x):
        """
        the string is reversed
        :param x:
        :return:
        """
        import math
        result = ""
        while x > 0:
            digit = x % 10
            result += str(digit)
            x = math.floor(x/10)

        return result

    def isPalindromeStr(self, s: str):
        idx = 0
        while idx < len(s)/2:
            if s[idx] != s[-idx - 1]:
                return False
            idx += 1

        return True
