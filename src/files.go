package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/aws/resource/elasticloadbalancing/aws_lb_target_group.json
var awsLbTargetGroup []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_lb_target_group_attachment.json
var awsLbTargetGroupAttachment []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_lb_listener.json
var awsLbListener []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_lb.json
var awsLb []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket.json
var awsS3Bucket []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_acl.json
var awsS3BucketACL []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_versioning.json
var awsS3BucketVersioning []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_server_side_encryption_configuration.json
var awsS3BucketServerSideEncryptionConfiguration []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_public_access_block.json
var awsS3BucketPublicAccessBlock []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_logging.json
var awsS3BucketLogging []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_lifecycle_configuration.json
var awsS3BucketLifecycleConfiguration []byte

//go:embed mapping/aws/resource/s3/aws_s3_bucket_policy.json
var awsS3BucketPolicy []byte

//go:embed mapping/aws/resource/s3/aws_s3_object.json
var awsS3Object []byte

//go:embed mapping/aws/resource/ec2/aws_instance.json
var awsInstance []byte

//go:embed mapping/aws/resource/ec2/aws_security_group.json
var awsSecurityGroup []byte

//go:embed mapping/aws/resource/ec2/aws_security_group_rule.json
var awsSecurityGroupRule []byte

//go:embed mapping/aws/resource/lambda/aws_lambda_function.json
var awsLambdaFunction []byte

//go:embed mapping/aws/resource/lambda/aws_lambda_alias.json
var awsLambdaAlias []byte

//go:embed mapping/aws/resource/lambda/aws_lambda_permission.json
var awsLambdaPermission []byte

//go:embed mapping/aws/resource/ec2/aws_vpc.json
var awsVpc []byte

//go:embed mapping/aws/resource/ec2/aws_subnet.json
var awsSubnet []byte

//go:embed mapping/aws/resource/ec2/aws_network_acl.json
var awsNetworkACL []byte

//go:embed mapping/aws/resource/kms/aws_kms_key.json
var awsKmsKey []byte

//go:embed mapping/aws/resource/iam/aws_iam_role.json
var awsIamRole []byte

//go:embed mapping/aws/resource/iam/aws_iam_role_policy.json
var awsIamRolePolicy []byte

//go:embed mapping/aws/resource/iam/aws_iam_role_policy_attachment.json
var awsIamRolePolicyAttachment []byte

//go:embed mapping/aws/resource/iam/aws_iam_policy.json
var awsIamPolicy []byte

//go:embed mapping/aws/resource/iam/aws_iam_instance_profile.json
var awsIamInstanceProfile []byte

//go:embed mapping/aws/resource/iam/aws_iam_access_key.json
var awsIamAccessKey []byte

//go:embed mapping/aws/resource/iam/aws_iam_group.json
var awsIamGroup []byte

//go:embed mapping/aws/resource/iam/aws_iam_group_membership.json
var awsIamGroupMembership []byte

//go:embed mapping/aws/resource/iam/aws_iam_group_policy.json
var awsIamGroupPolicy []byte

//go:embed mapping/aws/resource/iam/aws_iam_group_policy_attachment.json
var awsIamGroupPolicyAttachment []byte

//go:embed mapping/aws/resource/iam/aws_iam_policy_attachment.json
var awsIamPolicyAttachment []byte

//go:embed mapping/aws/resource/iam/aws_iam_service_linked_role.json
var awsIamServiceLinkedRole []byte

//go:embed mapping/aws/resource/iam/aws_iam_user.json
var awsIamUser []byte

//go:embed mapping/aws/resource/iam/aws_iam_user_login_profile.json
var awsIamUserLoginProfile []byte

//go:embed mapping/aws/resource/iam/aws_iam_user_policy.json
var awsIamUserPolicy []byte

//go:embed mapping/aws/resource/iam/aws_iam_user_policy_attachment.json
var awsIamUserPolicyAttachment []byte

//go:embed mapping/aws/resource/mq/aws_mq_broker.json
var awsMqBroker []byte

//go:embed mapping/aws/resource/mq/aws_mq_configuration.json
var awsMqConfiguration []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_group.json
var awsCloudwatchLogGroup []byte

//go:embed mapping/aws/resource/events/aws_cloudwatch_event_rule.json
var awsCloudwatchEventRule []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_metric_filter.json
var awsCloudwatchLogMetricFilter []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_resource_policy.json
var awsCloudwatchLogResourcePolicy []byte

