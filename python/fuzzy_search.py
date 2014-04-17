WORDS = open('/usr/share/dict/words', 'r')

def search(input_word, delta):
    matched = []
    for word in WORDS:
        i, w = 0, 0
        diff_seen = 0
        larger = input_word if len(input_word) > len(word) else word
        leave = False
        for _ in larger:
            if (i == len(input_word)) or (w == len(word)):
                if diff_seen == delta:
                    leave = True
                    break
                else:
                    if i == len(input_word):
                        i = i - 1
                    elif w == len(word):
                        w = w - 1

            if not input_word[i] == word[w]:
                if diff_seen < delta:
                    if len(input_word) > len(word):
                        i = i + 1
                    elif len(word) > len(input_word):
                        w = w + 1
                    else:
                        i = i + 1
                        w = w + 1
                    diff_seen = diff_seen + 1
                else:
                    leave = True
                    break
            else:
                i = i + 1
                w = w + 1

        if leave:
            continue

        matched.append(word.replace("\n", ""))

    return matched

to_search = raw_input("Enter word to search: ")
delta = int(raw_input("Delta: "))
print search(to_search, delta) 

WORDS.close()
