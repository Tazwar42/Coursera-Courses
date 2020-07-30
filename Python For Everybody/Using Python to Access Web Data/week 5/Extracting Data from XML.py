import urllib.request, urllib.parse, urllib.error
import xml.etree.ElementTree as ET
import ssl

# Ignore SSL certificate errors
ctx = ssl.create_default_context()
ctx.check_hostname = False
ctx.verify_mode = ssl.CERT_NONE

while True:
    sum = 0
    #taking input
    url = input('Enter location: ')
    if len(url) < 1: break
    # retrieving url and decoding it
    print('Retrieving', url)
    uh = urllib.request.urlopen(url, context=ctx)
    data = uh.read()
    print('Retrieved', len(data), 'characters')
    print(data.decode())
    # make a tree of xml and convert it to a list
    tree = ET.fromstring(data)
    counts = tree.findall('comments/comment')
    # retrieve data from the list
    for item in counts:
        sum = sum + int(item.find('count').text)
    print(sum)
