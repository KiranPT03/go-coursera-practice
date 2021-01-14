

def Solution(D):
    
    if D >= 0: 
        res = [int(x) for x in str(D)] 
        for i in range (len(res)):
            if res[i]<= 5:
                res = res[:i] + [5] + res[i:]
                print(res)
                break
            if i == len(res)-1:
                res = res + [5]
                print(res)            
                
   
Solution(678)