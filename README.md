# aws-go

A starter template for creating a serverless API in Go on AWS Lambda

---

## Required Technologies

* [Serverless Framework](https://www.serverless.com/framework/)
* [AWS](https://aws.amazon.com)
* [Go](https://go.dev/learn/)

---

## Getting Started

### 1. Create a new project using this template

```
serverless create --template-url https://github.com/nativecrane/aws-go --path <my-cool-go-api>
```

```
cd <my-cool-go-api>
```

### 2. Run unit tests
```
make test
```

### 3. Configure an AWS profile in .env.dev

### 4. Deploy starter project to AWS
```
make deploy-dev
```

### 5. Test that the API is working
```
curl <api-endpoint>
```

### 6. Start writing your API!
```
open . -a Visual\ Studio\ Code
```

---

## What am I doing wrong?

* Is your AWS profile configured in `~/.aws/credentials`?
* Does your configured AWS profile have appropriate permissions?

---

### Notes

It's probably a good idea to add your .env files to .gitignore
