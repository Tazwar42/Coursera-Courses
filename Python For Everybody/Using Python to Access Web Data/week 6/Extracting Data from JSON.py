import urllib.request, urllib.parse, urllib.error
import xml.etree.ElementTree as ET
import ssl
import json

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
    data = uh.read().decode()
    print('Retrieved', len(data), 'characters')
    #print(data)
    # loop through the dict and find the sum
    info = json.loads(data)
    for number in range(len(info['comments'])):
        sum = sum + int(info['comments'][number]['count'])
    print(sum)
