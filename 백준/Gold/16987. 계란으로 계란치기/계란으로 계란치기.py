n = int(input())
eggs = []
for i in range(n):
    s, w = map(int, input().split())
    eggs.append([s, w])

cnt = 0 # 깨져있는 계란 수
ans = 0


# k번째 계란으로 다른걸 칠 차례
def hit(k):
    global ans, cnt

    # 가장 최근에 든 계란이 가장 오른쪽에 위치한 경우
    if k == n:
        ans = max(ans, cnt)
        return
    # 현재 들고 있는 계란이 깨져있거나, 깨지지않은 다른 계란이 없다면
    if eggs[k][0] <= 0 or cnt == n-1:
        hit(k+1)
        return

    for i in range(n):
        if k != i and eggs[i][0] > 0:
            eggs[k][0] -= eggs[i][1]
            eggs[i][0] -= eggs[k][1]

            x = (eggs[k][0] <= 0) + (eggs[i][0] <= 0)

            cnt += x
            hit(k+1)
            cnt -= x

            eggs[k][0] += eggs[i][1]
            eggs[i][0] += eggs[k][1]

hit(0)
print(ans)