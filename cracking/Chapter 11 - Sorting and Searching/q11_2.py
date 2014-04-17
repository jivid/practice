"""
Write a method to sort an array of strings such that anagrams are next to
each other
"""

SEEN_DICT = {}

to_ord = lambda x: [ord(c) for c in x]

def find_anagram(word):
    for k in SEEN_DICT:
        if sorted(to_ord(k)) == sorted(to_ord(word)):
            return k

    return None

def sort(array):
    finished = []
    for word in array:
        anagram = find_anagram(word)
        if anagram is None:
            SEEN_DICT[word] = [word]
        else:
            SEEN_DICT[anagram].append(word)

    for k in SEEN_DICT:
        finished = finished + SEEN_DICT[k]

    return finished

to_sort = ["divij", "anuj", "jivid", "nujra", "apple", "orange", "juna",
            "arjun", "racecar", "ppale", "pale", "viidj"]
print sort(to_sort)
