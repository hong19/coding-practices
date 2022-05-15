class Solution(object):
    def isPalindrome(self, x: int) -> bool:
        """
        :type x: int
        :rtype: bool
        """
        if x < 0:
            return False

        exp = self.get_exp(x)
        while x > 0:
            equal, x = self.isFirstAndLastDigitSame(x, exp)
            if not equal:
                return False
            exp = exp/100

        return True


    def get_exp(self, x: int) -> int:
        exp = 1
        x = int(x/10)
        while x > 0:
            x = int(x/10)
            exp = exp*10

        return exp


    def isFirstAndLastDigitSame(self, x: int, exp):
        first_digit = int(x/exp)
        last_digit = x % 10

        x = x - first_digit*exp
        x = int(x / 10)

        return first_digit == last_digit, x

