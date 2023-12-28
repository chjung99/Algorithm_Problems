import heapq

N = int(input())
A, B, C = map(int, input().split())
M = int(input())

graph = [[] for _ in range(N + 1)]
friends = [A, B, C]
cand = []  # [[노드 번호 X, X에서 가장 가까운 친구의 집까지 거리], ..]


def dijkstra(start):
    dist = [float("inf") for _ in range(N + 1)]
    dist[start] = 0
    pq = []
    heapq.heappush(pq, [0, start])
    while pq:
        cur_dist, cur = heapq.heappop(pq)
        if cur_dist > dist[cur]:
            continue
        for nxt_dist, nxt in graph[cur]:
            if dist[nxt] > cur_dist + nxt_dist:
                dist[nxt] = cur_dist + nxt_dist
                heapq.heappush(pq, [dist[nxt], nxt])
    return dist

for i in range(M):
    D, E, L = map(int, input().split())
    graph[D].append([L, E])
    graph[E].append([L, D])

friends_dist = []
friends_dist.append(dijkstra(A))
friends_dist.append(dijkstra(B))
friends_dist.append(dijkstra(C))

for i in range(1, N + 1):
    if i not in friends:
        cand.append([i, min(friends_dist[0][i], friends_dist[1][i], friends_dist[2][i])])

cand.sort(key=lambda x: (-x[1], x[0]))
print(cand[0][0])
