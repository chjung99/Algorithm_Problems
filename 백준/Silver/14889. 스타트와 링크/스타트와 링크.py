n = int(input())

team = []
pick = []
visit = [0 for i in range(21)]
arr = list(range(n))
v = []
score = 0
flag = 0
for i in range(n):
    team.append(list(map(int, input().split())))


def func(idx):
    global v, flag, score
    if len(v) == n // 2:

        rv = [arr[i] for i in range(len(arr)) if visit[i] == 0]
        v_score = 0
        for i in range(len(v)):
            for j in range(len(v)):
                v_score += team[v[i]][v[j]]
        rv_score = 0
        for i in range(len(rv)):
            for j in range(len(rv)):
                rv_score += team[rv[i]][rv[j]]
        if flag == 0:
            flag = 1
            if v_score-rv_score < 0:
                score = rv_score - v_score
            else:
                score = v_score - rv_score
        else:
            if v_score - rv_score < 0:
                score = min(score, rv_score - v_score)
            else:
                score = min(score, v_score - rv_score)
        return
    for i in range(idx, len(arr)):
        if visit[i] == 0:
            v.append(arr[i])
            visit[i] = 1
            func(i)
            v.pop()
            visit[i] = 0
    return


func(0)
print(score)