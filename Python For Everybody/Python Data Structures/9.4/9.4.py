name = input("Enter file:")
if len(name) < 1 : name = "mbox-short.txt"
handle = open(name)

counts_dict = dict()
ls = list()

# refine the mails to a list
for line in handle:
    if line.startswith('From: '):
        words = line.split()
        #print(words[1])
        ls.append(words[1])
#print(counts_dict)


# retrieve/create/update counter in a dictionary
for word in ls:
    counts_dict[word] = counts_dict.get(word,0) + 1
#print(ls)

# find the largest number and word
bigcount = None
bigword = None
#print(counts)
for word,count in counts_dict.items():
    if bigcount is None or count>bigcount:
        bigword = word #capture the largest number
        bigcount = count
#print(bigword)
print(bigword,bigcount)
