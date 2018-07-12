bucket=ivan-sam-bucket-testing
output_template_file=serverless-output.yaml
stack_name=prod-sns-service

GOOS=linux go build -o main
sam package --template-file $(pwd)/template.yaml --s3-bucket $bucket --output-template-file $output_template_file
aws cloudformation deploy --template-file $(pwd)/$output_template_file \
           --stack-name $stack_name \
           --capabilities CAPABILITY_IAM \
           --parameter-overrides  SomeString=hello_prod
