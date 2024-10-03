# Encryption
All about encryption

# Asymmetric Key Pair

How It Works:

Public Key (encryption): You can give this key to anyone. They can use it to encrypt a message, but they cannot decrypt it.

Private Key (decryption): This key is kept secret by you. You use it to decrypt messages that were encrypted with your public key.

Example Use Case:
Alice wants to send a secure message to Bob. Bob gives Alice his public key. Alice encrypts her message with Bob's public key, and only Bob can decrypt it with his private key.

Another Usage is for message SIGNATURES

Alice wants to prove that she sent this message. Alice uses her Private Key to Encrypt the message. Alice gives her Public Key to Bob. If Bob can decrypt with Alice's Public key:

* Bob knows Alice must have sent the message
* Bob knows message was not modified in transit

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
