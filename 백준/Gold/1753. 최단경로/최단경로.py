import heapq

V, E = map(int, input().split())
K = int(input())
adj_list = [[] for _ in range(V + 1)]

for _ in range(E):
    u, v, w = map(int, input().split())
    adj_list[u].append([v, w])

d = [float("inf") for _ in range(V + 1)]

d[K] = 0

heap = []
heapq.heappush(heap, [0, K])

while heap:
    cur_d, cur = heapq.heappop(heap)
    if d[cur] != cur_d:
        continue
    for nxt, weight in adj_list[cur]:
        if d[nxt] > d[cur] + weight:
            d[nxt] = d[cur] + weight
            heapq.heappush(heap, [d[nxt], nxt])


for i in range(1, V+1):
    if d[i] == float("inf"):
        print("INF")
    else:
        print(d[i])
# sample input
# 6 8
# 1
# 1 2 3
# 1 3 2
# 1 4 5
# 2 3 2
# 2 5 8
# 3 4 2
# 4 5 6
# 5 6 1