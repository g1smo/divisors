#!/usr/bin/python

# Get number of divisors (cache results)
numDivisors = {}
primes = {}

def numDiv(num):
    if num in numDivisors:
        return numDivisors[num]

    if num == 0 or num == 1:
        numDivisors[num] = set([1])
        return numDivisors[num]

    count = int(num / 2)
    while count > 1 and num % count != 0:
        count -= 1

    subDivs = numDiv(count)
    factor = int(num / count)
    divisors = set([x * factor for x in subDivs]).union(subDivs)

    numDivisors[num] = divisors
    return divisors

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
