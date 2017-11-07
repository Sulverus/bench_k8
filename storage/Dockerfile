FROM tarantool/tarantool:1.7
COPY master.lua /opt/tarantool/
WORKDIR /opt/tarantool

EXPOSE 3315
CMD ["tarantool", "/opt/tarantool/master.lua"]
