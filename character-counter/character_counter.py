import string
import re

input_data = input()
print(type(input_data), input_data)

print(type(string.ascii_letters), string.ascii_letters)
# <class 'str'> abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ

def str_count():
    count_en = count_dg = count_sp = count_zh = count_other = 0

    for char in input_data:
        # print(type(char), char)
        if char in string.ascii_letters:
            count_en+=1
        elif char in string.digits:
            count_dg+=1
        elif char.isspace():
            count_sp+=1
        # elif char.isalpha():
        else:
            m = re.findall(u"[\u4e00-\u9fa5]+", char)
            
            if len(m) == 1:
                count_zh+=1
            else:
                count_other+=1
        
    result = f'en:{count_en}, dg:{count_dg}, sp:{count_sp}, zh:{count_zh}, other:{count_other}'
    print(result)

    total_chars = count_en + count_dg + count_sp + count_zh + count_other
    if len(input_data) != total_chars:
        raise('error')

str_count()

# Python 3
# str 代表的是 Unicode

# 用 encode 方法指定編碼，傳回一個 bytes 實例
# type_bytes = type_str.encode('utf-8')

# 用 decode 指定編碼，傳回代表 Unicode 的 str 實例
# str = b'abc'.decode('utf-8')

# test = '測試'
# test2 = u'測試'
# test3 = b'abc'
# print(type(test), len(test), test)
# print(type(test2), len(test2), test2)
# print(type(test3), len(test3), test3)

