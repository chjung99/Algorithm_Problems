N = int(input())
A = list(map(int, input().split()))
M = int(input())
B = list(map(int, input().split()))

A.sort()


def binary_search(x, A):
    left = 0
    right = len(A) - 1
    while left <= right:
        mid = (left + right) // 2
        if A[mid] == x:
            return 1
        elif A[mid] > x:
            right = mid - 1
        else:
            left = mid + 1
    return 0


for b in B:
    print(binary_search(b, A))