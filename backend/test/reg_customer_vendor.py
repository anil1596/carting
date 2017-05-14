import requests
import json
customer_to_register = {
     "customer_name" : "aman",
     "emailid" : "dt@gmail.com",
     "mobile" : ["1234567810"],
     "address" : "Hydrabaad",
     "password" : "aman123"
}

resp = requests.get("http://localhost:7070/registercustomer", data=json.dumps(customer_to_register))
body = resp.text
print(body)

vendor_to_register = {
     "owner" : "Rahul Sharma",
     "vendorname" : "Desi Tadka",
     "email" : "dt@gmail.com",
     "mobile" : ["1234567890"],
     "address" : "Hamirpur (H.P.)",
     "imageaddress" : "not available",
     "password" : "desitadka123"
}

resp = requests.post("http://localhost:7070/registervendor", data=json.dumps(vendor_to_register))
body = resp.text
print(body)
