-- 코드를 입력하세요
SELECT A.flavor
FROM FIRST_HALF as A
JOIN ICECREAM_INFO as B
ON A.flavor = B.flavor
WHERE TOTAL_ORDER > 3000 and INGREDIENT_TYPE = "fruit_based" 
ORDER BY A.TOTAL_ORDER DESC;  -- 오름차순 정렬
