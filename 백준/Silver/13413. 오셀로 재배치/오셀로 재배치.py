t = int(input())
for _ in range(t):
    n = int(input())
    src = input()
    dest = input()

    answer = 0
    diff_cnt = 0
    dest_cnt = 0
    src_cnt = 0
    for i in range(n):
        if src[i] == 'B':
            src_cnt += 1
        if dest[i] == 'B':
            dest_cnt += 1
        if src[i] != dest[i]:
            diff_cnt += 1
    if src_cnt == dest_cnt:
        answer = diff_cnt//2
    else:
        answer = abs(src_cnt - dest_cnt) + (diff_cnt - abs(src_cnt - dest_cnt))//2
    print(answer)