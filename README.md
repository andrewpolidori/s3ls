# s3ls

A simple static binary that outputs the contents of a given S3 Bucket and key prefix.

## Motivation

This is a sad tool with an expected audience of 1 user, me. I was unable to successfully execute [awscli](
https://aws.amazon.com/cli/) commands using the Golang [ssh.Session.CombinedOutput](
https://godoc.org/golang.org/x/crypto/ssh#Session.CombinedOutput) function. Output was hidden; mysterious error messages
were returned and I was generally confused. After a few hours of unsuccessfully debugging this, I spent 15 minutes to 
write this statically compilable tool that does exactly what I need with no dependencies!

## Usage

```
s3ls <region> s3://<bucket-name>/<prefix>/ 
```

This will return the names of all objects stored at `s3://<bucket-name>/<prefix>/`.

## Example

```
> s3ls us-west-2 s3://mongodb-with-backup-qqlpuh-backup/daily/

daily/mongodump.bak.2017-04-03-194012
daily/mongodump.bak.2017-04-03-194641
daily/mongodump.bak.2017-04-03-194846
daily/mongodump.bak.2017-04-03-195023
daily/mongodump.bak.2017-04-03-195149
daily/mongodump.bak.2017-04-03-195454
daily/mongodump.bak.2017-04-03-195656
...
```
