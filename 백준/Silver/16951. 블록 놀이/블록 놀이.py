from collections import deque

N, K = map(int, input().split())
A = list(map(int, input().split()))
answer = -1


def make_arr(idx, A, K, N):
    flag = True
    queue = deque([A[idx]])
    tmp = A[idx]
    for i in range(idx):
        tmp -= K
        if tmp <= 0:
            flag = False
        queue.appendleft(tmp)
    tmp = A[idx]
    for i in range(idx+1, N):
        tmp += K
        if tmp <= 0:
            flag = False
        queue.append(tmp)
    return list(queue), flag

def get_diff(A, B):
    cnt = 0
    for i in range(N):
        if A[i] != B[i]: cnt+=1
    return cnt


for i in range(N):
    B, is_valid = make_arr(i, A, K, N)
    if is_valid:
        diff = get_diff(A, B)
        if answer == -1:
            answer = diff
        else:
            answer = min(answer, diff)

print(answer)
