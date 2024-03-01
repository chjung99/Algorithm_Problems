N = int(input())
num_of_names = [0 for _ in range(26)]  # a to z
for _ in range(N):
    first_name = input()[0]
    num_of_names[ord(first_name)-ord('a')] += 1

flag = True
for i, x in enumerate(num_of_names):
    if x >= 5:
        print(chr(ord('a')+i), end="")
        flag = False
if flag:
    print("PREDAJA")



