V, E = map(int, input().split())
edges = []
answer = 0
cnt = 0
parent = [i for i in range(V + 1)]


def find_parent(x):
    if parent[x] != x:
        parent[x] = find_parent(parent[x])
    return parent[x]


def union_parent(a, b):
    global parent
    a = find_parent(a)
    b = find_parent(b)
    if a < b:
        parent[b] = a
    else:
        parent[a] = b


def is_diff_group(x, y):
    return parent[x] != parent[y]


for _ in range(E):
    A, B, C = map(int, input().split())
    edges.append([A, B, C])


edges.sort(key=lambda x: x[2])


# for i in range(E):
#     if is_diff_group(edges[i][0], edges[i][1]):
#         union_parent(edges[i][0], edges[i][1])
#         answer += edges[i][2]
#         cnt += 1
#     if cnt == V-1:
#         break
for s, e, w in edges:
    sRoot = find_parent(s)
    eRoot = find_parent(e)
    if sRoot != eRoot:
        if sRoot > eRoot:
            parent[sRoot] = eRoot
        else:
            parent[eRoot] = sRoot
        answer += w

print(answer)
