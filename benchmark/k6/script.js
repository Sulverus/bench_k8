import { check } from "k6";
import http from "k6/http";
import parseHTML from "k6/html";
 
export default function() {
    let res = http.request(
        "GET",
        "http://192.168.99.100/test-service",
        {},
        {}
    );
 
    check(res.body, {
        "contains hash": (r) => r.indexOf("hash") != -1
    });
 
    check(res, {
        "is status 200": (r) => r.status === 200
    });
};

