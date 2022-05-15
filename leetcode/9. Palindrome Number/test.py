from main import Solution


def test_isPalindromeStr():
    c = Solution()

    s = "abc"
    ans = c.isPalindromeStr(s)
    assert ans == False

    s = "aba"
    ans = c.isPalindromeStr(s)
    assert ans == True

    s = "c"
    ans = c.isPalindromeStr(s)
    assert ans == True

def test_intToString():
    c = Solution()

    x = 123
    result = c.intToString(x)

    assert result == "321"

def test_isPalindrome():
    c = Solution()

    x = 123
    result = c.isPalindrome(x)
    assert result is False

    x = 0
    result = c.isPalindrome(x)
    assert result is True

    x = 1221
    result = c.isPalindrome(x)
    assert result is True
