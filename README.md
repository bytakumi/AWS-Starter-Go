# AWS-Starter-Go

## Usage
```bash

# AWS Credentials(~/.aws/credentials)にて、{service名}-{stage名} のProfileを設定します。
# serivce=aws-starter-go, stage=dev の場合は、 aws-starter-go-dev のProfileが必要になります。
#   -> service名は、serverless.ymlのserviceパラメーターに記載されています。
#   -> stage名は、deploy時のコマンドで指定されます。
vim ~/.aws/credentials

# Deploy
make deploy_all stage=dev
```

## Resources

- IAM Role
- Lambda
- DynamoDB
- StepFunction
- APIGateway
- VPC
  - Private Subnet
    - NAT Gateway
  - Public Subnet
    - Internet Gateway
  - Elastic IP
