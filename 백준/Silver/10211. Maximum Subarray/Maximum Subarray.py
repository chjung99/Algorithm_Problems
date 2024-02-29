T = int(input())
ans = []


def get_maximum_subarray(arr_size, arr):
    max_value = arr[0]

    for i in range(arr_size):
        sub_arr_sum = 0
        for j in range(i, arr_size):
            sub_arr_sum += arr[j]
            if max_value < sub_arr_sum:
                max_value = sub_arr_sum

    return max_value


for _ in range(T):
    N = int(input())
    input_arr = list(map(int, input().split()))
    ans.append(get_maximum_subarray(N, input_arr))
for k in ans:
    print(k)
