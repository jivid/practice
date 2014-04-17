#!/usr/bin/env python

'''
Check if two strings are anagrams. This solution involves
XORing every character in both strings. If the final value
after all the operations is not 0, they are not anagrams
'''

first = raw_input("First word: ")
second = raw_input("Second word: ")
sentences = [first, second]
x = 0

for s in sentences:
    for c in s:
        # Convert the char to binary
        b = int(bin(ord(c))[2:])
        
        # XOR the current char with the existing XOR count
        x = x^b

if not (x == 0):
    print "Not anagrams!"
else:
    print "Anagrams!"   
