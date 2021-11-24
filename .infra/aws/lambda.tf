resource "aws_lambda_function" "lambda" {
  filename         = "archive.zip"
  function_name    = var.lambda_function_name
  description      = ""
  runtime          = "go1.x"
  architectures    = ["x86_64"]
  handler          = "main"
  layers           = []
  tags             = {}
  tags_all         = {}
  source_code_hash = filebase64sha256("archive.zip")
  role             = aws_iam_role.lambda_execution_role.arn

  depends_on = [
    aws_iam_role_policy_attachment.lambda_logs,
    aws_cloudwatch_log_group.lambda_log_group,
  ]
}