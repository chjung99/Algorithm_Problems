s = input()
t = input()
flag = 0


def dfs(src, tar):
    global flag

    if len(src) == len(tar):
        if src == tar:
            flag = 1
        return
    if src[0] == 'B':

        dfs(src[::-1][:-1], tar)
    if src[-1] == 'A':
        dfs(src[:-1], tar)


dfs(t, s)
print(flag)
