n = int(input())
arr = list(map(int, input().split()))
count = list(map(int, input().split()))
v = []
ans_min = 0
ans_max = 0
flag = 0


def pi():
    global v, count, flag, ans_min, ans_max
    if len(v) == n - 1:
        x = arr[0]
        for i in range(n-1):
            if v[i] == 0:
                x += arr[i+1]
            elif v[i] == 1:
                x -= arr[i+1]
            elif v[i] == 2:
                x *= arr[i+1]
            else:
                if x < 0:
                    x *= -1
                    x = x // arr[i+1]
                    x *= -1
                else:
                    x = x // arr[i+1]
        if flag == 0:
            flag = 1
            ans_min = x
            ans_max = x
        else:
            ans_min = min(ans_min, x)
            ans_max = max(ans_max, x)
        return
    for i in range(4):
        if count[i] > 0:
            count[i] -= 1
            v.append(i)
            pi()
            count[i] += 1
            v.pop()


pi()
print(ans_max)
print(ans_min)
