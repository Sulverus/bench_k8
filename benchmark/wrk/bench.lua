done = function(summary, latency, requests)
    io.write("---------------[Distribution]---------------\n")
    for _, p in pairs({ 99.999, 99, 90, 80, 70, 60, 50}) do
        n = latency:percentile(p) / 1000.0
        io.write(string.format(" -- %g%%:\t%dms\n", p, n))
    end
end

request = function()
    wrk.headers["Content-Type"] = "application/json; charset=utf-8;"
    return wrk.format("GET", wrk.path, wrk.headers)
end
