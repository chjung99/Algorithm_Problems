n, m = map(int, input().split())

def factorial(n):
    res = 1
    for i in range(1, n+1):
        res *= i
    return res
def comb(n, m):
    return factorial(n)//(factorial(m)*factorial(n-m))


print(comb(n, m))
