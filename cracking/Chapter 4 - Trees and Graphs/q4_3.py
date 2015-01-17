class Node(object):
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None


def make_tree(arr):
    if len(arr) == 1:
        return Node(arr[0])

    if len(arr) == 2:
        root = Node(arr[0])
        root.right = Node(arr[1])
        return root

    mid = len(arr)/2
    left = arr[:mid]
    right = arr[mid+1:]

    root = Node(arr[mid])
    root.left = make_tree(left)
    root.right = make_tree(right)
    return root
