import sys
from awsglue.transforms import *
from awsglue.utils import getResolvedOptions
from pyspark.context import SparkContext
from awsglue.context import GlueContext
from awsglue.job import Job

args = getResolvedOptions(sys.argv, ['JOB_NAME'])
sc = SparkContext()
glueContext = GlueContext(sc)
spark = glueContext.spark_session
job = Job(glueContext)
job.init(args['JOB_NAME'], args)

# Script generated for node AWS Glue Data Catalog
AWSGlueDataCatalog_node1715961140345 = glueContext.create_dynamic_frame.from_catalog(database="iu-dataengineering-patrick-glue-database", table_name="data", transformation_ctx="AWSGlueDataCatalog_node1715961140345")

# Script generated for node Drop Fields
DropFields_node1715961142823 = DropFields.apply(frame=AWSGlueDataCatalog_node1715961140345, paths=["col2"], transformation_ctx="DropFields_node1715961142823")

# Script generated for node Amazon S3
AmazonS3_node1715961144750 = glueContext.write_dynamic_frame.from_options(frame=DropFields_node1715961142823, connection_type="s3", format="csv", connection_options={"path": "s3://iu-dataengineering-patrick/result/", "compression": "snappy", "partitionKeys": []}, transformation_ctx="AmazonS3_node1715961144750")

job.commit()