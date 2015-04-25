import sys
import utils

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

def frequentWordsBySorting(text, k):
    frequentPatterns = []
    index = {}
    count = {}
    for i in range(0, len(text) - k):
        pattern = text[i:i+k]
        index[i] = utils.patternToNumber(pattern)
        count[i] = 1

    sortedIndex = sorted(index.values())
    for i in range(1, len(text) - k):
        if sortedIndex[i] == sortedIndex[i-1]:
            count[i] = count[i - 1] + 1

    maxCount = max(count.values())

    for i in range(0, len(text) - k):
        if count[i] == maxCount:
            frequentPatterns.append(utils.numberToPattern(sortedIndex[i], k))

    return " ".join(frequentPatterns)

print(frequentWordsBySorting(text, k))
