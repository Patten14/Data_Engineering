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

# Script generated for node Amazon S3
AmazonS3_node1715374377686 = glueContext.create_dynamic_frame.from_options(format_options={"quoteChar": "\"", "withHeader": False, "separator": ";", "optimizePerformance": False}, connection_type="s3", format="csv", connection_options={"paths": ["s3://iu-test-keanu/data/ML_DATA.csv"], "recurse": True}, transformation_ctx="AmazonS3_node1715374377686")

# Script generated for node Drop Fields
DropFields_node1715374387163 = DropFields.apply(frame=AmazonS3_node1715374377686, paths=["col2"], transformation_ctx="DropFields_node1715374387163")

# Script generated for node Amazon S3
AmazonS3_node1715374394703 = glueContext.write_dynamic_frame.from_options(
    frame=DropFields_node1715374387163,
    connection_type="s3",
    connection_options={
        "path": "s3://iu-test-keanu/result/",
    },
    format="csv",
    #format_options={"writeHeader": True},  # Wenn du möchtest, dass die Headerzeile geschrieben wird
    transformation_ctx="AmazonS3_node1715374394703"
)

job.commit()