import sys

lines = sys.stdin.read().splitlines()

text = lines[0]
k = int(lines[1])

def computeFrequency(text, k):
    frequencies = {}
    for i in range(0, 4**k):
        frequencies[i] = 0

    glue = ""
    out = ""
    for y in range(0, len(text) - k+1):
        pattern = text[y:y+k]
        n = patternToNumber(pattern)
        frequencies[n] = frequencies[n] + 1
        #print(("%s:%d") % (pattern,n))

    out = ""
    glue = ""
    for key in frequencies:
        out += glue + str(frequencies[key])
        glue = " "

    return out

def symbolToNumber(symbol):
    bases = { "A": 0, "C": 1, "G": 2, "T": 3 }
    if symbol in bases:
        return bases[symbol]
    else:
        return None

def patternToNumber(pattern):
    if pattern == "":
        return 0
    # Get last char of pattern
    symbol = pattern[-1:]
    # Now remove last char of pattern
    pattern = pattern[0:-1]
    # compute number based on base 4 using recurusion
    number = 4 * patternToNumber(pattern) + symbolToNumber(symbol)

    return number

def numberToPattern(number, k):
    bases = { 0: "A", 1: "C", 2: "G", 3: "T" }

    n = number
    s = ""
    for i in range(0, k):
        n, r = divmod(n, 4)
        s += str(r)

    return s[::-1]

print(computeFrequency(text, k))
