text = "X-DSPAM-Confidence:    0.8475";

index_num = text.find('0')
num = float(text[index_num:])
print(num)
