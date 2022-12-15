import base64
import secrets
import hashlib
# csgrgr7878#$ is your-secret-key and fef787@# is salt
hmac_key = hashlib.pbkdf2_hmac('sha256', b'csgrgr7878#$', b'fef787@#', 100000)
hmac_key_ascii = base64.b64encode(hmac_key).decode('ascii')
print(hmac_key_ascii)

signing_key = secrets.token_urlsafe(32)
print(signing_key)

pepper = secrets.token_bytes()
pepper_ascii = base64.b64encode(pepper).decode('ascii')
print(pepper_ascii)