//go:embed mapping/aws/resource/logs/aws_cloudwatch_log_subscription_filter.json
var awsCloudwatchLogSubscriptionFilter []byte

//go:embed mapping/aws/resource/cloudwatch/aws_cloudwatch_metric_alarm.json
var awsCloudwatchMetricAlarm []byte

//go:embed mapping/aws/resource/events/aws_cloudwatch_event_target.json
var awsCloudwatchEventTarget []byte

//go:embed mapping/aws/resource/route53/aws_route53_record.json
var awsRoute53Record []byte

//go:embed mapping/aws/resource/route53/aws_route53_zone.json
var awsRoute53Zone []byte

//go:embed mapping/aws/resource/sns/aws_sns_topic.json
var awsSnsTopic []byte

//go:embed mapping/aws/resource/sns/aws_sns_topic_subscription.json
var awsSnsTopicSubscription []byte

//go:embed mapping/aws/resource/sns/aws_sns_topic_policy.json
var awsSnsTopicPolicy []byte

//go:embed mapping/aws/resource/ec2/aws_key_pair.json
var awsKeyPair []byte

//go:embed mapping/aws/resource/rds/aws_db_instance.json
var awsDbInstance []byte

//go:embed mapping/aws/resource/dynamodb/aws_dynamodb_table.json
var awsDynamodbTable []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_parameter.json
var awsSsmParameter []byte

//go:embed mapping/aws/resource/ec2/aws_route.json
var awsRoute []byte

//go:embed mapping/aws/resource/ec2/aws_default_security_group.json
var awsDefaultSecurityGroup []byte

//go:embed mapping/aws/resource/rds/aws_db_subnet_group.json
var awsDbSubnetGroup []byte

//go:embed mapping/aws/resource/wafv2/aws_wafv2_web_acl.json
var awsWafv2WebACL []byte

//go:embed mapping/aws/resource/wafv2/aws_wafv2_ip_set.json
var awsWafv2IpSet []byte

//go:embed mapping/aws/resource/wafv2/aws_wafv2_rule_group.json
var awsWafv2RuleGroup []byte

//go:embed mapping/aws/resource/wafv2/aws_wafv2_regex_pattern_set.json
var awsWafv2RegexPatternSet []byte

//go:embed mapping/aws/resource/ec2/aws_vpn_gateway.json
var awsVpnGateway []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_endpoint.json
var awsVpcEndpoint []byte

//go:embed mapping/aws/resource/apigateway/aws_api_gateway_rest_api.json
var awsAPIGatewayRestAPI []byte

//go:embed mapping/aws/resource/apigateway/aws_apigatewayv2_api.json
var awsApigatewayv2Api []byte

//go:embed mapping/aws/resource/apigateway/aws_api_gateway_account.json
var awsAPIGatewayAccount []byte

//go:embed mapping/aws/resource/sqs/aws_sqs_queue.json
var awsSqsQueue []byte

//go:embed mapping/aws/resource/sqs/aws_sqs_queue_policy.json
var awsSqsQueuePolicy []byte

//go:embed mapping/aws/resource/ec2/aws_ebs_volume.json
var awsEbsVolume []byte

//go:embed mapping/aws/resource/autoscaling/aws_autoscaling_group.json
var awsAutoscalingGroup []byte

//go:embed mapping/aws/resource/autoscaling/aws_autoscaling_attachment.json
var awsAutoscalingAttachment []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_elb.json
var awsElb []byte

//go:embed mapping/aws/resource/ec2/aws_internet_gateway.json
var awsInternetGateway []byte

//go:embed mapping/aws/resource/autoscaling/aws_launch_configuration.json
var awsLaunchConfiguration []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_capacity_reservation.json
var awsEc2CapacityReservation []byte

//go:embed mapping/aws/resource/ec2/aws_network_interface.json
var awsNetworkInterface []byte

//go:embed mapping/aws/resource/ec2/aws_placement_group.json
var awsPlacementGroup []byte

//go:embed mapping/aws/resource/ec2/aws_spot_instance_request.json
var awsSpotInstanceRequest []byte

//go:embed mapping/aws/resource/ec2/aws_volume_attachment.json
var awsVolumeAttachment []byte

//go:embed mapping/aws/resource/ec2/aws_eip.json
var awsEip []byte

//go:embed mapping/aws/resource/firehose/aws_kinesis_firehose_delivery_stream.json
var awsKinesisFirehoseDeliveryStream []byte

