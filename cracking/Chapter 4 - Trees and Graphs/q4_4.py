"""
Given a binary tree, design an algorithm which creates a linked list of all
the nodes at each depth. (eg. if you have a tree with depth D, you'll have D
linked lists)
"""

def convert_to_linked_lists(tree):
    # generator to yield a linked list for each level of the tree
    nodes = [tree.root]
    while not (len(nodes) == 0):
        children = []
        l = LinkedList()
        for node in nodes:
            l.add_to_back(node)
            children += [c for c in node.children]
            nodes.remove(node)
        nodes = children
        yield l

