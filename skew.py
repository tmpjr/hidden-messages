import sys
import utils

lines = sys.stdin.read().splitlines()

text = lines[0]

def computeSkew(text):
    skews = [0]
    skew = 0
    for i in range(0, len(text)):
        w = text[i]
        if w == "C":
            skew = skew - 1
        elif w == "G":
            skew = skew + 1
        else:
            skew = skew
        skews.append(skew)

    return skews


print(" ".join(str(s) for s in computeSkew(text)))
#print(" ".join(skew))
