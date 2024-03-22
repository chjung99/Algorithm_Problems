
def bfs(board, sx, sy):
    N = len(board)
    dx = [0, 1, 0, -1]
    dy = [1, 0, -1, 0]
    
    cnt = 0
    
    for i in range(4):
        nx, ny = sx + dx[i], sy + dy[i]
        if nx < 0 or nx >= N or ny < 0 or ny >= N:
            continue
        if board[nx][ny] == board[sx][sy]:
            cnt += 1
            
    return cnt

def solution(board, h, w):
    answer = bfs(board, h, w)
    
    return answer