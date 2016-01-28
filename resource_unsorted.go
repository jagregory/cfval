package main

func cacheCluster() Resource {
	return Resource{
		AwsType:    "AWS::ElastiCache::CacheCluster",
		Properties: map[string]Schema{},
	}
}

func subnetGroup() Resource {
	return Resource{
		AwsType:    "AWS::ElastiCache::SubnetGroup",
		Properties: map[string]Schema{},
	}
}

func instanceProfile() Resource {
	return Resource{
		AwsType:    "AWS::IAM::InstanceProfile",
		Properties: map[string]Schema{},
	}
}

func loadBalancer() Resource {
	return Resource{
		AwsType:    "AWS::ElasticLoadBalancing::LoadBalancer",
		Properties: map[string]Schema{},
	}
}

func application() Resource {
	return Resource{
		AwsType:    "AWS::ElasticBeanstalk::Application",
		Properties: map[string]Schema{},
	}
}

func configurationTemplate() Resource {
	return Resource{
		AwsType:    "AWS::ElasticBeanstalk::ConfigurationTemplate",
		Properties: map[string]Schema{},
	}
}

func environment() Resource {
	return Resource{
		AwsType:    "AWS::ElasticBeanstalk::Environment",
		Properties: map[string]Schema{},
	}
}

func applicationVersion() Resource {
	return Resource{
		AwsType:    "AWS::ElasticBeanstalk::ApplicationVersion",
		Properties: map[string]Schema{},
	}
}

func topic() Resource {
	return Resource{
		AwsType:    "AWS::SNS::Topic",
		Properties: map[string]Schema{},
	}
}
