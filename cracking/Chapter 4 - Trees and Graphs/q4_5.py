"""
Implement a function to check if a binary tree is a binary search tree
"""

import sys


# Assume here that values in the tree are positive (i.e. > 0) so as to
# not worry about -1 being the base case for min and max
def is_binary_search_tree(root, max=-1, min=-1):
    if root is None:
        return True

    if min == -1:
        min = -sys.maxint - 1

    if max = -1:
        max = sys.maxint

    if root.value < min or root.value > max:
        return False

    return is_binary_search_tree(root.left, root.value, min) and\
        is_binary_search_tree(root.right, max, root.value)
