#!/bin/sh

# Export DB_DRIVER from GitHub secrets
export DB_DRIVER="postgres"

# Export DB_SOURCE from GitHub secrets
export DB_SOURCE=${{ secrets.DB_SOURCE }}

# Export NOTIFDB_SOURCE from GitHub secrets
export NOTIFDB_SOURCE=${{ secrets.NOTIFDB_SOURCE }}

# Export REDIS_ADDR from GitHub secrets
export REDIS_ADDR=${{ secrets.REDIS_ADDR }}

# Export HTTP_SERVER_ADDRESS from GitHub secrets
export HTTP_SERVER_ADDRESS=${{ secrets.HTTP_SERVER_ADDRESS }}

# Export GRPC_SERVER_ADDRESS from GitHub secrets
export GRPC_SERVER_ADDRESS=${{ secrets.GRPC_SERVER_ADDRESS }}

# Export TOKEN_SYMMETRIC_KEY from GitHub secrets
export TOKEN_SYMMETRIC_KEY=${{ secrets.TOKEN_SYMMETRIC_KEY }}

# Export ACCESS_TOKEN_DURATION from GitHub secrets
export ACCESS_TOKEN_DURATION=${{ secrets.ACCESS_TOKEN_DURATION }}

# Export REFRESH_TOKEN_DURATION from GitHub secrets
export REFRESH_TOKEN_DURATION=${{ secrets.REFRESH_TOKEN_DURATION }}

# Export REDIRECT_AUTH_TOKEN from GitHub secrets
export REDIRECT_AUTH_TOKEN=${{ secrets.REDIRECT_AUTH_TOKEN }}

# Export GOOGLE_APPLICATION_CREDENTIALS from GitHub secrets
export GOOGLE_APPLICATION_CREDENTIALS=${{ secrets.GOOGLE_APPLICATION_CREDENTIALS }}

# Export EMAIL_HOST from GitHub secrets
export EMAIL_HOST=${{ secrets.EMAIL_HOST }}

# Export EMAIL_PASSWORD from GitHub secrets
export EMAIL_PASSWORD=${{ secrets.EMAIL_PASSWORD }}

# Export EMAIL_SMTP from GitHub secrets
export EMAIL_SMTP=${{ secrets.EMAIL_SMTP }}

# Export BASE_URL from GitHub secrets
export BASE_URL=${{ secrets.BASE_URL }}

# Export GOOGLE_OAUTH_CLIENT_ID from GitHub secrets
export GOOGLE_OAUTH_CLIENT_ID=${{ secrets.GOOGLE_OAUTH_CLIENT_ID }}

# Export GOOGLE_OAUTH_SECRET from GitHub secrets
export GOOGLE_OAUTH_SECRET=${{ secrets.GOOGLE_OAUTH_SECRET }}

# Export USER_DIR from GitHub secrets
export USER_DIR=${{ secrets.USER_DIR }}

# Export GCP_BUCKET_NAME from GitHub secrets
export GCP_BUCKET_NAME=${{ secrets.GCP_BUCKET_NAME }}

# Export MAX_UPLOAD_SIZE from GitHub secrets
export MAX_UPLOAD_SIZE=${{ secrets.MAX_UPLOAD_SIZE }}
