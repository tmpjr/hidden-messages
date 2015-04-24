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

def patternToNumber(pattern):
    bases = { "A": 0, "C": 1, "G": 2, "T": 3 }
    reverse = pattern[::-1]
    number = 0
    for i in reversed(range(0, len(pattern))):
        base = reverse[i]
        index = bases[base]
        #print(("%d * 4**%d") % (index, i))
        number += index * 4**i

    return number

def numberToPattern(number, k):
    bases = { 0: "A", 1: "C", 2: "G", 3: "T" }



#print(frequentWords(text, k))
print(patternToNumber('ATGCAA'))
