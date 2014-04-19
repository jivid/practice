"""
Question
========
Design an algorithm and write code to find the first common ancestor of two
nodes in a binary tree. Avoid storing additional nodes in a data structure.
NOTE: This is not necessarily a binary search tree

Solution
========
The algorithm implemented is as follows:
    - If both values are in the left sub tree, recurse with left node as root
    - If both values are in the right sub tree, recurse with right node as root
    - If both values aren't contained in the same sub-tree, return current node

The helper function exists_in checks to see if a given value belongs as an
ancestor of a supplied node
"""

def exists_in(node, value):
    if node is None:
        return False

    if node.value == value:
        return True

    if node.left is not None:
        return exists_in(node.left, value)
    elif node.right is not None:
        return exists_in(node.right, value)
    else:
        return False

def find_common_ancestor(root, val1, val2):
    if not exists_in(root, val1) or not exists_in(root, val2):
        return None

    if exists_in(root.left, val1) and exists_in(root.left, val2):
        return find_common_ancestor(root.left, val1, val2)
    elif exists_in(root.right, val1) and exists_in(root.right, val2):
        return find_common_ancestor(root.right, val1, val2)
    else:
        return root.value
