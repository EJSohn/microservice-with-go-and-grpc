
from __future__ import print_function

import random
import grpc
import api_pb2
import api_pb2_grpc


def get_point(stub, profile):
    point = stub.GetPoint(profile)
    print(point)

def list_users(stub, point):
    users = stub.ListUsers(point)
    for user in users:
        print(user.profile.name)

def run():
    with grpc.insecure_channel('localhost:10000') as channel:
        stub = api_pb2_grpc.UserGuideStub(channel)
        print('-----GetPoint-----')
        profile = api_pb2.Profile(
            age=26,
            name="james"
        )
        get_point(stub, profile)
        print('-----ListUsers-----')
        point = api_pb2.Point(
            value=1000
        )
        list_users(stub, point)
        

if __name__ == "__main__":
    run()

