def node_height(node):
    if len(node.children) == 0:
        return 0

    heights = map(node_height, node.children)
    return 1 + max(heights)


class Node(object):
    def __init__(self, value=None):
        self.value = value
        self.children = []

    def add_child(self, value=None):
        if len(self.children) == 2:
            raise TypeError('Binary nodes can only have two children')

        if not isinstance(value, Node):
            value = Node(value)

        self.children.append(value)

    @property
    def balanced(self):
        if len(self.children) == 0:
            return False

        left = self.children[0]
        left_height = 1 + node_height(left)

        # If there's no second child, the node is only balanced
        # if the first child had a height of 0
        if not len(self.children) >= 2:
            return True if left_height <= 1 else False

        right = self.children[1]
        right_height = 1 + node_height(right)

        return abs(left_height - right_height) <= 1;

