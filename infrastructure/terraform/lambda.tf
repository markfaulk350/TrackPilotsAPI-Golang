resource "aws_lambda_function" "TrackPilotsAPI" {
  function_name = "TrackPilotsAPI"
  handler = "main"
  filename = "../../main.zip"
  source_code_hash = "${filebase64sha256("../../main.zip")}"
  runtime = "go1.x"
  timeout = "30"
  role = "${aws_iam_role.TrackPilotsAPI_lambda.arn}"

  environment {
      variables = {
      DB_NAME = "${var.DB_NAME}"
      DB_PASSWORD = "${var.DB_PASSWORD}"
      DB_SESSION = "${var.DB_SESSION}"
      DB_USER = "${var.DB_USER}"
  }
 }
}
