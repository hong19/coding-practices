from solution import Solution


def test_maxWidthRamp():
    c = Solution()

    nums = [6,0,8,2,1,5]
    ans = c.maxWidthRamp(nums)
    assert ans == 4

    nums = [9,8,1,0,1,9,4,0,4,1]
    ans = c.maxWidthRamp(nums)
    assert ans == 7 

    nums = [0, 1]
    ans = c.maxWidthRamp(nums)
    assert ans == 1 
