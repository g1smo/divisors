#!/usr/bin/python

import math

# Get number of divisors (cache results)
numDivisors = {
    0: set([1]),
    1: set([1]),
    2: set([1, 2])
}

def numDiv(num):
    if num in numDivisors:
        return numDivisors[num]

    count = 2
    limit = int(math.sqrt(num))
    while (num % count) != 0:
        # Is number prime?
        if count >= limit:
            numDivisors[num] = set([num, 1])
            return numDivisors[num]

        count += 1

    divisors = numDiv(int(num / count))
    divisors = divisors.union(set([x * count for x in divisors]))

    numDivisors[num] = divisors
    return numDivisors[num]

def omega(num):
    return sum(numDiv(num))

def omDiv(maxNum, divider):
    results = []

    for i in range(1, maxNum + 1):
        if i % int(maxNum / 100) == 0:
            print(i / (maxNum / 100), "%")

        if omega(i) % divider == 0:
            results.append(i)

    return sum(results)

print(omDiv(10**6, 2017))
