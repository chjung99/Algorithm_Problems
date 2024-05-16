board = []
call_num = []
visit = [[0 for _ in range(5)] for __ in range(5)]

for _ in range(5):
    board.append(list(map(int, input().split())))
for _ in range(5):
    call_num.extend(list(map(int, input().split())))


def find_and_visit(num, board, visit):
    for i in range(5):
        for j in range(5):
            if board[i][j] == num:
                visit[i][j] = 1
                return


def get_bingo_cnt(board):
    bingo_cnt = 0
    for i in range(5):
        visit_cnt = 0
        for j in range(5):
            if visit[i][j]:
                visit_cnt += 1
            else:
                break
        if visit_cnt == 5:
            bingo_cnt += 1

    for i in range(5):
        visit_cnt = 0
        for j in range(5):
            if visit[j][i]:
                visit_cnt += 1
            else:
                break
        if visit_cnt == 5:
            bingo_cnt += 1

    cx, cy = 0, 0
    visit_cnt = 0
    for i in range(5):
        if visit[cx][cy]:
            visit_cnt += 1
            cx += 1
            cy += 1
        else:
            break
    if visit_cnt == 5:
        bingo_cnt += 1

    cx, cy = 0, 4
    visit_cnt = 0
    for i in range(5):
        if visit[cx][cy]:
            visit_cnt += 1
            cx += 1
            cy -= 1
        else:
            break
    if visit_cnt == 5:
        bingo_cnt += 1

    return bingo_cnt


for idx, num in enumerate(call_num):
    find_and_visit(num, board, visit)

    if get_bingo_cnt(board) >= 3:
        print(idx+1)
        break
