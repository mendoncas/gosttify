import grpc from 'k6/net/grpc';
import { sleep } from 'k6';

const client = new grpc.Client()
client.load([''], 'gostify.proto')

export default function () {
  client.connect('grpcbin.test.k6.io:50000', {
  })

  const data = {username: "ricardo"}
  const response = cient.invoke('gostify.Gostify/GetUsers', data)

  console.log(JSON.stringify(response.message))
  client.close()
}

