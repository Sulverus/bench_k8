#!/usr/bin/python
# -*- coding: utf-8 -*-

class AmmoGenerator(object):
    """
    Yandex tand ammo generator based on documentation:
    http://yandextank.readthedocs.org/en/latest/tutorial.html#preparing-requests
    """
    # tarantool upstream template
    smart_template = '%s\n'
    req_template = (
          "%s %s HTTP/1.1\r\n"
          "%s\r\n"
          "\r\n"
    )

    #http request with entity body template
    req_template_w_entity_body = (
          "%s %s HTTP/1.1\r\n"
          "%s\r\n"
          "Content-Length: %d\r\n"
          "\r\n"
          "%s\r\n"
    )

    # phantom ammo template
    ammo_template = (
        "%d %s\n"
        "%s"
    )
    headers = (
        "Host: 192.168.99.100\r\n"
        "User-agent: Tank\r\n"
        "Connection: Keep-Alive\r\n"
    )

    def __init__(self, method='POST', user_agent='tank'):
        self.requests = []
        self.method = method
        self.user_agent = user_agent

    def request(self, host, body, case='/'):
        """
        Returns tank ammo string: for uripost
        """
        req = self.smart_template % body
        return self.ammo_template % (len(req), case, req)

    def full_request(self, host, body, case='/service/'):
        """
        Returns tank ammo string: generic request
        """
        if not body:
            req = self.req_template % (self.method, case, self.headers)
        else:
            req = self.req_template_w_entity_body % (
                 self.method, case, self.headers, len(body), body
            )
        return self.ammo_template % (len(req), 'test', req)

    def append(self, host, body, case='/'):
        """
        Add request to container
        """
        self.requests.append(self.request(host, body, case))

    def export(self, filename):
        """
        Export requests to file
        """
        if not len(self.requests):
            return
        with open(filename, 'w') as ammo_file:
            for ammo in self.requests:
                ammo_file.write(ammo)
