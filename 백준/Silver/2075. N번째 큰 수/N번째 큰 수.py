import heapq

N = int(input())
pq = []
for _ in range(N):
    tmp = list(map(int, input().split()))
    for x in tmp:
        if len(pq) < N:
            heapq.heappush(pq, x)
        else:
            if pq[0] < x:
                heapq.heapreplace(pq, x)
print(pq[0])



