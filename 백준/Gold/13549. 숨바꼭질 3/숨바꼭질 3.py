import heapq

N, K = map(int, input().split())
dist = [float("inf") for _ in range(100000 + 1)]


def dijkstra(start):
    pq = []
    dist[start] = 0
    heapq.heappush(pq, (0, start))

    while pq:
        cur_dist, cur = heapq.heappop(pq)
        if cur_dist > dist[cur]:
            continue

        for i in range(3):
            if i == 0:
                nxt = cur - 1
            elif i == 1:
                nxt = cur + 1
            else:
                nxt = cur * 2

            if nxt > 100000 or nxt < 0:
                continue

            if i < 2 and dist[nxt] > cur_dist + 1:
                dist[nxt] = cur_dist + 1
                heapq.heappush(pq, (dist[nxt], nxt))
            if i == 2 and dist[nxt] > cur_dist:
                dist[nxt] = cur_dist
                heapq.heappush(pq, (dist[nxt], nxt))

dijkstra(N)
print(dist[K])