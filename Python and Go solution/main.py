def sqrt(n):
    if n == 0:
        return 0
    left, right = 1, n
    while left <=right:
        mid = (left+right)//2
        # print("left:", left, "right: ", right)
        # print("mid:", mid)
        if mid*mid == n:
            return mid
        elif mid*mid < n:
            left = mid+1
        else:
            right = mid-1   
    return right


def valid_palindrome(str) -> bool:
    i, j = 0, len(str)-1
    while i<j:
        var = str[i].isalnum()
        var2 = str[j].isalnum()
        if var == False:
            i = i+1
            continue
        if var2 == False:
            j = j-1
            continue
        i_low = str[i].lower()
        j_low = str[j].lower()
        if i_low != j_low:
            # print("compare", i, i_low,"with", j_low, j)
            return False
        i = i+1
        j = j-1
    return True
    

def valid_anagram(s: str, t: str) -> bool:
    count_s = {}
    count_t = {}
    for i in s:
        if i in count_s:
            count_s[i] +=1
        else: 
            count_s[i] = 1
    for i in t:
        if i in count_t:
            count_t[i] +=1
        else: 
            count_t[i] = 1

    return count_s == count_t

def keyboard_row(swords: list[str]) -> list[str]:
    result = list()
    keyRow = {"q": 1, "w": 1, "e": 1, "r": 1, "t": 1, "y": 1, "u": 1, "i": 1, "o": 1, "p": 1,
		"a": 2, "s": 2, "d": 2, "f": 2, "l": 2, "g": 2, "h": 2, "j": 2, "k": 2,
		"z": 3, "x": 3, "c": 3, "v": 3, "b": 3, "n": 3, "m": 3}
    for i in swords:
        oneRow = True
        char = i.lower()
        for j in range(1, len(char)):
            if keyRow[char[j]] != keyRow[char[j-1]]:
                oneRow = False
            
        if oneRow:
            result.append(i)
    
    return result

print(sqrt(4))

print(valid_palindrome("A man, a plan, a canal: Panama"))

print(valid_anagram("anagram", "nagaram"))

print(keyboard_row({"Hello", "Alaska", "Dad", "Peace"}))
