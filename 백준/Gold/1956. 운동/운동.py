V, E = map(int, input().split())
matrix = [[float('inf') for _ in range(V+1)]for __ in range(V+1)]

for _ in range(E):
    a, b, c = map(int, input().split())
    matrix[a][b] = c

for i in range(1, V+1):
    for j in range(1, V+1):
        for k in range(1, V+1):
            matrix[i][j] = min(matrix[i][j], matrix[i][k] + matrix[k][j])
answer = float('inf')
for i in range(1, V+1):
    answer = min(answer, matrix[i][i])
if answer == float('inf'):
    answer = -1
print(answer)

