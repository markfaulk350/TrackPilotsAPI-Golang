resource "aws_lambda_function" "TrackPilotsAPI" {
  function_name = "TrackPilotsAPI"
  handler = "main"
  filename = "../../main.zip"
  source_code_hash = "${base64sha256(file("../../main.zip"))}"
  runtime = "go1.x"
  timeout = "30"
  role = "${aws_iam_role.lambda_exec.arn}"

  environment {
      variables = {
      DB_NAME = "${var.DB_NAME}"
      DB_PASSWORD = "${var.DB_PASSWORD}"
      DB_SESSION = "${var.DB_SESSION}"
      DB_USER = "${var.DB_USER}"
  }
 }
}

resource "aws_iam_role" "lambda_exec" {
  name = "serverless_example_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.example.arn}"
  principal     = "apigateway.amazonaws.com"
  source_arn = "${aws_api_gateway_deployment.example.execution_arn}/*/*"
}