import heapq

def solution(scoville, K):
    answer = 0
    # 최소값이 K 이상이면 됨
    # 우선순위 큐에 리스트 삽입
    heapq.heapify(scoville)
    while True:
        if len(scoville) == 1:
            if scoville[0] < K:
                answer = -1
            break
        elif scoville[0] >= K:
            break
        first = heapq.heappop(scoville)
        second = heapq.heappop(scoville)
        mixed = first + (second*2)
        heapq.heappush(scoville, mixed)
        answer += 1
            
    return answer