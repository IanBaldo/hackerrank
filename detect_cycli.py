# https://www.hackerrank.com/challenges/ctci-linked-list-cycle

"""
Detect a cycle in a linked list. Note that the head pointer may be 'None' if the list is empty.

A Node is defined as: 
 
    class Node(object):
        def __init__(self, data = None, next_node = None):
            self.data = data
            self.next = next_node
"""


def has_cycle(head):
    
    if head == None:
        return 0 # false

    # Slow Runner
    slow = head

    while  slow.next != None:
        # Fast Runner
        fast = slow.next

        if fast == slow:
            return 1 # true

        slow = fast

    return 0 # false
    
