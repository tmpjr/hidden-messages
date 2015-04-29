import sys
import utils

lines = sys.stdin.read().splitlines()

index = int(lines[0])
k = int(lines[1])


#text = lines[0]
#print(utils.patternToNumber(text))
print(utils.numberToPattern(index, k))
