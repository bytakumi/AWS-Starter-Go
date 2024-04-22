# AWS-Starter-Go

## File Tree
```
.
├── app              // Lambdaアプリケーション
│   └── hello        // Lambda関数
│       └── main.go
├── serverless       // IaC　インフラ定義
│   ├── dynamodb     // DynamoDB関連
│   ├── function     // Lambda関連
│   ├── iam          // IAM/Role関連
│   ├── loggroup     // CloudWatchLogs関連
│   ├── param        // Stageごとのパラメーター関連
│   ├── stepfunction // StepFunction関連
│   └── vpc          // VPC関連
├── tools
│   └── go_build.sh  // Goビルドツール
├── Makefile         // Goビルドやデプロイコマンド
└── serverless.yml   // sls定義
```

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
