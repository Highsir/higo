import requests

filename = "640.gif"
path = "test"

url = "https://mirrors.tencent.com/repository/generic/v_gjbgao/%s/%s" % (path, filename)

username = "v_gjbgao"
token = "b81e339e069d11eeb44c5254006c56f5"

with open(r"C:\Users\v_gjbgao\Desktop\6401.gif", "rb") as fp:
  response = requests.request("PUT", url, auth=(username, token), data=fp)

print(response.text)