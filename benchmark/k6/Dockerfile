FROM loadimpact/k6

RUN mkdir -p /opt/test/
COPY script.js /opt/test/
CMD ["run", "--vus", "5", "--duration", "60s", "/opt/test/script.js"]
