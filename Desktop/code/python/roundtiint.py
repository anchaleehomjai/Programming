import math
def round_to_int(x):
    if (x%1>=0.5 and x>0 and x%1 != 0) or (x < 0 and abs(x)%1<0.5 and x%1 != 0):
        return math.ceil(x)
    elif x==0:
        return 0
    elif (x%1<0.5 and x>0 and x%1 != 0) or (x < 0 and abs(x)%1>=0.5 and x%1 != 0):
        return math.ceil(x)-1
    else:
        return int(x)
def main() :
    x = float (input(""))
    y = round_to_int(x)
    print(y)
if  __name__ == '__main__':
    main()