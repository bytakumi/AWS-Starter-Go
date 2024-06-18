# AWS-Starter-Go

## File Tree
```
.
├── app              // Lambda
│   └── hello        // Lambda functions
│       └── main.go
├── serverless       // Infra
│   ├── dynamodb     // DynamoDB
│   ├── function     // Lambda
│   ├── iam          // IAM/Role
│   ├── loggroup     // CloudWatchLogs
│   ├── param        // Stage Parameters
│   ├── stepfunction // StepFunction
│   └── vpc          // VPC
├── tools
│   └── go_build.sh  // Build tools
├── Makefile         // Build/Deploy commands
└── serverless.yml   // sls
```

## Usage
```bash

# Set up the profile in AWS Credentials (~/.aws/credentials) as {service-name}-{stage-name}.
# For example, if service=aws-starter-go and stage=dev, the profile should be aws-starter-go-dev.
#   -> The service name is specified in the service parameter in serverless.yml.
#   -> The stage name is specified during the deploy command.
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
