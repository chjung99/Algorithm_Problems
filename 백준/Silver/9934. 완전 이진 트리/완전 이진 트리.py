k = int(input())
tree = list(map(int, input().split()))
ans = [[] for _ in range(12)]
def move(cur, depth):
    if depth >= 2:
        left_child = cur - (2 ** (depth - 2))
        right_child = cur + (2 ** (depth - 2))
        ans[depth - 1].append(tree[left_child])
        ans[depth - 1].append(tree[right_child])
        # print(f"{} {tree[right_child]}", end=" ")
        move(left_child, depth-1)
        move(right_child, depth-1)
    else:
        return
root = ((2 ** k) -1) // 2
ans[k].append(tree[root])
move(root, k)
for d in range(k, 0, -1):
    for t in ans[d]:
        print(t, end=" ")
    print()