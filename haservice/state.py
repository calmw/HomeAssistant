import requests

token = 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiI5NTQxY2M2NmUwNWQ0ZjJhOGU2MjZkZGU0Yzg0MjU3NyIsImlhdCI6MTcxMDYwNDk0MSwiZXhwIjoyMDI1OTY0OTQxfQ.jWAHyZttnndgOr4tnEfUEeM3FlJdJdalckJVFGWDBbg'

headers = {
    "Authorization": token,
    "content-type": "application/json",
}


def change_state(entity_id, state):
    url = "http://localhost:8123/api/states/light.study_light"
    response = requests.post(url, '{"state":"off"}', headers=headers)
    print(response.text)


def get_state(entity_id):
    url = "http://localhost:8123/api/states/" + entity_id
    return requests.get(url, headers=headers)
