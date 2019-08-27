resource "aws_api_gateway_rest_api" "TrackPilotsAPI" {
  name        = "TrackPilotsAPI-gateway"
  description = "This is my API description"
}

resource "aws_api_gateway_deployment" "deployment_v1" {
  rest_api_id = "${aws_api_gateway_rest_api.TrackPilotsAPI.id}"
  stage_name  = "api"

  depends_on = [
    "aws_api_gateway_integration.TrackPilotsAPI",
  ]
}

resource "aws_api_gateway_resource" "TrackPilotsAPIProxy" {
  parent_id   = "${aws_api_gateway_rest_api.TrackPilotsAPI.root_resource_id}"
  path_part   = "{proxy+}"
  rest_api_id = "${aws_api_gateway_rest_api.TrackPilotsAPI.id}"
}

resource "aws_api_gateway_method" "TrackPilotsAPI" {
  http_method      = "ANY"
  resource_id      = "${aws_api_gateway_resource.TrackPilotsAPIProxy.id}"
  rest_api_id      = "${aws_api_gateway_rest_api.TrackPilotsAPI.id}"
  authorization    = "NONE"
  api_key_required = false

  request_parameters = {
    "method.request.path.proxy" = true
  }
}

resource "aws_api_gateway_integration" "TrackPilotsAPI" {
  http_method             = "${aws_api_gateway_method.TrackPilotsAPI.http_method}"
  resource_id             = "${aws_api_gateway_resource.TrackPilotsAPIProxy.id}"
  rest_api_id             = "${aws_api_gateway_rest_api.TrackPilotsAPI.id}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "${aws_lambda_function.TrackPilotsAPI.invoke_arn}"

  depends_on = [
    "aws_api_gateway_method.TrackPilotsAPI",
  ]
}
