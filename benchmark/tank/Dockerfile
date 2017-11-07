FROM direvius/yandex-tank
COPY ammo.txt /var/loadtest
COPY tank.yaml /var/loadtest
COPY token.txt /var/loadtest
COPY run.sh /var/loadtest
EXPOSE 8080
ENTRYPOINT ["/bin/bash"]
CMD ["run.sh"]
