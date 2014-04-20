class Node(object):
    def __init__(self, value, next=None):
        self.value = value
        self.next = next

class LinkedList(object):
    def __init__(self, root):
        self.root = Node(root)
        self.end = self.root

    def add_to_back(self, value):
        n = Node(value)
        self.end.next = n
        self.end = n

    def show(self):
        to_show = self.root
        l = ""
        while not to_show.next == None:
            l += "%d -> " % to_show.value
            to_show = to_show.next

        l += str(to_show.value)
        print l
