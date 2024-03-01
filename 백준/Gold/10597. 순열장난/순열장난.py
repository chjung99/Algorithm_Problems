input_string = input()
L = len(input_string)
visit = [0 for _ in range(51)]
ans = set()


def valid(chk_str):
    num = int(chk_str)
    return chk_str[0] != '0' and 1 <= num <= 50 and not visit[num]


def search(start, end, accum_str):
    if start == L:
        if sum(visit[0:len(accum_str)+1]) == len(accum_str):
            ans.add(tuple(accum_str))
        return
    if start > L:
        # print(accum_str)
        return

    if start != end:
        cur_str = input_string[start:end + 1]
    else:
        cur_str = input_string[start]

    if not valid(cur_str):
        return

    cur_num = int(cur_str)
    visit[cur_num] = 1
    accum_str.append(cur_str)

    search(end + 1, end + 1, accum_str)
    if end + 2 < L:
        search(end + 1, end + 2, accum_str)
    accum_str.pop()
    visit[cur_num] = 0


search(0, 0, [])
search(0, 1, [])
# print(ans)
print(" ".join(list(ans.pop())))
