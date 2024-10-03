# Encryption
All about encryption

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
