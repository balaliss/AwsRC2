# Cloud Resume Challenge

This repository contains the source code and infrastructure configuration for my Cloud Resume Challenge. The challenge involves deploying a resume website to AWS and implementing a variety of AWS services and best practices.

## Static Website

The resume website is hosted as a static site on Amazon S3 and is accessible via the following link: [sbawsresumechallenge.com](http://sbawsresumechallenge.com)

## Overview

The Cloud Resume Challenge is a comprehensive project that demonstrates cloud computing skills and knowledge of AWS services. The project includes the following components:

1. **HTML/CSS**: The resume is built using HTML and styled with CSS. It is a static website hosted on Amazon S3.

2. **JavaScript**: The website includes a visitor counter implemented in JavaScript. The counter tracks unique visits to the site.

3. **Database**: Visitor data is stored in a DynamoDB table.

4. **API**: A REST API is created using AWS API Gateway and Lambda to handle requests from the website to the DynamoDB table. The API is secured using AWS IAM roles and policies.

5. **Python**: The Lambda function that interacts with DynamoDB is written in Python, utilizing the `boto3` library.

6. **CI/CD**: The project uses GitHub Actions for Continuous Integration and Continuous Deployment (CI/CD). When updates are pushed to the repository, the changes are automatically deployed to the S3 bucket and the Lambda function.

7. **Infrastructure as Code**: AWS resources are defined using the AWS Serverless Application Model (SAM) and deployed using the AWS SAM CLI.

8. **HTTPS & DNS**: The website is served securely over HTTPS using Amazon CloudFront, with a custom domain configured via Route 53.

## Project Structure

```bash
.
├── front-end/               # Contains the HTML, CSS, and JavaScript for the resume website
├── back-end/                # Contains the Lambda function code and SAM template
├── .github/workflows/       # CI/CD pipelines for the front-end and back-end
├── README.md                # This README file
└── samconfig.toml           # Configuration for the SAM CLI
