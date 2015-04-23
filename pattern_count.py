import sys

lines = sys.stdin.read().splitlines()

text = lines[0]
pattern = lines[1]

def patternCount(text, pattern):
    count = 0
    for i in range(0, len(text) - len(pattern)):
        if text[i:i+len(pattern)] == pattern:
            count += 1

    return count

count = patternCount(text, pattern)
print(count)

