'''
Let's follow the 5' → 3' direction of DNA and walk along the chromosome from terC to oriC (along a reverse half-strand), then continue on from oriC to terC (along a forward half-strand). In our previous discussion, we saw that the skew is decreasing along the reverse half-strand and increasing along the forward half-strand. Thus, the skew should achieve a minimum at the position where the reverse half-strand ends and the forward half-strand begins, which is exactly the location of oriC!

We have just developed an insight for a new algorithm for locating oriC: it should be found where the skew attains a minimum.

Minimum Skew Problem: Find a position in a genome minimizing the skew.
     Input: A DNA string Genome.
     Output: All integer(s) i minimizing Skewi (Genome) among all values of i (from 0 to |Genome|).

CODE CHALLENGE: Solve the Minimum Skew Problem.

Try the code-graded problem!

Sample Input:
TAAAGACTGCCGAGAGGCCAACACGAGTGCTAGAACGAGGGGCGTAAACGCGGGTCCGAT

Sample Output:
11 24
'''
import sys
import utils

lines = sys.stdin.read().splitlines()
text = lines[0]


skew = utils.computeSkew(text)
minval = min(skew)
s = ""
for i in range(0, len(skew)):
    if skew[i] == minval:
        s += str(i) + " "

print(s)
