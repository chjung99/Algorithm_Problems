from collections import deque

N = int(input())
adj_list = [[] for _ in range(N + 1)]
visit = [0 for _ in range(N + 1)]


def bfs(cur):
    global visit
    visit[cur] = 1
    queue = deque([cur])

    while queue:
        cur = queue.popleft()
        for nxt in adj_list[cur]:
            if not visit[nxt]:
                visit[nxt] = 1
                queue.append(nxt)
    return


for _ in range(N - 2):
    s, e = map(int, input().split())
    adj_list[s].append(e)
    adj_list[e].append(s)

for i in range(1, N + 1):
    if not visit[i]:
        print(i, end=" ")
        bfs(i)
