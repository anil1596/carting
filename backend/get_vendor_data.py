import requests
import json

resp = requests.get("http://localhost:7070/getvendors")
body = resp.text
print(body)

print("Item Menu")
id = {
    "vendorid":1
}

resp = requests.get("http://localhost:7070/getvendorsmenu",data=json.dumps(id))
body = resp.text
print(body)
