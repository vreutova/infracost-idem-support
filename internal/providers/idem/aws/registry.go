package aws

import "github.com/infracost/infracost/internal/schema"

var ResourceRegistry []*schema.RegistryItem = []*schema.RegistryItem{
	// GetAPIGatewayRestAPIRegistryItem(),
	// GetAPIGatewayStageRegistryItem(),
	// GetAPIGatewayv2ApiRegistryItem(),
	GetAutoscalingGroupRegistryItem(),
	GetACMCertificate(),
	// GetACMPCACertificateAuthorityRegistryItem(),
	// GetCloudfrontDistributionRegistryItem(),
	// GetCloudwatchDashboardRegistryItem(),
	// GetCloudwatchEventBusItem(),
	GetCloudwatchLogGroupItem(),
	// GetCloudwatchMetricAlarmRegistryItem(),
	// GetCodebuildProjectRegistryItem(),
	// GetConfigRuleItem(),
	// GetConfigurationRecorderItem(),
	// GetConfigOrganizationCustomRuleItem(),
	// GetConfigOrganizationManagedRuleItem(),
	// getDataTransferRegistryItem(),
	// GetDBInstanceRegistryItem(),
	// GetDMSRegistryItem(),
	// GetDocDBClusterInstanceRegistryItem(),
	// GetDocDBClusterRegistryItem(),
	// GetDocDBClusterSnapshotRegistryItem(),
	// GetDXConnectionRegistryItem(),
	// GetDXGatewayAssociationRegistryItem(),
	// GetDynamoDBTableRegistryItem(),
	// GetEBSSnapshotCopyRegistryItem(),
	// GetEBSSnapshotRegistryItem(),
	// GetEBSVolumeRegistryItem(),
	// GetEC2ClientVPNEndpointRegistryItem(),
	// GetEC2ClientVPNNetworkAssociationRegistryItem(),
	// GetEC2TrafficMirroSessionRegistryItem(),
	// GetEC2TransitGatewayPeeringAttachmentRegistryItem(),
	// GetEC2TransitGatewayVpcAttachmentRegistryItem(),
	// GetECRRegistryItem(),
	// GetECSServiceRegistryItem(),
	// GetEFSFileSystemRegistryItem(),
	// GetEIPRegistryItem(),
	// GetElastiCacheClusterItem(),
	GetElastiCacheReplicationGroupItem(),
	// GetElasticsearchDomainRegistryItem(),
	// GetELBRegistryItem(),
	// GetFSXWindowsFSRegistryItem(),
	GetInstanceRegistryItem(),
	// GetLambdaFunctionRegistryItem(),
	// GetLBRegistryItem(),
	// GetLightsailInstanceRegistryItem(),
	// GetMSKClusterRegistryItem(),
	// GetALBRegistryItem(),
	// GetMQBrokerRegistryItem(),
	GetNATGatewayRegistryItem(),
	GetRDSClusterRegistryItem(),
	GetRDSClusterInstanceRegistryItem(),
	// GetRedshiftClusterRegistryItem(),
	// GetRoute53HealthCheck(),
	// GetRoute53ResolverEndpointRegistryItem(),
	// GetRoute53RecordRegistryItem(),
	// GetRoute53ZoneRegistryItem(),
	GetS3BucketRegistryItem(),
	// GetS3BucketAnalyticsConfigurationRegistryItem(),
	// GetS3BucketInventoryRegistryItem(),
	// GetSecretsManagerSecret(),
	// GetSSMActivationRegistryItem(),
	// GetSSMParameterRegistryItem(),
	// GetSNSTopicRegistryItem(),
	// GetSNSTopicSubscriptionRegistryItem(),
	// GetSQSQueueRegistryItem(),
	// GetNewEKSNodeGroupItem(),
	// GetNewEKSFargateProfileItem(),
	GetNewEKSClusterItem(),
	// GetNewKMSKeyRegistryItem(),
	// GetNewKMSExternalKeyRegistryItem(),
	// GetVPNConnectionRegistryItem(),
	// GetVpcEndpointRegistryItem(),

}

// FreeResources grouped alphabetically
var FreeResources = []string{
	// Idem Core Helper States
	"states.exec.run",
	"states.acct.profile",

	// AWS States
	"states.aws.ec2.launch_template.present",
}

var UsageOnlyResources = []string{
	"aws_data_transfer",
}