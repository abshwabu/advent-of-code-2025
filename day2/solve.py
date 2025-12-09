def is_invalid_id(num):
    s = str(num)
    if len(s) > 1 and s.startswith('0'):
        return False
    
    # An ID is invalid if it's made only of some sequence of digits repeated twice.
    # This means the string length must be even.
    if len(s) % 2 != 0:
        return False
    
    half_length = len(s) // 2
    return s[:half_length] == s[half_length:]

def solve():
    with open('input.txt', 'r') as f:
        line = f.read().strip()
    
    ranges_str = line.split(',')
    total_sum = 0
    
    for r_str in ranges_str:
        start_str, end_str = r_str.split('-')
        start = int(start_str)
        end = int(end_str)
        
        for i in range(start, end + 1):
            if is_invalid_id(i):
                total_sum += i
                
    print(total_sum)

if __name__ == "__main__":
    solve()