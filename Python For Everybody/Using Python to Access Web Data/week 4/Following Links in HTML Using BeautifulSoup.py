from urllib.request import urlopen
from bs4 import BeautifulSoup
import ssl

global url, count_2
# Ignore SSL certificate errors
ctx = ssl.create_default_context()
ctx.check_hostname = False
ctx.verify_mode = ssl.CERT_NONE

count_2 = 0
# retrieve data using parser
url = input('Enter URL - ')
count = int(input('Enter count: '))
position = int(input('Enter Position: '))
print("Retrieving..",url)
while(count>0):
    #print("before:..",url)
    html = urlopen(url, context=ctx).read()
    soup = BeautifulSoup(html, "html.parser")
    # Retrieve all of the anchor tags
    tags = soup('a')
    for tag in tags:
        count_2 = count_2+1
        if (count_2 == position):
            url = tag.get('href', None)
            print("Retrieving..",url)
            count_2 = 0
            break
    count = count - 1
