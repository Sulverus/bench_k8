#!/usr/bin/python                                                               
# -*- coding: utf-8 -*-                                                         
from ammo  import AmmoGenerator

import random

HOST="192.168.99.100"
URI='/test-service/tx?from=%d&to=%d'
FILE="ammo.txt"
COUNT=3000

def main():
    gen = AmmoGenerator(method="GET")
    with open(FILE, 'w') as expfile:
        for x in xrange(COUNT):
            a = random.randint(1, 100)
            b = random.randint(1, 100)
            expfile.write(gen.full_request(
                HOST, None, case=URI % (a, b)
            ))
    print "Created %d ammo in %s" % (COUNT, FILE)

if __name__ == '__main__':
    main()
