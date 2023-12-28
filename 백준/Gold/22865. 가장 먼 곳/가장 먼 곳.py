import sys,heapq
input = sys.stdin.readline
INF = int(1e9)

n = int(input())
A,B,C = map(int,input().split())
g = [[] for _ in range(n+1)]
m = int(input())
for _ in range(m):
    a,b,c = map(int,input().split())
    g[a].append((b,c))
    g[b].append((a,c))

def dijstra(start):
    dis = [INF] * (n+1)
    dis[start] = 0 
    q = [(0,start)]

    while q:
        now_dist,now_node = heapq.heappop(q)
        if dis[now_node] != now_dist : continue

        for next_node,next_dist in g[now_node]:
            if dis[now_node] + next_dist < dis[next_node]:
                dis[next_node] = dis[now_node] + next_dist
                heapq.heappush(q,(dis[next_node],next_node))
    
    return dis

dijstra_a = dijstra(A)
dijstra_b = dijstra(B)
dijstra_c = dijstra(C)

MAX = 0; answer = 0
for i in range(1,n+1):
    
    if MAX < min(dijstra_a[i],dijstra_b[i],dijstra_c[i]):
        MAX = min(dijstra_a[i],dijstra_b[i],dijstra_c[i])
        answer = i

print(answer)






