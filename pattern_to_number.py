import sys

lines = sys.stdin.read().splitlines()

index = int(lines[0])
k = int(lines[1])

# convert number to its base4 value
def numberToSymbol(index):
    symbols = { 0: "A", 1: "C", 2: "G", 3: "T" }
    if index in symbols:
        return symbols[index]
    else:
        return None

# convert symbol to its base4 value
def symbolToNumber(symbol):
    bases = { "A": 0, "C": 1, "G": 2, "T": 3 }
    if symbol in bases:
        return bases[symbol]
    else:
        return None


'''
When computing Pattern = NumberToPattern(9904, 7), we divide 9904 by 4
to obtain a quotient of 2476 and a remainder of 0. This remainder represents
the final nucleotide of Pattern, or NumberToSymbol(0) = A.  We then iterate
this process, dividing each subsequent quotient by 4, until we obtain a quotient
of 0.  The symbols in the nucleotide column, read upward from the bottom,
yield Pattern = GCGGTAA.
##############################################################################
n               Quotient(n,4)           Remainder(n,4)          NumberToSymbol
9904            2476                    0                       A
2476             619                    0                       A
 619             154                    3                       T
 154              38                    2                       G
  38               9                    2                       G
   9               2                    1                       C
   2               0                    2                       G
##############################################################################
'''
# convert base4 number to its nucleotide base
def numberToPattern(index, k):
    if k == 1:
        return numberToSymbol(index)

    prefixIndex, r = divmod(index, 4)
    prefixPattern = numberToPattern(prefixIndex, k - 1)
    symbol = numberToSymbol(r)

    return prefixPattern + symbol

# convert pattern to base4 number
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

#print(patternToNumber(text))
print(numberToPattern(index, k))
