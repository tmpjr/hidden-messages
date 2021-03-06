import sys

lines = sys.stdin.read().splitlines()

genome = lines[0]
parts = lines[1].split()
k = int(parts[0])
L = int(parts[1])
t = int(parts[2])

#print(("k: %d, L: %d, t: %d") % (k, L, t))

def findClumps(genome, k, L, t):
    clumps = ""
    data = {}
    for i in range(0, len(genome) - L):
        # text is a genome of length L
        text = genome[i:i+L]
        #print(text)
        d = {}
        for y in range(0, len(text) - k):
            c = text[y:y+k]
            if c in d:
                d[c] += 1
            else:
                d[c] = 1

        for key in d:
            if d[key] >= t:
                data[key] = d[key]

    glue = ""
    for clump in data:
        clumps += glue + clump
        glue = " "

    return clumps

clumps = findClumps(genome, k, L, t)
print(clumps)
