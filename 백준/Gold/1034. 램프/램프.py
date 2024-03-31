n, m = map(int, input().split())
board = []
for _ in range(n):
    board.append(list(input()))
k = int(input())
answer = 0
for col in range(n):
    col_zero_cnt = board[col].count("0")

    cnt = 0
    if col_zero_cnt <= k and col_zero_cnt % 2 == k % 2:
        for col2 in range(n):
            if board[col] == board[col2]:
                cnt += 1
    answer = max(answer, cnt)

print(answer)