//go:embed mapping/aws/resource/firehose/aws_kinesis_stream.json
var awsKinesisStream []byte

//go:embed mapping/aws/resource/firehose/aws_kinesis_video_stream.json
var awsKinesisVideoStream []byte

//go:embed mapping/aws/resource/budgets/aws_budgets_budget.json
var awsBudgetsBudget []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_elastic_beanstalk_application.json
var awsElasticBeanstalkApplication []byte

//go:embed mapping/aws/resource/ec2/aws_flow_log.json
var awsFlowLog []byte

//go:embed mapping/aws/resource/ecr/aws_ecr_lifecycle_policy.json
var awsEcrLifecyclePolicy []byte

//go:embed mapping/aws/resource/ecr/aws_ecr_pull_through_cache_rule.json
var awsEcrPullThroughCacheRule []byte

//go:embed mapping/aws/resource/ecr/aws_ecr_repository.json
var awsEcrRepository []byte

//go:embed mapping/aws/resource/kms/aws_kms_alias.json
var awsKmsAlias []byte

//go:embed mapping/aws/resource/ec2/aws_route_table.json
var awsRouteTable []byte

//go:embed mapping/aws/resource/ec2/aws_route_table_association.json
var awsRouteTableAssociation []byte

//go:embed mapping/aws/resource/ec2/aws_nat_gateway.json
var awsNatGateway []byte

//go:embed mapping/aws/resource/rds/aws_db_option_group.json
var awsDbOptionGroup []byte

//go:embed mapping/aws/resource/rds/aws_db_parameter_group.json
var awsDbParameterGroup []byte

//go:embed mapping/aws/resource/secretsmanager/aws_secretsmanager_secret.json
var awsSecretsmanagerSecret []byte

//go:embed mapping/aws/resource/secretsmanager/aws_secretsmanager_secret_version.json
var awsSecretsmanagerSecretVersion []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_document.json
var awsSsmDocument []byte

//go:embed mapping/aws/resource/glue/aws_glue_classifier.json
var awsGlueClassifier []byte

//go:embed mapping/aws/resource/glue/aws_glue_crawler.json
var awsGlueCrawler []byte

//go:embed mapping/aws/resource/glue/aws_glue_catalog_database.json
var awsGlueCatalogDatabase []byte

//go:embed mapping/aws/resource/glue/aws_glue_catalog_table.json
var awsGlueCatalogTable []byte

//go:embed mapping/aws/resource/glue/aws_glue_connection.json
var awsGlueConnection []byte

//go:embed mapping/aws/resource/glue/aws_glue_data_catalog_encryption_settings.json
var awsGlueDataCatalogEncryptionSettings []byte

//go:embed mapping/aws/resource/glue/aws_glue_ml_transform.json
var awsGlueMlTransform []byte

//go:embed mapping/aws/resource/glue/aws_glue_trigger.json
var awsGlueTrigger []byte

//go:embed mapping/aws/resource/glue/aws_glue_job.json
var awsGlueJob []byte

//go:embed mapping/aws/resource/glue/aws_glue_registry.json
var awsGlueRegistry []byte

//go:embed mapping/aws/resource/glue/aws_glue_resource_policy.json
var awsGlueResourcePolicy []byte

//go:embed mapping/aws/resource/glue/aws_glue_schema.json
var awsGlueSchema []byte

//go:embed mapping/aws/resource/glue/aws_glue_security_configuration.json
var awsGlueSecurityConfiguration []byte

//go:embed mapping/aws/resource/glue/aws_glue_user_defined_function.json
var awsGlueUserDefinedFunction []byte

//go:embed mapping/aws/resource/glue/aws_glue_workflow.json
var awsGlueWorkflow []byte

//go:embed mapping/aws/resource/codebuild/aws_codebuild_project.json
var awsCodebuildProject []byte

//go:embed mapping/aws/resource/codecommit/aws_codecommit_repository.json
var awsCodecommitRepository []byte

//go:embed mapping/aws/resource/codepipeline/aws_codepipeline.json
var awsCodepipeline []byte

//go:embed mapping/aws/resource/codeartifact/aws_codeartifact_repository.json
var awsCodeartifactRepository []byte

//go:embed mapping/aws/resource/codeartifact/aws_codeartifact_domain.json
var awsCodeartifactDomain []byte

