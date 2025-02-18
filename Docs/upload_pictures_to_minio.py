# https://www.cnblogs.com/xujunkai/p/14749165.html

from minio import Minio
from minio.error import S3Error
import os

MINIO_SERVER="192.168.2.128:9000"
ACCESS_KEY="admin"
SECRET_KEY="12345678"


try:
    # 初始化 MinIO 客户端
    minio_client = Minio(
        MINIO_SERVER,  #  MinIO 服务地址
        access_key=ACCESS_KEY,  #  Access Key
        secret_key=SECRET_KEY,  #  Secret Key
        secure=False  # 如果是 HTTP，则设置为 False，如果是 HTTPS，则设置为 True
    )
    print("MinIO client initialized successfully.")
except S3Error as err:
    print(f"Failed to initialize MinIO client. Error: {err}")
    exit(1)


# 桶名称
bucket_name = 'test'

# 检查桶是否存在
found = minio_client.bucket_exists(bucket_name)
if not found:
    # 创建桶
    minio_client.make_bucket(bucket_name)
    print(f"Bucket '{bucket_name}' created successfully.")
else:
    print(f"Bucket '{bucket_name}' already exists.")

# 上传 ./pictures 文件夹下的所有图片到桶
for file in os.listdir('./pictures'):
    if file.lower().endswith(('.png', '.jpg', '.jpeg', '.gif', '.bmp')):
        file_path = os.path.join('./pictures', file)
        try:
            # 上传文件
            minio_client.fput_object(
                bucket_name,
                file,
                file_path
            )
            print(f"File '{file}' uploaded to bucket '{bucket_name}'.")
        except S3Error as err:
            print(f"File '{file}' failed to upload. Error: {err}")


url=minio_client.presigned_get_object("test","1.png")

print(f'测试下载：{url}')