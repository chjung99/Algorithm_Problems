MAX_COL = 0
MAX_ROW = 5
arr = []

for _ in range(MAX_ROW):
    tmp = list(input())
    if len(tmp) > MAX_COL:
        MAX_COL = len(tmp)
    arr.append(tmp)

col_cnt = 0
while col_cnt < MAX_COL:
    for i in range(MAX_ROW):
        if len(arr[i]) < col_cnt + 1:
            continue
        print(arr[i][col_cnt], end="")
    col_cnt += 1