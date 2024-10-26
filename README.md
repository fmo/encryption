# General TLS info

Laith Academy general 40 mins video

https://www.youtube.com/watch?v=EnY6fSng3Ew

# Self Signed Certificates

Jay LaCroix Video

https://www.youtube.com/watch?v=Qg5ghpiEHm0

```
openssl req -new -newkey rsa:4096 -x509 -days 365 -nodes -out MyCertificate.crt -keyout MyKey.key
```

# Asymmetric Key Pair

How It Works:

Public Key (encryption): You can give this key to anyone. They can use it to encrypt a message, but they cannot decrypt it.

Private Key (decryption): This key is kept secret by you. You use it to decrypt messages that were encrypted with your public key.

Example Use Case:
Alice wants to send a secure message to Bob. Bob gives Alice his public key. Alice encrypts her message with Bob's public key, and only Bob can decrypt it with his private key.

# Signatures

Video: https://www.youtube.com/watch?v=_zyKvPvh808

Alice wants to prove that she sent this message. Alice uses her Private Key to Encrypt the message. Alice gives her Public Key to Bob. If Bob can decrypt with Alice's Public key:

* Bob knows Alice must have sent the message (AUTHENTICATION)
* Bob knows message was not modified in transit (INTEGRITY)

## Shortcoming of Asymmetric Key pairs

Can't use for Bulk Data, but can use it for Limited Data.

Bulk Data should be protected with Symmetric Encryption. But thats not super secure. 
What if we used Asymmetric Keys to share Symmetric Keys? (SSL and TLS based on this)

* Pam randomly generates a Symmetric Secret key
* Pam encrypts Symmetric Key with Jim's Public key
* Jim decrypts Symmetric Key with Jim's Private Key
* Bulk Data can now be Symmetrically Encrypted
* Both parties has identical Symmetric key, so they can share the encrypted data.

This concept of using both Asymmetric and Symmetric Encryption means Hybrid Encryption.

* Asymmetric Encryption to facilitate a Key Exchange
* Secret Key used with Symmetric Encryption for Bulk Data

Asymmetric Encryption

* Weakness: Slower - Reqquires much larger key sizes
* Weakness: Cipher text expansion
* Strength: More Secure - Private Key is never shared.

Symmetric Encryption

* Strength: Faster - Lower CPU Cost
* Strength: Cipher text is same size as Plain Text
* Weaknes: Less Secure - Secret key must be shared

# AES-256 
AES-256 encryption is a method that scrambles data using a 256-bit key

Encrypting with AES-256 and -iter:

```
echo "hello world" | openssl enc -aes-256-cbc -base64 -salt -iter 100000 -pass pass:yourpassword
```

Decrypting with AES-256 and -iter:

```
echo "ENCRYPTED_MESSAGE" | openssl enc -aes-256-cbc -base64 -d -salt -iter 100000 -pass pass:yourpassword
```

<p>Purpose of -salt:
Prevents Precomputed Attacks: When encrypting, the same input data (e.g., "hello world") with the same password would always produce the same encrypted output without a salt. Attackers could precompute encrypted versions of common passwords or messages (a technique called "rainbow tables"). Adding a random salt ensures that even if the same data and password are used, the resulting encrypted message will be different each time.

Adds Security: The salt is combined with the password during key derivation to make it harder for attackers to use brute-force or dictionary attacks to guess the password. It increases the uniqueness of the encryption process.

How it works:
When you use -salt, OpenSSL generates a random value (salt) and incorporates it into the key derivation. This makes the resulting encryption unique even with the same password.

Without a salt, if you encrypt the same data with the same password multiple times, you would get identical encrypted output, which is less secure.
</p>

# Certificate generation

## Create private key and self-signed certificate

```
openssl req -x509 \
-sha256 \
-newkey rsa:4096 \
-days 365 \
-keyout ca-key.pem \
-out ca-cert.pem \
-subj "/C=TR/ST=EURASIA/L=ISTANBUL/O=Software/OU=Microservices/CN=*.microservices.dev/emailAddress=test@test.com" \
-nodes
```

Validate

```
openssl x509 -in ca-cert.pem -noout -text
```

Certificate Signing Request

```
openssl req \
-newkey rsa:4096 \
-keyout server-key.pem \
-out server-req.pem \
-subj "/C=TR/ST=EURASIA/L=ISTANBUL/O=Microservices/OU=PaymentService/CN=*.microservices.dev/emailAddress=test@test.com" \
-nodes \
-sha256
```

Use CA's private key to sign the request

```
openssl x509 \
-req -in server-req.pem \
-days 60 \
-CA ca-cert.pem \
-CAkey ca-key.pem \
-CAcreateserial \
-out server-cert.pem \
-extfile server-ext.cnf \
-sha256
```

server-ext.cnf

```
subjectAltName=DNS:*.microservices.dev,DNS:*.microsercices.dev,IP:0.0.0.0
```

Validate

```
openssl x509 -in server-cert.pem -noout -text
```
