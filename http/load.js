import http from 'k6/http';
import { sleep } from 'k6';

export default function () {
  let target = "http://localhost:8085"
  let res = http.get(target+"/users");
  // let res = http.post(target+"/soap", "<SoapRequest> <action>getUsers</action> </SoapRequest>")
  console.log(res.status)
}

