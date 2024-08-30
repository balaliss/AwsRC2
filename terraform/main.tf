provider "aws" {
  region = "us-east-1"
}

# S3 bucket configuration
resource "aws_s3_bucket" "resume_bucket" {
  bucket = "steve-balals-resume"
  
  website {
    index_document = "index.html"
    error_document = "error.html"
  }
}

# Public ACL for the S3 bucket
resource "aws_s3_bucket_acl" "resume_bucket_acl" {
  bucket = aws_s3_bucket.resume_bucket.bucket
  acl    = "public-read"
}


# Bucket policy to allow public read access
resource "aws_s3_bucket_policy" "resume_bucket_policy" {
  bucket = aws_s3_bucket.resume_bucket.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid       = "PublicReadGetObject",
        Effect    = "Allow",
        Principal = "*",
        Action    = "s3:GetObject",
        Resource  = "${aws_s3_bucket.resume_bucket.arn}/*"
      }
    ]
  })
}
