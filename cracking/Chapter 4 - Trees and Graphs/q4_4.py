import copy

class TreeNode(object):
    def __init__(self, val, left, right):
        self.val = val
        self.left = left
        self.right = right

class Tree(object):
    def __init__(self, root):
        self.root = root


class LinkedListNode(object):
    def __init__(self, val):
        self.val = val
        self.next_node = None

    @staticmethod
    def from_nums(arr):
        n = LinkedListNode(arr[0].val)
        res = n
        for num in arr[1:]:
            n.next_node = LinkedListNode(num.val)
            n = n.next_node

        return res


def make_linked_lists(tree):
    def _populate(arr, node):
        if node.left:
            arr.append(node.left)

        if node.right:
            arr.append(node.right)

        return arr

    c1 = [tree.root]
    c2 = _populate([], tree.root)
    lists = []

    while c1:
        lists.append(LinkedListNode.from_nums(c1))
        c1 = copy.copy(c2)
        c2 = []
        for c in c1:
            c2 = _populate(c2, c)

    return lists
