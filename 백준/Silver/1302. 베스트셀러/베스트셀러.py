from collections import defaultdict
# 입력
N = int(input())
# 딕셔너리 정의 (문자열, 개수)
books = defaultdict(int)
# 문자열 입력
for _ in range(N):
    book_name = input()
    books[book_name] += 1
# 정렬
sorted_books = sorted(books.items(),key=lambda x:(-x[1], x[0]))
# 출력
print(sorted_books[0][0])