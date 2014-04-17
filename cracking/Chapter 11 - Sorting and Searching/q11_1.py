"""
You are given two sorted arrays, A and B, where A has enough buffer at the end
to hold B. Write a method to merge B into A in sorted order
"""
def merge(A, B, startA=0, startB=0):
    if (startB == len(B)):
        return A, B

    to_compare = A[startA]
    to_insert = B[startB]

    while to_compare < to_insert:
        startA += 1
        if (startA == len(A)):
            A = A + B[startB:]
            return A, B

        to_compare = A[startA]

    insert_at = startA if startA > 0 else 0
    A.insert(insert_at, to_insert)

    A, B = merge(A, B, startA, startB+1)

    return A, B

A = [1, 5, 7, 9]
B = [2, 6, 12]
print merge(A, B)[0]
