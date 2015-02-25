import sys

length = int(sys.argv[1])
added = ','
for i in range(length):
    if i + 1 == length:
        added = ''
    print str([1.0] * (i + 1) + [0.0] * (length - i - 1)) + added
