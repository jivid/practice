"""
Write code to partition a linked list around a value x, such that all nodes
less than x come before all nodes greater than or equal to x
"""

from linked_list import LinkedList


def partition(l, x):
    lesser_list = None
    greater_list = None

    check = l.root
    while not check.next is None:
        if check.value < x:
            if lesser_list is None:
                lesser_list = LinkedList(check.value)
            else:
                lesser_list.add_to_back(check.value)
        elif check.value >= x:
            if greater_list is None:
                greater_list = LinkedList(check.value)
            else:
                greater_list.add_to_back(check.value)
        check = check.next

    if check.value < x:
        if lesser_list is None:
            lesser_list = LinkedList(check.value)
        else:
            lesser_list.add_to_back(check.value)
    elif check.value >= x:
        if greater_list is None:
            greater_list = LinkedList(check.value)
        else:
            greater_list.add_to_back(check.value)

    final = lesser_list
    to_add = greater_list.root
    while not to_add.next is None:
        final.add_to_back(to_add.value)
        to_add = to_add.next

    final.add_to_back(to_add.value)
    return final 
