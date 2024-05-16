from itertools import permutations

def solution(k, dungeons):
    answer = 0
    routes = permutations(dungeons, len(dungeons))
    
    for route in routes:
        user_cost = k
        visit_cnt = 0
        for first_cost, second_cost in route:
            if first_cost > user_cost:
                break
            else:
                user_cost -= second_cost
                visit_cnt += 1
        answer = max(answer, visit_cnt)
    return answer