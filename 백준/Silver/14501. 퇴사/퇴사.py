n = int(input())
t = [0]
p = [0]
ans = 0

for i in range(n):
    ti, pi = map(int, input().split())
    t.append(ti)
    p.append(pi)


def func(day, off, pay):
    global ans

    if day == n + 1:
        ans = max(ans, pay)
        return

    if off != 0:
        func(day + 1, off - 1, pay)
    else:
        func(day + 1, off, pay)
        if t[day] - 1 <= n - day:
            func(day + 1, t[day]-1, pay + p[day])


func(1, 0, 0)
print(ans)