FROM tarantool/tarantool:1.7
COPY slave.lua /opt/tarantool/
WORKDIR /opt/tarantool

EXPOSE 3316
CMD ["tarantool", "/opt/tarantool/slave.lua"]
