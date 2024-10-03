N = int(input())
pq = []
for _ in range(N):
    tmp = list(map(int, input().split()))
    pq.extend(tmp)
    pq.sort(reverse=True)
    pq = pq[:N]
print(pq[-1])


