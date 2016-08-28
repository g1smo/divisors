#!/usr/bin/python

import math

# Get number of divisors (cache results)
numDivisors = {
    0: 1,
    1: 1,
    2: 3
}

def numDiv(num):
    if num in numDivisors:
        return numDivisors[num]

    count = 2
    limit = int(math.sqrt(num))
    while (num % count) != 0:
        # Is number prime?
        if count >= limit:
            numDivisors[num] = num + 1
            return numDivisors[num]

        count += 1

    divisors = numDiv(int(num / count))

    numDivisors[num] = divisors
    return numDivisors[num]

def omDiv(maxNum, divider):
    results = 0

    tick = int(maxNum/1000)
    for i in range(1, maxNum + 1):
        if i % tick == 0:
            print(i / (maxNum / 100), "%")

        if numDiv(i) % divider == 0:
            results += i

    return results

print(omDiv(10**6, 2017))
