import sys

lines = sys.stdin.read().splitlines()

text = lines[0]
k = int(lines[1])

def frequentWords(text, k):
    patterns = {}
    maximum = 0
    for i in range(0, len(text) - k):
        w = text[i:i+k]
        if w in patterns:
            patterns[w] += 1
        else:
            patterns[w] = 1

        if patterns[w] > maximum:
            maximum = patterns[w]

    glue = ""
    words = ""
    for word in patterns:
        if patterns[word] == maximum:
            words += glue + word
            glue = " "

    return words

print(frequentWords(text, k))
