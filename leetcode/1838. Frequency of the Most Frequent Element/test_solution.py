from solution import Solution


def test_maxWidthRamp():
    c = Solution()

    nums = [1,2,4]
    k = 5
    ans = c.maxFrequency(nums, k)
    assert ans == 3 

    nums = [1,4,8,13]   
    k = 5
    ans = c.maxFrequency(nums, k)
    assert ans == 2 

    nums = [3,9,6]
    k = 2
    ans = c.maxFrequency(nums, k)
    assert ans == 1 
