while True:
    try:
        answer = ""
        a = input()
        b = input()
        k = ord('z') - ord('a') + 1
        cnt_a = [0 for _ in range(k)]
        cnt_b = [0 for _ in range(k)]

        for i in range(len(a)):
            idx_a = ord(a[i]) - ord('a')
            cnt_a[idx_a] += 1
        for i in range(len(b)):
            idx_b = ord(b[i]) - ord('a')
            cnt_b[idx_b] += 1

        for i in range(k):
            min_cnt = min(cnt_a[i], cnt_b[i])
            x = chr(ord('a') + i)
            answer += x * min_cnt
        print(answer)


    except:
        break
