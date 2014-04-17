"""
Given a sorted array of n integers that has been rotated an unknown number of
times, write code to find an element in the array. You may assume that the
array was originally sorted in increasing order
"""

def find(value, array):
    found = False
    start = 0
    end = len(array) - 1

    while not found:
        print "Searching in %s" % array
        middle = len(array)/2
        index = start + middle

        if array[middle] == value:
            return index

        left = array[:middle]
        right = array[middle:]

        if value >= left[0] and value <= left[-1]:
            array = left
            end = middle
        elif value >= right[0] and value <= right[-1]:
            array = right
            start = index
        else:
            return -1

to_find = int(raw_input("Number: "))
print find(to_find, [31,37,44,59,1,6,11,17,23,28])
