/*
   Copyright 2010-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

   This file is licensed under the Apache License, Version 2.0 (the "License").
   You may not use this file except in compliance with the License. A copy of
   the License is located at

    http://aws.amazon.com/apache2.0/

   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied. See the License for the
   specific language governing permissions and limitations under the License.
*/

package main

import (
    "fmt"
    "os"

    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

// Gets the ACL for a bucket
//
// Usage:
//    go run s3_get_bucket_acl.go BUCKET
func main() {
    if len(os.Args) != 2 {
        exitErrorf("Bucket name required\nUsage: go run", os.Args[0], "BUCKET")
    }

    bucket := os.Args[1]

    // Initialize a session in us-west-2 that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials.
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-2")},
    )

    // Create S3 service client
    svc := s3.New(sess)

    // Get bucket ACL
    result, err := svc.GetBucketAcl(&s3.GetBucketAclInput{Bucket: &bucket})

    if err != nil {
        exitErrorf(err.Error())
    }

    fmt.Println("Owner:", *result.Owner.DisplayName)

    fmt.Println("Grants")

    for _, g := range result.Grants {
        fmt.Println("  Grantee:   ", *g.Grantee.DisplayName)
        fmt.Println("  Type:      ", *g.Grantee.Type)
        fmt.Println("  Permission:", *g.Permission)
        fmt.Println("")
    }
}

func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}
