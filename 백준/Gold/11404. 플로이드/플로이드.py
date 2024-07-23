n = int(input())
m = int(input())

matrix = [[float('inf') for _ in range(n+1)]for __ in range(n+1)]

for _ in range(m):
    a, b, c = map(int, input().split())
    matrix[a][b] = min(matrix[a][b], c)
for i in range(1, n+1):
    matrix[i][i] = 0

for k in range(1, n + 1):
    for i in range(1, n+1):
        for j in range(1, n+1):
            matrix[i][j] = min(matrix[i][j], matrix[i][k] + matrix[k][j])

for i in range(1, n+1):
    for j in range(1, n+1):
        if matrix[i][j] == float('inf'):
            print(0, end=" ")
        else:
            print(matrix[i][j], end=" ")
    print("")