//go:embed mapping/aws/resource/codeartifact/aws_codeartifact_repository_permissions_policy.json
var awsCodeartifactRepositoryPermissionsPolicy []byte

//go:embed mapping/aws/resource/codeartifact/aws_codeartifact_domain_permissions_policy.json
var awsCodeartifactDomainPermissionsPolicy []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_patch_baseline.json
var awsSsmPatchBaseline []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_patch_group.json
var awsSsmPatchGroup []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_maintenance_window.json
var awsSsmMaintenanceWindow []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_maintenance_window_target.json
var awsSsmMaintenanceWindowTarget []byte

//go:embed mapping/aws/resource/ssm/aws_ssm_maintenance_window_task.json
var awsSsmMaintenanceWindowTask []byte

//go:embed mapping/aws/resource/ec2/aws_launch_template.json
var awsLaunchTemplate []byte

//go:embed mapping/aws/resource/directoryservice/aws_directory_service_directory.json
var awsDirectoryServiceDirectory []byte

//go:embed mapping/aws/resource/directoryservice/aws_directory_service_log_subscription.json
var awsDirectoryServiceLogSubscription []byte

//go:embed mapping/aws/resource/cloudtrail/aws_cloudtrail.json
var awsCloudtrail []byte

//go:embed mapping/aws/resource/rds/aws_rds_cluster_parameter_group.json
var awsRdsClusterParameterGroup []byte

//go:embed mapping/aws/resource/ec2/aws_network_acl_rule.json
var awsNetworkACLRule []byte

//go:embed mapping/aws/resource/acm/aws_acm_certificate.json
var awsAcmCertificate []byte

//go:embed mapping/aws/resource/acm-pa/aws_acmpca_certificate_authority.json
var awsAcmpcaCertificateAuthority []byte

//go:embed mapping/aws/resource/ecs/aws_ecs_cluster.json
var awsEcsCluster []byte

//go:embed mapping/aws/resource/ecs/aws_ecs_task_definition.json
var awsEcsTaskDefinition []byte

//go:embed mapping/aws/resource/ecs/aws_ecs_service.json
var awsEcsService []byte

//go:embed mapping/aws/resource/application-autoscaling/aws_appautoscaling_scheduled_action.json
var awsAppautoscalingScheduledAction []byte

//go:embed mapping/aws/resource/application-autoscaling/aws_appautoscaling_target.json
var awsAppautoscalingTarget []byte

//go:embed mapping/aws/resource/application-autoscaling/aws_appautoscaling_policy.json
var awsAppautoscalingPolicy []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_identity_provider.json
var awsCognitoIdentityProvider []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_resource_server.json
var awsCognitoResourceServer []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_identity_provider.json
var awsCognitoRiskConfiguration []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_user.json
var awsCognitoUser []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_user_group.json
var awsCognitoUserGroup []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_user_in_group.json
var awsCognitoUserInGroup []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_user_pool.json
var awsCognitoUserPool []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_user_pool_client.json
var awsCognitoUserPoolClient []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_user_pool_domain.json
var awsCognitoUserPoolDomain []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_user_pool_ui_customization.json
var awsCognitoUserPoolUICustomization []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_subnet_group.json
var awsRedshiftSubnetGroup []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_cluster.json
var awsRedshiftCluster []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_scheduled_action.json
var awsRedshiftScheduledAction []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_authentication_profile.json
var awsRedshiftAuthenticationProfile []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_cluster_iam_roles.json
var awsRedshiftClusterIamRoles []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_event_subscription.json
var awsRedshiftEventSubscription []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_hsm_client_certificate.json
var awsRedshiftHsmClientCertififcate []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_hsm_configuration.json
var awsRedshiftHsmConfiguration []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_usage_limit.json
var awsRedshiftUsageLimit []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_parameter_group.json
var awsRedshiftParameterGroup []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_snapshot_copy_grant.json
var awsRedshiftSnapshotCopyGrant []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_snapshot_schedule.json
var awsRedshiftSnapshotSchedule []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_snapshot_schedule_association.json
var awsRedshiftSnapshotScheduleAssociation []byte

//go:embed mapping/aws/resource/inspector/aws_inspector_assessment_target.json
var awsInspectorAssessmentTarget []byte

//go:embed mapping/aws/resource/inspector/aws_inspector_assessment_template.json
var awsInspectorAssessmentTemplate []byte

//go:embed mapping/aws/resource/inspector/aws_inspector_resource_group.json
var awsInspectorResouceGroup []byte
