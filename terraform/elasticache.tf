resource "aws_elasticache_cluster" "redis" {
  cluster_id           = "${local.name_base}-redis"
  engine               = "redis"
  node_type            = "cache.t3.micro"
  num_cache_nodes      = 1
  parameter_group_name = aws_elasticache_parameter_group.redis.id
  engine_version       = "7.0"
  port                 = 6379
  subnet_group_name    = module.vpc.elasticache_subnet_group_name
  security_group_ids   = [module.security-group_kvs.security_group_id]

  snapshot_retention_limit = 5
}

resource "aws_elasticache_parameter_group" "redis" {
  name   = "${local.name_base}-redis-param-group"
  family = "redis7"
}
