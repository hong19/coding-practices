from solution import Solution


def test_productExceptSelf():
    c = Solution()

    nums = [1, 2, 3, 4]
    ans = c.productExceptSelf(nums)
    assert ans == [24, 12, 8, 6]

    nums = [-1, 1, 0, -3, 3]
    ans = c.productExceptSelf(nums)
    assert ans == [0, 0, 9, 0, 0]
