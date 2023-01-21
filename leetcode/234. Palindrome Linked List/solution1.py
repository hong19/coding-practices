# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        num = 0
        while head.next:
            num = num*10 + head.val
            head = head.next

        return self.isPalindromeInt(num)

    def isPalindromeInt(self, x: int) -> bool:
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
