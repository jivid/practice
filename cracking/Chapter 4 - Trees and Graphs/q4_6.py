"""
Write an algorithm to find the 'next' node (i.e. in-order successor) of a given
node in a binary search tree. You may assume that each node has a link to its
parent
"""

def find_next(value, node, root=None):
    if value == node.value:
        if node.right is not None:
            return node.right.left.value if node.right.left is not None
                else node.right.value
        else:
            return root

    if value < node.value:
        return find_next(value, node.left, node.value)
    elif value > node.value:
        return find_next(value, node.right, root)
