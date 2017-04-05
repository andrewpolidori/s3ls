# s3dl

A simple static binary for downloading the S3 file in a directory from AWS to the given local directory

## Motivation

This is a sad tool with an expected audience of 1 user, me. I was unable to successfully execute [awscli](
https://aws.amazon.com/cli/) commands using the Golang [ssh.Session.CombinedOutput](
https://godoc.org/golang.org/x/crypto/ssh#Session.CombinedOutput) function. Output was hidden; mysterious error messages
were returned and I was generally confused. After a few hours of unsuccessfully debugging this, I spent 15 minutes to 
write this statically compilable tool that does exactly what I need with no depenencies!

## Usage

```
s3dl <region> s3://<bucket-name>/<prefix>/ <local-path> 
```

This will download the *first* file from all files stored at `s3://<bucket-name>/<prefix>/` to the local directory `<local-path>`>

## Example

```
s3dl us-west-2 s3://my-bucket/folder1/ /tmp/file1
```
