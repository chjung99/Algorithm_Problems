from collections import deque
from math import gcd

N = int(input())
arr_A = list(map(int, input().split()))
M = int(input())
arr_B = list(map(int, input().split()))

A = 1
B = 1
for x in arr_A:
    A *= x
for x in arr_B:
    B *= x

ans = str(gcd(A, B))
print(ans[-9:])