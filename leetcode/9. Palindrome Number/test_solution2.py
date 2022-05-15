from solution2 import Solution
import pytest

c = Solution()


def test_get_exp():
    x = 101
    result = c.get_exp(x)
    assert result == 100

@pytest.mark.parametrize(
    "x, exp, expected_is_equal, expected_x",
    [
        pytest.param(
            101,
            100,
            True,
            0
        ),
        pytest.param(
            1,
            1,
            True,
            0
        ),
        pytest.param(
            0,
            1,
            True,
            0
        ),
        pytest.param(
            11,
            10,
            True,
            0
        ),
        pytest.param(
            1221,
            1000,
            True,
            22
        ),
    ]

)
def test_isFirstAndLastDigitSame(x, exp, expected_is_equal, expected_x):
    is_equal, new_x = c.isFirstAndLastDigitSame(x, exp)
    assert is_equal is expected_is_equal
    assert new_x == expected_x

def test_isPalindrome():
    x = 1221
    result = c.isPalindrome(x)
    assert result is True
